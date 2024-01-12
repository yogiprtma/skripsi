package handler

import (
	"net/http"
	"project-skripsi/user"
	"strconv"

	"github.com/gin-gonic/gin"
)

type adminKunciHandler struct {
	userService user.Service
}

func NewAdminKunciHandler(userService user.Service) *adminKunciHandler {
	return &adminKunciHandler{userService}
}

func (h *adminKunciHandler) Index(c *gin.Context) {
	users, err := h.userService.GetAllUsers()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.HTML(http.StatusOK, "kunci_index.html", gin.H{
		"users": user.FormatDataUsers(users),
	})
}

func (h *adminKunciHandler) Detail(c *gin.Context) {
	idParam := c.Param("id")

	id, _ := strconv.Atoi(idParam)

	user, err := h.userService.GetUserByID(id)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}
	c.HTML(http.StatusOK, "kunci_detail.html", user)

}
