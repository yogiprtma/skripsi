package document_legalization

import (
	"errors"
	"project-skripsi/user"
	"time"
)

type Service interface {
	CreateAppDocumentLegalization(input FormCreateAppDocumentLegalizationInput) (DocumentLegalization, error)
	CheckUserIDWhereStatusInProcess(userID int) (bool, error)
	GetAllDocumentLegalizationsWhereStatusUnsigned() ([]DocumentLegalization, error)
	GetAllDocumentLegalizationsByUserIDWhereStatusSigned(userID int) ([]DocumentLegalization, error)
	GetDocumentLegalizationByID(ID int) (DocumentLegalization, error)
	UpdateDocumentLegalization(input FormUpdateDocumentLegalizationInput, fileLocation, signature string) (DocumentLegalization, error)
	GetMessageDigest(msgDigest string) (DocumentLegalization, error)
	CheckMessageDigest(msgDigest string) (bool, error)
	GetAllDocForKaryawan(firstDepart, secondDepart string) ([]DocumentLegalization, error)
	UpdateDocumentToApprovedByKaryawan(input FormUpdateDocToApprovedByKaryawanInput) (DocumentLegalization, error)
	UpdateDocumentToRejected(ID int) (DocumentLegalization, error)
	GetAllDocForKaprodi(department string) ([]DocumentLegalization, error)
	UpdateDocumentToApprovedByKaprodi(ID int) (DocumentLegalization, error)
	GetAllDocForWadek() ([]DocumentLegalization, error)
	UpdateDocumentToSignedByWadek(input FormUpdateDocToSignedByWadekInput) (DocumentLegalization, error)
	GetDocumentLegalizationByUUID(UUID string) (DocumentLegalization, error)
	GetAllDocumentLegalizationsWhereStatusSigned() ([]DocumentLegalization, error)
}

type service struct {
	repository     Repository
	userRepository user.Repository
}

func NewService(repository Repository, userRepository user.Repository) *service {
	return &service{repository, userRepository}
}

func (s *service) CreateAppDocumentLegalization(input FormCreateAppDocumentLegalizationInput) (DocumentLegalization, error) {
	documentLegalization := DocumentLegalization{}
	documentLegalization.Email = input.Email
	documentLegalization.UserID = input.UserID
	documentLegalization.SubjectID = input.SubjectID
	documentLegalization.Status = "UNSIGNED"

	newDocumentLegalization, err := s.repository.Save(documentLegalization)
	if err != nil {
		return newDocumentLegalization, err
	}

	return newDocumentLegalization, nil
}

func (s *service) CheckUserIDWhereStatusInProcess(userID int) (bool, error) {
	documentLegalization, err := s.repository.FindUserIDWhereStatusInProcess(userID)
	if err != nil {
		return false, err
	}

	if documentLegalization.ID == 0 {
		return false, nil
	}

	return true, nil
}

func (s *service) GetAllDocumentLegalizationsWhereStatusUnsigned() ([]DocumentLegalization, error) {
	documentLegalizations, err := s.repository.FindAllWhereStatusUnsigned()
	if err != nil {
		return documentLegalizations, err
	}

	return documentLegalizations, nil
}

func (s *service) GetAllDocumentLegalizationsByUserIDWhereStatusSigned(userID int) ([]DocumentLegalization, error) {
	documentLegalizations, err := s.repository.FindAllByUserIDWhereStatusSigned(userID)
	if err != nil {
		return documentLegalizations, err
	}

	return documentLegalizations, nil
}

func (s *service) GetDocumentLegalizationByID(ID int) (DocumentLegalization, error) {
	documentLegalization, err := s.repository.FindByID(ID)
	if err != nil {
		return documentLegalization, err
	}

	if documentLegalization.ID == 0 {
		return documentLegalization, errors.New("dokumen tidak ditemukan")
	}

	return documentLegalization, nil
}

func (s *service) UpdateDocumentLegalization(input FormUpdateDocumentLegalizationInput, fileLocation, signature string) (DocumentLegalization, error) {
	documentLegalization, err := s.repository.FindByID(input.ID)
	if err != nil {
		return documentLegalization, err
	}

	documentLegalization.FileNameDocument = fileLocation
	documentLegalization.Signature = signature

	updatedDocumentLegalization, err := s.repository.UpdateSignature(documentLegalization.ID, documentLegalization.FileNameDocument, documentLegalization.Signature)
	if err != nil {
		return updatedDocumentLegalization, err
	}

	return updatedDocumentLegalization, nil
}

func (s *service) GetMessageDigest(msgDigest string) (DocumentLegalization, error) {
	documentLegalization, err := s.repository.FindByMessageDigest(msgDigest)
	if err != nil {
		return documentLegalization, err
	}

	if documentLegalization.ID == 0 {
		return documentLegalization, errors.New("dokumen transkrip tidak valid")
	}

	return documentLegalization, nil
}

func (s *service) CheckMessageDigest(msgDigest string) (bool, error) {
	documentLegalization, err := s.repository.FindByMessageDigest(msgDigest)
	if err != nil {
		return false, err
	}

	if documentLegalization.ID == 0 {
		return false, nil
	}

	return true, nil
}

func (s *service) GetAllDocForKaryawan(firstDepart, secondDepart string) ([]DocumentLegalization, error) {
	users, err := s.userRepository.FindByTwoDepartments(firstDepart, secondDepart)
	if err != nil {
		return nil, err
	}

	var usersID []int
	for _, user := range users {
		usersID = append(usersID, user.ID)
	}

	documentLegalizations, err := s.repository.FindForKaryawan(usersID)
	if err != nil {
		return documentLegalizations, err
	}

	return documentLegalizations, nil
}

func (s *service) UpdateDocumentToApprovedByKaryawan(input FormUpdateDocToApprovedByKaryawanInput) (DocumentLegalization, error) {
	approvedByKaryawanAt := time.Now()
	updatedDocumentLegalization, err := s.repository.UpdateStatusToApprovedByKaryawan(input.ID, input.FileNameDocument, input.UUID, approvedByKaryawanAt)
	if err != nil {
		return updatedDocumentLegalization, err
	}

	return updatedDocumentLegalization, nil
}

func (s *service) UpdateDocumentToRejected(ID int) (DocumentLegalization, error) {
	updatedDocumentLegalization, err := s.repository.UpdateStatusToRejected(ID)
	if err != nil {
		return updatedDocumentLegalization, err
	}

	return updatedDocumentLegalization, nil
}

func (s *service) GetAllDocForKaprodi(department string) ([]DocumentLegalization, error) {
	users, err := s.userRepository.FindByDepartment(department)
	if err != nil {
		return nil, err
	}

	var usersID []int
	for _, user := range users {
		usersID = append(usersID, user.ID)
	}

	documentLegalizations, err := s.repository.FindForKaprodi(usersID)
	if err != nil {
		return documentLegalizations, err
	}

	return documentLegalizations, nil
}

func (s *service) UpdateDocumentToApprovedByKaprodi(ID int) (DocumentLegalization, error) {
	approvedAt := time.Now()
	updatedDocumentLegalization, err := s.repository.UpdateStatusToApprovedByKaprodi(ID, approvedAt)
	if err != nil {
		return updatedDocumentLegalization, err
	}

	return updatedDocumentLegalization, nil
}

func (s *service) GetAllDocForWadek() ([]DocumentLegalization, error) {
	documentLegalizations, err := s.repository.FindForWadek()
	if err != nil {
		return documentLegalizations, err
	}

	return documentLegalizations, nil
}

func (s *service) UpdateDocumentToSignedByWadek(input FormUpdateDocToSignedByWadekInput) (DocumentLegalization, error) {
	updatedDocumentLegalization, err := s.repository.UpdateStatusToSignedByWadek(input.ID, input.MessageDigest, input.Signature, input.SignedAt, input.ExpiredAt)
	if err != nil {
		return updatedDocumentLegalization, err
	}

	return updatedDocumentLegalization, nil
}

func (s *service) GetDocumentLegalizationByUUID(UUID string) (DocumentLegalization, error) {
	documentLegalization, err := s.repository.FindByUUID(UUID)
	if err != nil {
		return documentLegalization, err
	}

	if documentLegalization.ID == 0 {
		return documentLegalization, errors.New("dokumen tidak ditemukan")
	}

	return documentLegalization, nil
}

func (s *service) GetAllDocumentLegalizationsWhereStatusSigned() ([]DocumentLegalization, error) {
	documentLegalizations, err := s.repository.FindAllWhereStatusSigned()
	if err != nil {
		return documentLegalizations, err
	}

	return documentLegalizations, nil
}
