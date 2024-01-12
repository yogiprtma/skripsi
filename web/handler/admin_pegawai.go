package handler

import (
	"net/http"
	"project-skripsi/employee"
	"strconv"

	"github.com/gin-gonic/gin"
)

type adminPegawaiHandler struct {
	employeeService employee.Service
}

func NewAdminPegawaiHandler(employeeService employee.Service) *adminPegawaiHandler {
	return &adminPegawaiHandler{employeeService}
}

func (h *adminPegawaiHandler) Index(c *gin.Context) {
	employees, err := h.employeeService.GetAllEmployees()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.HTML(http.StatusOK, "pegawai_index.html", gin.H{
		"employees": employee.FormatEmployees(employees),
	})
}

func (h *adminPegawaiHandler) New(c *gin.Context) {
	c.HTML(http.StatusOK, "pegawai_new.html", nil)
}

func (h *adminPegawaiHandler) Create(c *gin.Context) {
	var input employee.FormCreateEmployeeInput

	err := c.ShouldBind(&input)
	if err != nil {
		input.Error = input.Validate()
		c.HTML(http.StatusOK, "pegawai_new.html", input)
		return
	}

	_, err = h.employeeService.CreateEmployee(input)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.Redirect(http.StatusFound, "/admin/pegawai")
}

func (h *adminPegawaiHandler) Edit(c *gin.Context) {
	idParam := c.Param("id")

	id, _ := strconv.Atoi(idParam)

	getEmployee, err := h.employeeService.GetEmployeeByID(id)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	input := employee.FormUpdateEmployeeInput{}
	input.ID = getEmployee.ID
	input.Name = getEmployee.Name
	input.NIP = getEmployee.Nip
	input.Role = getEmployee.Role

	c.HTML(http.StatusOK, "pegawai_edit.html", input)
}

func (h *adminPegawaiHandler) Update(c *gin.Context) {
	idParam := c.Param("id")

	id, _ := strconv.Atoi(idParam)

	var input employee.FormUpdateEmployeeInput

	err := c.ShouldBind(&input)
	if err != nil {
		input.ID = id
		input.Error = input.Validate()
		c.HTML(http.StatusOK, "pegawai_edit.html", input)
		return
	}
	input.ID = id

	_, err = h.employeeService.UpdateEmployee(input)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.Redirect(http.StatusFound, "/admin/pegawai")
}

func (h *adminPegawaiHandler) EditPassword(c *gin.Context) {
	idParam := c.Param("id")

	id, _ := strconv.Atoi(idParam)

	registeredEmployee, err := h.employeeService.GetEmployeeByID(id)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	input := employee.FormUpdatePasswordEmployeeInput{}
	input.ID = registeredEmployee.ID
	input.Password = registeredEmployee.PasswordHash

	c.HTML(http.StatusOK, "pegawai_edit_password.html", input)
}

func (h *adminPegawaiHandler) UpdatePassword(c *gin.Context) {
	idParam := c.Param("id")

	id, _ := strconv.Atoi(idParam)

	var input employee.FormUpdatePasswordEmployeeInput

	err := c.ShouldBind(&input)
	if input.Password != input.NewPassword {
		input.ID = id
		input.ErrorPassword = "Password baru tidak sama."
		c.HTML(http.StatusOK, "pegawai_edit_password.html", input)
		return
	}
	if err != nil {
		input.ID = id
		input.Error = input.Validate()
		c.HTML(http.StatusOK, "pegawai_edit_password.html", input)
		return
	}
	input.ID = id

	_, err = h.employeeService.UpdatePasswordEmployee(input)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.Redirect(http.StatusFound, "/admin/pegawai")
}

func (h *adminPegawaiHandler) Delete(c *gin.Context) {
	idParam := c.Param("id")

	id, _ := strconv.Atoi(idParam)

	_, err := h.employeeService.DeleteEmployee(id)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.Redirect(http.StatusFound, "/admin/pegawai")
}
