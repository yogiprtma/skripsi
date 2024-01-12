package handler

import (
	"net/http"
	"project-skripsi/document_legalization"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type userDashboardHandler struct {
	documentLegalizationService document_legalization.Service
}

func NewUserDashboardHandler(documentLegalizationService document_legalization.Service) *userDashboardHandler {
	return &userDashboardHandler{documentLegalizationService}
}

func (h *userDashboardHandler) Index(c *gin.Context) {
	session := sessions.Default(c)

	if session.Get("userID") == nil {
		return
	}
	currentUserID := session.Get("userID")
	userID := currentUserID.(int)
	currentUserName := session.Get("userName").(string)
	currentUserNPM := session.Get("userNPM").(string)

	documentLegalizations, err := h.documentLegalizationService.GetAllDocumentLegalizationsByUserIDWhereStatusSigned(userID)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	ok, err := h.documentLegalizationService.CheckUserIDWhereStatusInProcess(userID)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.HTML(http.StatusOK, "user_dashboard_index.html", gin.H{
		"Ok":                    ok,
		"documentLegalizations": document_legalization.FormatDataDocumentLegalizations(documentLegalizations),
		"name":                  currentUserName,
		"npm":                   currentUserNPM,
	})
}

func (h *userDashboardHandler) Detail(c *gin.Context) {
	idParam := c.Param("id")

	id, _ := strconv.Atoi(idParam)

	documentLegalization, err := h.documentLegalizationService.GetDocumentLegalizationByID(id)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}
	c.HTML(http.StatusOK, "user_dashboard_detail.html", document_legalization.FormatDataDocumentLegalization(documentLegalization))
}
