package user

import (
	"errors"
	"fmt"
	"project-skripsi/signature"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input FormCreateUserInput) (User, error)
	Login(input LoginInput) (User, error)
	GetAllUsers() ([]User, error)
	GetUserByID(ID int) (User, error)
	UpdateUser(input FormUpdateUserInput) (User, error)
	UpdatePasswordUser(input FormUpdatePasswordUserInput) (User, error)
	DeleteUser(ID int) (User, error)
	GetPublicAndPrivateKeyByNPM(input FormFindPublicAndPrivateKeyUserInput) (User, error)
	GetUsersByTwoDepartments(first string, second string) ([]User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(input FormCreateUserInput) (User, error) {
	user := User{}
	user.Name = input.Name
	user.NPM = input.NPM
	user.Department = input.Department

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	user.PasswordHash = string(passwordHash)

	// buat public key dan private key akun
	e, d, n := signature.GenerateKey()

	user.PublicKey = fmt.Sprintf("%v, %v", d, n)
	user.PrivateKey = fmt.Sprintf("%v, %v", e, n)

	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}
	return newUser, nil
}

func (s *service) Login(input LoginInput) (User, error) {
	npm := input.NPM
	password := input.Password

	user, err := s.repository.FindByNPM(npm)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("user tidak ditemukan")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) GetAllUsers() ([]User, error) {
	users, err := s.repository.FindAll()
	if err != nil {
		return users, err
	}

	return users, nil
}

func (s *service) GetUserByID(ID int) (User, error) {
	user, err := s.repository.FindByID(ID)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("user tidak ditemukan")
	}

	return user, nil
}

func (s *service) UpdateUser(input FormUpdateUserInput) (User, error) {
	user, err := s.repository.FindByID(input.ID)
	if err != nil {
		return user, err
	}

	user.Name = input.Name
	user.NPM = input.NPM
	user.Department = input.Department

	updatedUser, err := s.repository.Update(user)
	if err != nil {
		return updatedUser, err
	}

	return updatedUser, nil
}

func (s *service) UpdatePasswordUser(input FormUpdatePasswordUserInput) (User, error) {
	user, err := s.repository.FindByID(input.ID)
	if err != nil {
		return user, err
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	user.PasswordHash = string(passwordHash)

	updatedPasswordUser, err := s.repository.UpdatePassword(user.ID, user.PasswordHash)
	if err != nil {
		return updatedPasswordUser, err
	}

	return updatedPasswordUser, nil
}

func (s *service) DeleteUser(ID int) (User, error) {
	user, err := s.repository.FindByID(ID)
	if err != nil {
		return user, err
	}

	deletedUser, err := s.repository.DeleteByID(user.ID)
	if err != nil {
		return deletedUser, err
	}

	return deletedUser, nil
}

func (s *service) GetPublicAndPrivateKeyByNPM(input FormFindPublicAndPrivateKeyUserInput) (User, error) {
	user, err := s.repository.FindByNPM(input.NPM)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("NPM tidak ditemukan")
	}

	return user, nil
}

func (s *service) GetUsersByTwoDepartments(first string, second string) ([]User, error) {
	users, err := s.repository.FindByTwoDepartments(first, second)
	if err != nil {
		return users, err
	}

	return users, nil
}
