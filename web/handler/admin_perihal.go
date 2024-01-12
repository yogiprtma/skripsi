package handler

import (
	"net/http"
	"project-skripsi/subject"
	"strconv"

	"github.com/gin-gonic/gin"
)

type adminPerihalHandler struct {
	subjectService subject.Service
}

func NewAdminPerihalHandler(subjectService subject.Service) *adminPerihalHandler {
	return &adminPerihalHandler{subjectService}
}

func (h *adminPerihalHandler) Index(c *gin.Context) {
	subjects, err := h.subjectService.GetAllSubjects()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.HTML(http.StatusOK, "perihal_index.html", gin.H{
		"subjects": subject.FormatDataSubjects(subjects),
	})
}

func (h *adminPerihalHandler) New(c *gin.Context) {
	c.HTML(http.StatusOK, "perihal_new.html", nil)
}

func (h *adminPerihalHandler) Create(c *gin.Context) {
	var input subject.FormCreateSubjectInput

	err := c.ShouldBind(&input)
	if err != nil {
		input.Error = input.Validate()
		c.HTML(http.StatusOK, "perihal_new.html", input)
		return
	}

	_, err = h.subjectService.CreateSubject(input)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.Redirect(http.StatusFound, "/admin/perihal")
}

func (h *adminPerihalHandler) Edit(c *gin.Context) {
	idParam := c.Param("id")

	id, _ := strconv.Atoi(idParam)

	getSubject, err := h.subjectService.GetSubjectByID(id)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	input := subject.FormUpdateSubjectInput{}
	input.ID = getSubject.ID
	input.Name = getSubject.Name

	c.HTML(http.StatusOK, "perihal_edit.html", input)
}

func (h *adminPerihalHandler) Update(c *gin.Context) {
	idParam := c.Param("id")

	id, _ := strconv.Atoi(idParam)

	var input subject.FormUpdateSubjectInput

	err := c.ShouldBind(&input)
	if err != nil {
		input.ID = id
		input.Error = input.Validate()
		c.HTML(http.StatusOK, "perihal_edit.html", input)
		return
	}

	input.ID = id

	_, err = h.subjectService.UpdateSubject(input)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.Redirect(http.StatusFound, "/admin/perihal")
}

func (h *adminPerihalHandler) Delete(c *gin.Context) {
	idParam := c.Param("id")

	id, _ := strconv.Atoi(idParam)

	_, err := h.subjectService.DeleteSubject(id)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.Redirect(http.StatusFound, "/admin/perihal")
}
