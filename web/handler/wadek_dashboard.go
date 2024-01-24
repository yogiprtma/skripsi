package handler

import (
	"fmt"
	"log"
	"math/big"
	"net/http"
	"os"
	"project-skripsi/document_legalization"
	"project-skripsi/generate_pdf"
	"project-skripsi/helper"
	"project-skripsi/signature"
	"strconv"
	"strings"
	"sync"
	"time"

	_ "github.com/joho/godotenv/autoload"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/zeebo/blake3"
)

type wadekDashboardHandler struct {
	documentLegalizationService document_legalization.Service
}

func NewWadekDashboardHandler(documentLegalizationService document_legalization.Service) *wadekDashboardHandler {
	return &wadekDashboardHandler{documentLegalizationService}
}

func (h *wadekDashboardHandler) Index(c *gin.Context) {
	session := sessions.Default(c)

	if session.Get("nameWadek") == nil {
		return
	}

	currentNameWadek := session.Get("nameWadek").(string)

	documentLegalizations, err := h.documentLegalizationService.GetAllDocForWadek()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}
	c.HTML(http.StatusOK, "wadek_dashboard_index.html", gin.H{
		"documentLegalizations": document_legalization.FormatDataDocumentLegalizations(documentLegalizations),
		"name":                  currentNameWadek,
		"total":                 len(documentLegalizations),
	})
}

func (h *wadekDashboardHandler) Detail(c *gin.Context) {
	idParam := c.Param("id")

	id, _ := strconv.Atoi(idParam)

	documentLegalization, err := h.documentLegalizationService.GetDocumentLegalizationByID(id)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}
	c.HTML(http.StatusOK, "wadek_dashboard_detail.html", document_legalization.FormatDataDocumentLegalization(documentLegalization))
}

func (h *wadekDashboardHandler) Reject(c *gin.Context) {
	var input document_legalization.FormRejectedDocumentInput
	err := c.ShouldBind(&input)
	if err != nil {
		c.HTML(http.StatusOK, "wadek_dashboard_index.html", input)
		return
	}
	getDocumentLegalization, err := h.documentLegalizationService.GetDocumentLegalizationByID(input.ID)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	_, err = h.documentLegalizationService.UpdateDocumentToRejected(getDocumentLegalization.ID)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}
	go helper.SendEmailToUserForRejected(getDocumentLegalization.Email, getDocumentLegalization.User.Name, getDocumentLegalization.User.NPM, input.Reason)

	c.Redirect(http.StatusFound, "/wadek/dashboard")
}

func (h *wadekDashboardHandler) Update(c *gin.Context) {
	idParam := c.Param("id")

	id, _ := strconv.Atoi(idParam)

	var input document_legalization.FormUpdateDocToSignedByWadekInput
	var wg sync.WaitGroup

	documentLegalization, err := h.documentLegalizationService.GetDocumentLegalizationByID(id)
	if err != nil {
		log.Printf("Error : %v", err)
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	privateKey := strings.Split(documentLegalization.User.PrivateKey, ", ")

	d := new(big.Int)
	d.SetString(privateKey[0], 10)

	n := new(big.Int)
	n.SetString(privateKey[1], 10)

	err = c.ShouldBind(&input)
	if err != nil {
		input.ID = id
		input.Error = err
		c.HTML(http.StatusOK, "wadek_dashboard_new.html", input)
		return
	}

	input.ID = id
	pwd, _ := os.Getwd()
	file, err := os.ReadFile(pwd + "/file_document/" + documentLegalization.FileNameDocument)
	if err != nil {
		log.Printf("Error : %v", err)
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}
	expiredAt := time.Now().Add(time.Hour * 24 * 33)
	expiredAtString := expiredAt.Format("01/02/2006")

	session := sessions.Default(c)
	currentNameWadek := session.Get("nameWadek")
	SignedByWadek := fmt.Sprintf("%v (%v)", currentNameWadek.(string), time.Now().Format("01/02/2006"))

	path := fmt.Sprintf("file_signed/%v.pdf", documentLegalization.UUID)

	contentQRCode := fmt.Sprintf(
		`
	Disetujui Oleh :
	1. %v
	2. %v

	Ditandatangani Oleh :
	%v

	Pranala Dokumen :
	%v/dokumen/%v
	`, documentLegalization.ApprovedByKaryawanAkademik, documentLegalization.ApprovedByKaprodi, SignedByWadek, os.Getenv("BASE_URL_SERVER"), documentLegalization.UUID)

	wg.Add(1)
	go func() {
		generate_pdf.GeneratePDF(file, path, contentQRCode, expiredAtString)
		wg.Done()
	}()
	wg.Wait()
	message, err := os.ReadFile(pwd + "/file_signed/" + documentLegalization.FileNameDocument)
	if err != nil {
		log.Printf("Error : %v", err)
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	MD := signature.HashMsg(message)

	msgDigest := blake3.Sum256(message)

	generateSignature := signature.GenerateSignature(MD, d, n)

	input.MessageDigest = fmt.Sprintf("%x", msgDigest)
	input.Signature = generateSignature
	input.ExpiredAt = expiredAt
	input.SignedByWadek = SignedByWadek

	_, err = h.documentLegalizationService.UpdateDocumentToSignedByWadek(input)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	// go helper.SendEmailToUserForApproved(documentLegalization.Email, documentLegalization.User.Name, documentLegalization.User.NPM, fmt.Sprintf("%v.pdf", documentLegalization.UUID))

	c.Redirect(http.StatusFound, "/wadek/dashboard")
}
