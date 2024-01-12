package handler

import (
	"net/http"
	"project-skripsi/document_legalization"
	"project-skripsi/helper"
	"project-skripsi/subject"
	"project-skripsi/user"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type userLegalisasiHandler struct {
	subjectService              subject.Service
	documentLegalizationService document_legalization.Service
	userService                 user.Service
}

func NewUserLegalisasiHandler(subjectService subject.Service, documentLegalizationService document_legalization.Service, userService user.Service) *userLegalisasiHandler {
	return &userLegalisasiHandler{subjectService, documentLegalizationService, userService}
}

func (h *userLegalisasiHandler) Index(c *gin.Context) {
	subjects, err := h.subjectService.GetAllSubjects()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	session := sessions.Default(c)
	if session.Get("userID") == nil {
		return
	}
	currentUserID := session.Get("userID")
	userID := currentUserID.(int)

	ok, err := h.documentLegalizationService.CheckUserIDWhereStatusInProcess(userID)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	if ok {
		c.HTML(http.StatusOK, "legalisasi-not-available.html", nil)
		return
	}

	c.HTML(http.StatusOK, "user_legalisasi_index.html", gin.H{
		"Subjects": subjects,
	})
}

func (h *userLegalisasiHandler) Create(c *gin.Context) {
	var input document_legalization.FormCreateAppDocumentLegalizationInput
	subjects, err := h.subjectService.GetAllSubjects()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	input.Subjects = subjects
	session := sessions.Default(c)
	currentUserID := session.Get("userID")
	input.UserID = currentUserID.(int)
	userID := currentUserID.(int)

	err = c.ShouldBind(&input)
	if err != nil {
		input.Error = input.Validate()
		c.HTML(http.StatusOK, "user_legalisasi_index.html", input)
		return
	}

	// _ ganti user untuk kirim email
	user, err := h.userService.GetUserByID(userID)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	_, err = h.documentLegalizationService.CreateAppDocumentLegalization(input)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}
	// nanti balikkan untuk send email

	go helper.SendEmailToUserForSubmission(input.Email, user.Name, user.NPM)

	c.Redirect(http.StatusFound, "/pengajuan-legalisasi")
}
