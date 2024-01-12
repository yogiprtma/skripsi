package handler

import (
	"net/http"
	"project-skripsi/document_legalization"
	"project-skripsi/helper"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type kaprodiDashboardHandler struct {
	documentLegalizationService document_legalization.Service
}

func NewKaprodiDashboardHandler(documentLegalizationService document_legalization.Service) *kaprodiDashboardHandler {
	return &kaprodiDashboardHandler{documentLegalizationService}
}

func (h *kaprodiDashboardHandler) Index(c *gin.Context) {
	session := sessions.Default(c)

	if session.Get("roleKaprodi") == nil {
		return
	}
	if session.Get("nameKaprodi") == nil {
		return
	}

	currentRoleKaprodi := session.Get("roleKaprodi").(string)
	currentNameKaprodi := session.Get("nameKaprodi").(string)

	if currentRoleKaprodi == "Koordinator Prodi Informatika" {
		documentLegalizations, err := h.documentLegalizationService.GetAllDocForKaprodi("Informatika")
		if err != nil {
			c.HTML(http.StatusInternalServerError, "error.html", nil)
			return
		}
		c.HTML(http.StatusOK, "kaprodi_dashboard_index.html", gin.H{
			"documentLegalizations": document_legalization.FormatDataDocumentLegalizations(documentLegalizations),
			"name":                  currentNameKaprodi,
			"total":                 len(documentLegalizations),
		})
		return
	}

	if currentRoleKaprodi == "Koordinator Prodi Teknik Sipil" {
		documentLegalizations, err := h.documentLegalizationService.GetAllDocForKaprodi("Teknik Sipil")
		if err != nil {
			c.HTML(http.StatusInternalServerError, "error.html", nil)
			return
		}
		c.HTML(http.StatusOK, "kaprodi_dashboard_index.html", gin.H{
			"documentLegalizations": document_legalization.FormatDataDocumentLegalizations(documentLegalizations),
			"name":                  currentNameKaprodi,
			"total":                 len(documentLegalizations),
		})
		return
	}

	if currentRoleKaprodi == "Koordinator Prodi Teknik Elektro" {
		documentLegalizations, err := h.documentLegalizationService.GetAllDocForKaprodi("Teknik Elektro")
		if err != nil {
			c.HTML(http.StatusInternalServerError, "error.html", nil)
			return
		}
		c.HTML(http.StatusOK, "kaprodi_dashboard_index.html", gin.H{
			"documentLegalizations": document_legalization.FormatDataDocumentLegalizations(documentLegalizations),
			"name":                  currentNameKaprodi,
			"total":                 len(documentLegalizations),
		})
		return
	}

	if currentRoleKaprodi == "Koordinator Prodi Teknik Mesin" {
		documentLegalizations, err := h.documentLegalizationService.GetAllDocForKaprodi("Teknik Mesin")
		if err != nil {
			c.HTML(http.StatusInternalServerError, "error.html", nil)
			return
		}
		c.HTML(http.StatusOK, "kaprodi_dashboard_index.html", gin.H{
			"documentLegalizations": document_legalization.FormatDataDocumentLegalizations(documentLegalizations),
			"name":                  currentNameKaprodi,
			"total":                 len(documentLegalizations),
		})
		return
	}

	if currentRoleKaprodi == "Koordinator Prodi Arsitektur" {
		documentLegalizations, err := h.documentLegalizationService.GetAllDocForKaprodi("Arsitektur")
		if err != nil {
			c.HTML(http.StatusInternalServerError, "error.html", nil)
			return
		}
		c.HTML(http.StatusOK, "kaprodi_dashboard_index.html", gin.H{
			"documentLegalizations": document_legalization.FormatDataDocumentLegalizations(documentLegalizations),
			"name":                  currentNameKaprodi,
			"total":                 len(documentLegalizations),
		})
		return
	}

	if currentRoleKaprodi == "Koordinator Prodi Sistem Informasi" {
		documentLegalizations, err := h.documentLegalizationService.GetAllDocForKaprodi("Sistem Informasi")
		if err != nil {
			c.HTML(http.StatusInternalServerError, "error.html", nil)
			return
		}
		c.HTML(http.StatusOK, "kaprodi_dashboard_index.html", gin.H{
			"documentLegalizations": document_legalization.FormatDataDocumentLegalizations(documentLegalizations),
			"name":                  currentNameKaprodi,
			"total":                 len(documentLegalizations),
		})
		return
	}
}

func (h *kaprodiDashboardHandler) Detail(c *gin.Context) {
	idParam := c.Param("id")

	id, _ := strconv.Atoi(idParam)

	documentLegalization, err := h.documentLegalizationService.GetDocumentLegalizationByID(id)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}
	c.HTML(http.StatusOK, "kaprodi_dashboard_detail.html", document_legalization.FormatDataDocumentLegalization(documentLegalization))
}

func (h *kaprodiDashboardHandler) Reject(c *gin.Context) {
	var input document_legalization.FormRejectedDocumentInput
	err := c.ShouldBind(&input)
	if err != nil {
		c.HTML(http.StatusOK, "kaprodi_dashboard_index.html", input)
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

	c.Redirect(http.StatusFound, "/koordinator-prodi/dashboard")
}

func (h *kaprodiDashboardHandler) Update(c *gin.Context) {
	idParam := c.Param("id")

	id, _ := strconv.Atoi(idParam)

	_, err := h.documentLegalizationService.UpdateDocumentToApprovedByKaprodi(id)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.Redirect(http.StatusFound, "/koordinator-prodi/dashboard")
}
