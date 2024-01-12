package subject

import "errors"

type Service interface {
	CreateSubject(input FormCreateSubjectInput) (Subject, error)
	GetAllSubjects() ([]Subject, error)
	GetSubjectByID(ID int) (Subject, error)
	UpdateSubject(input FormUpdateSubjectInput) (Subject, error)
	DeleteSubject(ID int) (Subject, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateSubject(input FormCreateSubjectInput) (Subject, error) {
	subject := Subject{}
	subject.Name = input.Name

	subject, err := s.repository.Save(subject)
	if err != nil {
		return subject, err
	}

	return subject, nil
}

func (s *service) GetAllSubjects() ([]Subject, error) {
	subjects, err := s.repository.FindAll()
	if err != nil {
		return subjects, err
	}

	return subjects, nil
}

func (s *service) GetSubjectByID(ID int) (Subject, error) {
	subject, err := s.repository.FindByID(ID)
	if err != nil {
		return subject, err
	}

	if subject.ID == 0 {
		return subject, errors.New("perihal tidak ditemukan")
	}

	return subject, nil
}

func (s *service) UpdateSubject(input FormUpdateSubjectInput) (Subject, error) {
	subject, err := s.repository.FindByID(input.ID)
	if err != nil {
		return subject, err
	}

	subject.Name = input.Name

	updatedSubject, err := s.repository.Update(subject)
	if err != nil {
		return updatedSubject, err
	}

	return updatedSubject, nil
}

func (s *service) DeleteSubject(ID int) (Subject, error) {
	subject, err := s.repository.FindByID(ID)
	if err != nil {
		return subject, err
	}

	deletedSubject, err := s.repository.DeleteByID(subject.ID)
	if err != nil {
		return deletedSubject, err
	}

	return deletedSubject, nil
}
