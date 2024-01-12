package employee

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	CreateEmployee(input FormCreateEmployeeInput) (Employee, error)
	GetAllEmployees() ([]Employee, error)
	UpdateEmployee(input FormUpdateEmployeeInput) (Employee, error)
	GetEmployeeByID(ID int) (Employee, error)
	UpdatePasswordEmployee(input FormUpdatePasswordEmployeeInput) (Employee, error)
	DeleteEmployee(ID int) (Employee, error)
	Login(input LoginInput) (Employee, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateEmployee(input FormCreateEmployeeInput) (Employee, error) {
	employee := Employee{}
	employee.Name = input.Name
	employee.Nip = input.NIP
	employee.Role = input.Role

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return employee, err
	}
	employee.PasswordHash = string(passwordHash)

	newEmployee, err := s.repository.Save(employee)
	if err != nil {
		return newEmployee, err
	}

	return newEmployee, nil
}

func (s *service) GetAllEmployees() ([]Employee, error) {
	employees, err := s.repository.FindAll()
	if err != nil {
		return employees, err
	}

	return employees, nil
}

func (s *service) UpdateEmployee(input FormUpdateEmployeeInput) (Employee, error) {
	employee, err := s.repository.FindByID(input.ID)
	if err != nil {
		return employee, err
	}

	employee.Name = input.Name
	employee.Nip = input.NIP
	employee.Role = input.Role

	updatedEmployee, err := s.repository.Update(employee.ID, employee.Name, employee.Nip, employee.Role)
	if err != nil {
		return updatedEmployee, err
	}

	return updatedEmployee, nil
}

func (s *service) GetEmployeeByID(ID int) (Employee, error) {
	employee, err := s.repository.FindByID(ID)
	if err != nil {
		return employee, err
	}

	if employee.ID == 0 {
		return employee, errors.New("pegawai tidak ditemukan")
	}

	return employee, nil
}

func (s *service) UpdatePasswordEmployee(input FormUpdatePasswordEmployeeInput) (Employee, error) {
	employee, err := s.repository.FindByID(input.ID)
	if err != nil {
		return employee, err
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return employee, err
	}

	employee.PasswordHash = string(passwordHash)

	updatedPasswordEmployee, err := s.repository.UpdatePassword(employee.ID, employee.PasswordHash)
	if err != nil {
		return updatedPasswordEmployee, err
	}

	return updatedPasswordEmployee, nil
}

func (s *service) DeleteEmployee(ID int) (Employee, error) {
	employee, err := s.repository.FindByID(ID)
	if err != nil {
		return employee, err
	}

	deletedEmployee, err := s.repository.DeleteByID(employee.ID)
	if err != nil {
		return deletedEmployee, err
	}

	return deletedEmployee, nil
}

func (s *service) Login(input LoginInput) (Employee, error) {
	nip := input.Nip
	password := input.Password

	employee, err := s.repository.FindByNIP(nip)
	if err != nil {
		return employee, err
	}

	if employee.ID == 0 {
		return employee, errors.New("pegawai tidak ditemukan")
	}

	err = bcrypt.CompareHashAndPassword([]byte(employee.PasswordHash), []byte(password))
	if err != nil {
		return employee, err
	}

	return employee, nil
}
