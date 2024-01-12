package handler

import (
	"net/http"
	"project-skripsi/document_legalization"
	"project-skripsi/employee"
	"project-skripsi/subject"
	"project-skripsi/user"

	"github.com/gin-gonic/gin"
)

type adminDashboardHandler struct {
	documentLegalizationService document_legalization.Service
	employeeService             employee.Service
	userService                 user.Service
	SubjectService              subject.Service
}

func NewAdminDashboardHandler(documentLegalizationService document_legalization.Service, employeeService employee.Service, userService user.Service, subjectService subject.Service) *adminDashboardHandler {
	return &adminDashboardHandler{documentLegalizationService, employeeService, userService, subjectService}
}

func (h *adminDashboardHandler) Index(c *gin.Context) {
	users, err := h.userService.GetAllUsers()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	employees, err := h.employeeService.GetAllEmployees()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	subjects, err := h.SubjectService.GetAllSubjects()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.HTML(http.StatusOK, "dashboard_index.html", gin.H{
		"users":     len(users),
		"employees": len(employees),
		"subjects":  len(subjects),
	})
}
