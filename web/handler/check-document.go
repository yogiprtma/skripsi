package handler

import (
	"errors"
	"io"
	"math/big"
	"net/http"
	"project-skripsi/document_legalization"
	"project-skripsi/signature"
	"project-skripsi/user"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type CheckDocumentHandler struct {
	userService                 user.Service
	documentLegalizationService document_legalization.Service
}

func NewCheckDocumentHandler(userService user.Service, documentLegalizationService document_legalization.Service) *CheckDocumentHandler {
	return &CheckDocumentHandler{userService, documentLegalizationService}
}

func (h *CheckDocumentHandler) New(c *gin.Context) {
	c.HTML(http.StatusOK, "check_document", nil)
}

func (h *CheckDocumentHandler) Invalid(c *gin.Context) {
	c.HTML(http.StatusOK, "document_not_found", nil)
}

func (h *CheckDocumentHandler) Find(c *gin.Context) {
	var input user.FormFindPublicAndPrivateKeyUserInput

	err := c.ShouldBind(&input)
	if err != nil {
		input.Error = input.Validate()
		c.HTML(http.StatusOK, "check_document", input)
		return
	}

	user, err := h.userService.GetPublicAndPrivateKeyByNPM(input)
	if err != nil {
		input.ErrorFindNPM = err
		input.Error = input.Validate()
		c.HTML(http.StatusOK, "check_document", input)
		return
	}

	privateKey := strings.Split(user.PrivateKey, ", ")
	publicKey := strings.Split(user.PublicKey, ", ")

	d := new(big.Int)
	d.SetString(privateKey[0], 10)

	e := new(big.Int)
	e.SetString(publicKey[0], 10)

	n := new(big.Int)
	n.SetString(publicKey[1], 10)

	file, err := c.FormFile("file_transkrip")
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	openedFile, _ := file.Open()
	openedFile.Close()
	pdfBytes, err := io.ReadAll(openedFile)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	MD := signature.HashMsg(pdfBytes)

	generateSignature := signature.GenerateSignature(MD, d, n)

	msgDigest := signature.DecryptSignature(generateSignature, e, n)

	ok, err := h.documentLegalizationService.CheckMessageDigest(msgDigest)
	if err != nil {
		input.Error = err
		c.HTML(http.StatusOK, "check_document", input)
		return
	}

	if ok {
		documentLegalization, err := h.documentLegalizationService.GetMessageDigest(msgDigest)
		if err != nil {
			input.Error = err
			c.HTML(http.StatusOK, "check_document", input)
			return
		}

		currentUserPublicKey := strings.Split(documentLegalization.User.PublicKey, ", ")

		currentN := new(big.Int)
		currentN.SetString(currentUserPublicKey[1], 10)
		if n.String() != currentN.String() {
			input.ErrorKey = errors.New("kunci tidak sepadan")
			c.HTML(http.StatusOK, "check_document", input)
			return
		} else if n.String() == currentN.String() {
			c.Redirect(http.StatusMovedPermanently, "/dokumen/"+documentLegalization.UUID)
		} else {
			c.Redirect(http.StatusMovedPermanently, "/dokumen/dokumen-tidak-valid")
		}
	} else {
		c.Redirect(http.StatusMovedPermanently, "/dokumen/dokumen-tidak-valid")
	}
}

func (h *CheckDocumentHandler) Index(c *gin.Context) {
	uuid := c.Param("uuid")

	documentLegalization, err := h.documentLegalizationService.GetDocumentLegalizationByUUID(uuid)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	isExpired := "VALID"

	if time.Now().After(documentLegalization.ExpiredAt) {
		isExpired = "EXPIRED"
	}
	c.HTML(http.StatusOK, "document_found", gin.H{
		"DocumentLegalization": document_legalization.FormatDataDocumentLegalization(documentLegalization),
		"IsExpired":            isExpired,
	})
}
