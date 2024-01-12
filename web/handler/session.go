package handler

import (
	"net/http"
	"project-skripsi/employee"
	"project-skripsi/user"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type sessionHandler struct {
	userService     user.Service
	employeeService employee.Service
}

func NewSessionHandler(userService user.Service, employeeService employee.Service) *sessionHandler {
	return &sessionHandler{userService, employeeService}
}

func (h *sessionHandler) New(c *gin.Context) {
	c.HTML(http.StatusOK, "session_new", nil)
}

func (h *sessionHandler) Create(c *gin.Context) {
	var input user.LoginInput

	err := c.ShouldBind(&input)
	if err != nil {
		input.Error = input.Validate()
		c.HTML(http.StatusOK, "session_new", input)
		return
	}
	user, err := h.userService.Login(input)
	if err != nil {
		input.ErrorPassword = err
		c.HTML(http.StatusOK, "session_new", input)
		return
	}
	session := sessions.Default(c)
	session.Set("userName", user.Name)
	session.Set("userNPM", user.NPM)
	session.Set("userID", user.ID)
	session.Set("authUser", user.NPM)
	session.Save()

	c.Redirect(http.StatusFound, "/dashboard")
}

func (h *sessionHandler) NewSessionEmployee(c *gin.Context) {
	c.HTML(http.StatusOK, "session_employee_new", nil)
}

func (h *sessionHandler) CreateSessionEmployee(c *gin.Context) {
	var input employee.LoginInput

	err := c.ShouldBind(&input)
	if err != nil {
		input.Error = input.Validate()
		c.HTML(http.StatusOK, "session_employee_new", input)
		return
	}

	employee, err := h.employeeService.Login(input)
	if err != nil {
		input.ErrorPassword = err
		c.HTML(http.StatusOK, "session_employee_new", input)
		return
	}

	if employee.Role == "Admin" {
		session := sessions.Default(c)
		session.Set("authAdmin", employee.Nip)
		session.Save()

		c.Redirect(http.StatusFound, "/admin/dashboard")
	} else if employee.Role == "Wakil Dekan Bidang Akademik" {
		session := sessions.Default(c)
		session.Set("nameWadek", employee.Name)
		session.Set("authWadek", employee.Nip)
		session.Save()

		c.Redirect(http.StatusFound, "/wadek/dashboard")
	} else if employee.Role == "Koordinator Prodi Informatika" || employee.Role == "Koordinator Prodi Teknik Sipil" || employee.Role == "Koordinator Prodi Teknik Elektro" || employee.Role == "Koordinator Prodi Teknik Mesin" || employee.Role == "Koordinator Prodi Arsitektur" || employee.Role == "Koordinator Prodi Sistem Informasi" {
		session := sessions.Default(c)
		session.Set("nameKaprodi", employee.Name)
		session.Set("roleKaprodi", employee.Role)
		session.Set("authKaprodi", employee.Nip)
		session.Save()

		c.Redirect(http.StatusFound, "/koordinator-prodi/dashboard")
	} else if employee.Role == "Karyawan Akademik Bagian Informatika dan Sistem Informasi" || employee.Role == "Karyawan Akademik Bagian Teknik Sipil dan Arsitektur" || employee.Role == "Karyawan Akademik Bagian Teknik Mesin dan Teknik Elektro" {
		session := sessions.Default(c)
		session.Set("nameKaryawan", employee.Name)
		session.Set("roleKaryawan", employee.Role)
		session.Set("authKaryawan", employee.Nip)
		session.Save()

		c.Redirect(http.StatusFound, "/karyawan/dashboard")
	} else {

		c.Redirect(http.StatusFound, "/login")
	}

}

func (h *sessionHandler) Destroy(c *gin.Context) {
	session := sessions.Default(c)

	if session.Get("authUser") != nil {
		session.Clear()
		session.Options(sessions.Options{MaxAge: -1})
		session.Save()
		c.Redirect(http.StatusFound, "/")
	} else {
		session.Clear()
		session.Options(sessions.Options{MaxAge: -1})
		session.Save()
		c.Redirect(http.StatusFound, "/login")
	}

}
