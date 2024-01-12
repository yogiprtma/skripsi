package handler

import (
	"net/http"
	"project-skripsi/document_legalization"
	"strconv"

	"github.com/gin-gonic/gin"
)

type adminDokumenHandler struct {
	documentLegalizationService document_legalization.Service
}

func NewAdminDokumenHandler(documentLegalizationService document_legalization.Service) *adminDokumenHandler {
	return &adminDokumenHandler{documentLegalizationService}
}

func (h *adminDokumenHandler) Index(c *gin.Context) {
	documentLegalizations, err := h.documentLegalizationService.GetAllDocumentLegalizationsWhereStatusSigned()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.HTML(http.StatusOK, "admin_dokumen_index.html", gin.H{
		"documentLegalizations": document_legalization.FormatDataDocumentLegalizations(documentLegalizations),
	})
}

func (h *adminDokumenHandler) Detail(c *gin.Context) {
	idParam := c.Param("id")

	id, _ := strconv.Atoi(idParam)

	documentLegalization, err := h.documentLegalizationService.GetDocumentLegalizationByID(id)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}
	c.HTML(http.StatusOK, "admin_dokumen_detail.html", document_legalization.FormatDataDocumentLegalization(documentLegalization))
}
