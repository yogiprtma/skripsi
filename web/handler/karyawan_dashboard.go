package handler

import (
	"fmt"
	"net/http"
	"project-skripsi/document_legalization"
	"project-skripsi/helper"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type karyawanDashboardHandler struct {
	documentLegalizationService document_legalization.Service
}

func NewKaryawanDashboardHandler(documentLegalizationService document_legalization.Service) *karyawanDashboardHandler {
	return &karyawanDashboardHandler{documentLegalizationService}
}

func (h *karyawanDashboardHandler) Index(c *gin.Context) {
	session := sessions.Default(c)

	if session.Get("roleKaryawan") == nil {
		return
	}
	if session.Get("nameKaryawan") == nil {
		return
	}
	currentRolekaryawan := session.Get("roleKaryawan")
	roleKaryawan := currentRolekaryawan.(string)
	currentNameKaryawan := session.Get("nameKaryawan").(string)

	if roleKaryawan == "Karyawan Akademik Bagian Informatika dan Sistem Informasi" {
		documentLegalizations, err := h.documentLegalizationService.GetAllDocForKaryawan("Informatika", "Sistem Informasi")
		if err != nil {
			c.HTML(http.StatusInternalServerError, "error.html", nil)
			return
		}
		c.HTML(http.StatusOK, "karyawan_dashboard_index.html", gin.H{
			"documentLegalizations": document_legalization.FormatDataDocumentLegalizations(documentLegalizations),
			"name":                  currentNameKaryawan,
			"total":                 len(documentLegalizations),
		})
		return
	}

	if roleKaryawan == "Karyawan Akademik Bagian Teknik Sipil dan Arsitektur" {
		documentLegalizations, err := h.documentLegalizationService.GetAllDocForKaryawan("Teknik Sipil", "Teknik Elektro")
		if err != nil {
			c.HTML(http.StatusInternalServerError, "error.html", nil)
			return
		}
		c.HTML(http.StatusOK, "karyawan_dashboard_index.html", gin.H{
			"documentLegalizations": document_legalization.FormatDataDocumentLegalizations(documentLegalizations),
			"name":                  currentNameKaryawan,
			"total":                 len(documentLegalizations),
		})
		return
	}

	if roleKaryawan == "Karyawan Akademik Bagian Teknik Mesin dan Teknik Elektro" {
		documentLegalizations, err := h.documentLegalizationService.GetAllDocForKaryawan("Teknik Mesin", "Arsitektur")
		if err != nil {
			c.HTML(http.StatusInternalServerError, "error.html", nil)
			return
		}
		c.HTML(http.StatusOK, "karyawan_dashboard_index.html", gin.H{
			"documentLegalizations": document_legalization.FormatDataDocumentLegalizations(documentLegalizations),
			"name":                  currentNameKaryawan,
			"total":                 len(documentLegalizations),
		})
		return
	}
}

func (h *karyawanDashboardHandler) New(c *gin.Context) {
	idParam := c.Param("id")

	id, _ := strconv.Atoi(idParam)

	documentLegalization, err := h.documentLegalizationService.GetDocumentLegalizationByID(id)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	var input document_legalization.FormUpdateDocumentLegalizationInput

	input.ID = documentLegalization.ID
	input.NameUser = documentLegalization.User.Name
	input.NpmUser = documentLegalization.User.NPM
	input.NameSubject = documentLegalization.Subject.Name

	c.HTML(200, "karyawan_dashboard_new.html", input)
}

func (h *karyawanDashboardHandler) Reject(c *gin.Context) {
	var input document_legalization.FormRejectedDocumentInput
	err := c.ShouldBind(&input)
	if err != nil {
		c.HTML(http.StatusOK, "karyawan_dashboard_index.html", input)
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

	c.Redirect(http.StatusFound, "/karyawan/dashboard")
}

func (h *karyawanDashboardHandler) Update(c *gin.Context) {
	idParam := c.Param("id")

	id, _ := strconv.Atoi(idParam)

	var input document_legalization.FormUpdateDocToApprovedByKaryawanInput

	err := c.ShouldBind(&input)
	if err != nil {
		input.ID = id
		input.Error = err
		c.HTML(http.StatusOK, "karyawan_dashboard_new.html", input)
		return
	}

	file, err := c.FormFile("file_transkrip")
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	uuid := uuid.NewString()

	path := fmt.Sprintf("file_document/%v.pdf", uuid)
	fileName := fmt.Sprintf("%v.pdf", uuid)
	c.SaveUploadedFile(file, path)

	input.ID = id
	input.FileNameDocument = fileName
	input.UUID = uuid

	_, err = h.documentLegalizationService.UpdateDocumentToApprovedByKaryawan(input)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.Redirect(http.StatusFound, "/karyawan/dashboard")
}
