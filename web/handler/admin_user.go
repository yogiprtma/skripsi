package handler

import (
	"net/http"
	"project-skripsi/user"
	"strconv"

	"github.com/gin-gonic/gin"
)

type adminUserHandler struct {
	userService user.Service
}

func NewAdminUserHandler(userService user.Service) *adminUserHandler {
	return &adminUserHandler{userService}
}

func (h *adminUserHandler) Index(c *gin.Context) {
	users, err := h.userService.GetAllUsers()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.HTML(http.StatusOK, "user_index.html", gin.H{
		"users": user.FormatDataUsers(users),
	})

}

func (h *adminUserHandler) New(c *gin.Context) {
	c.HTML(http.StatusOK, "user_new.html", nil)
}

func (h *adminUserHandler) Create(c *gin.Context) {
	var input user.FormCreateUserInput

	err := c.ShouldBind(&input)
	if err != nil {
		input.Error = input.Validate()
		c.HTML(http.StatusOK, "user_new.html", input)
		return
	}

	_, err = h.userService.RegisterUser(input)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.Redirect(http.StatusFound, "/admin/akun")
}

func (h *adminUserHandler) Edit(c *gin.Context) {
	idParam := c.Param("id")

	id, _ := strconv.Atoi(idParam)

	registeredUser, err := h.userService.GetUserByID(id)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	input := user.FormUpdateUserInput{}
	input.ID = registeredUser.ID
	input.Name = registeredUser.Name
	input.NPM = registeredUser.NPM
	input.Department = registeredUser.Department

	c.HTML(http.StatusOK, "user_edit.html", input)
}

func (h *adminUserHandler) Update(c *gin.Context) {
	idParam := c.Param("id")

	id, _ := strconv.Atoi(idParam)

	var input user.FormUpdateUserInput

	err := c.ShouldBind(&input)
	if err != nil {
		input.ID = id
		input.Error = input.Validate()
		c.HTML(http.StatusOK, "user_edit.html", input)
		return
	}

	input.ID = id

	_, err = h.userService.UpdateUser(input)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.Redirect(http.StatusFound, "/admin/akun")
}

func (h *adminUserHandler) Delete(c *gin.Context) {
	idParam := c.Param("id")

	id, _ := strconv.Atoi(idParam)

	_, err := h.userService.DeleteUser(id)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.Redirect(http.StatusFound, "/admin/akun")
}

func (h *adminUserHandler) EditPassword(c *gin.Context) {
	idParam := c.Param("id")

	id, _ := strconv.Atoi(idParam)

	registeredUser, err := h.userService.GetUserByID(id)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	input := user.FormUpdatePasswordUserInput{}
	input.ID = registeredUser.ID
	input.Password = registeredUser.PasswordHash

	c.HTML(http.StatusOK, "user_edit_password.html", input)
}

func (h *adminUserHandler) UpdatePassword(c *gin.Context) {
	idParam := c.Param("id")

	id, _ := strconv.Atoi(idParam)

	var input user.FormUpdatePasswordUserInput

	err := c.ShouldBind(&input)
	if input.Password != input.NewPassword {
		input.ID = id
		input.ErrorPassword = "Password baru tidak sama."
		c.HTML(http.StatusOK, "user_edit_password.html", input)
		return
	}
	if err != nil {
		input.ID = id
		input.Error = input.Validate()
		c.HTML(http.StatusOK, "user_edit_password.html", input)
		return
	}

	input.ID = id

	_, err = h.userService.UpdatePasswordUser(input)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.Redirect(http.StatusFound, "/admin/akun")
}
