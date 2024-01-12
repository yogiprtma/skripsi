package document_legalization

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Repository interface {
	Save(documentLegalization DocumentLegalization) (DocumentLegalization, error)
	FindUserIDWhereStatusInProcess(userID int) (DocumentLegalization, error)
	FindAllWhereStatusUnsigned() ([]DocumentLegalization, error)
	FindAllWhereStatusSigned() ([]DocumentLegalization, error)
	FindAllByUserIDWhereStatusSigned(userID int) ([]DocumentLegalization, error)
	FindByID(ID int) (DocumentLegalization, error)
	UpdateSignature(ID int, fileNameDocument, signature string) (DocumentLegalization, error)
	FindByMessageDigest(msgDigest string) (DocumentLegalization, error)
	FindForKaryawan(usersID []int) ([]DocumentLegalization, error)
	UpdateStatusToApprovedByKaryawan(ID int, fileNameDocument, uuid string, approvedAt time.Time) (DocumentLegalization, error)
	UpdateStatusToRejected(ID int) (DocumentLegalization, error)
	FindForKaprodi(usersID []int) ([]DocumentLegalization, error)
	UpdateStatusToApprovedByKaprodi(ID int, approvedAt time.Time) (DocumentLegalization, error)
	FindForWadek() ([]DocumentLegalization, error)
	UpdateStatusToSignedByWadek(ID int, msgDigest, signature string, signedAt, expiredAt time.Time) (DocumentLegalization, error)
	FindByUUID(UUID string) (DocumentLegalization, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(documentLegalization DocumentLegalization) (DocumentLegalization, error) {
	err := r.db.Omit("ApprovedByKaryawanAkademikAt", "ApprovedByKaprodiAt", "SignedByWadekAt", "ExpiredAt").Create(&documentLegalization).Error
	if err != nil {
		return documentLegalization, err
	}

	return documentLegalization, nil
}

func (r *repository) FindUserIDWhereStatusInProcess(userID int) (DocumentLegalization, error) {
	var documentLegalization DocumentLegalization

	err := r.db.Where("user_id = ? AND status IN ?", userID, []string{"UNSIGNED", "APPROVED_BY_KARYAWAN", "APPROVED_BY_KAPRODI"}).Find(&documentLegalization).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return documentLegalization, nil
	}

	return documentLegalization, nil
}

func (r *repository) FindAllWhereStatusUnsigned() ([]DocumentLegalization, error) {
	var documentLegalizations []DocumentLegalization

	err := r.db.Preload("User").Preload("Subject").Where("status = ?", "UNSIGNED").Find(&documentLegalizations).Error
	if err != nil {
		return documentLegalizations, err
	}

	return documentLegalizations, nil
}

func (r *repository) FindAllWhereStatusSigned() ([]DocumentLegalization, error) {
	var documentLegalizations []DocumentLegalization

	err := r.db.Preload("User").Preload("Subject").Where("status = ?", "SIGNED").Find(&documentLegalizations).Error
	if err != nil {
		return documentLegalizations, err
	}

	return documentLegalizations, nil
}

func (r *repository) FindAllByUserIDWhereStatusSigned(userID int) ([]DocumentLegalization, error) {
	var documentLegalizations []DocumentLegalization

	err := r.db.Preload("User").Preload("Subject").Where("user_id = ? AND status = ?", userID, "SIGNED").Find(&documentLegalizations).Error
	if err != nil {
		return documentLegalizations, err
	}

	return documentLegalizations, nil
}

func (r *repository) FindByID(ID int) (DocumentLegalization, error) {
	var documentLegalization DocumentLegalization

	err := r.db.Preload("User").Preload("Subject").Where("id = ?", ID).Find(&documentLegalization).Error
	if err != nil {
		return documentLegalization, err
	}

	return documentLegalization, nil
}

func (r *repository) UpdateSignature(ID int, fileNameDocument, signature string) (DocumentLegalization, error) {
	var documentLegalization DocumentLegalization
	err := r.db.Model(&documentLegalization).Where("id = ?", ID).Updates(DocumentLegalization{FileNameDocument: fileNameDocument, Status: "SIGNED", Signature: signature}).Error
	if err != nil {
		return documentLegalization, err
	}

	return documentLegalization, nil
}

func (r *repository) FindByMessageDigest(msgDigest string) (DocumentLegalization, error) {
	var documentLegalization DocumentLegalization

	err := r.db.Preload("User").Preload("Subject").Where("message_digest = ?", msgDigest).First(&documentLegalization).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return documentLegalization, nil
	}

	return documentLegalization, nil
}

func (r *repository) FindForKaryawan(usersID []int) ([]DocumentLegalization, error) {
	var documentLegalizations []DocumentLegalization

	err := r.db.Where("status = ? AND user_id IN ?", "UNSIGNED", usersID).Preload("User").Preload("Subject").Find(&documentLegalizations).Error
	if err != nil {
		return documentLegalizations, err
	}

	return documentLegalizations, nil
}

func (r *repository) UpdateStatusToApprovedByKaryawan(ID int, fileNameDocument, uuid string, approvedAt time.Time) (DocumentLegalization, error) {
	var documentLegalization DocumentLegalization
	err := r.db.Model(&documentLegalization).Where("id = ?", ID).Updates(DocumentLegalization{FileNameDocument: fileNameDocument, Status: "APPROVED_BY_KARYAWAN", UUID: uuid, ApprovedByKaryawanAkademikAt: approvedAt}).Error
	if err != nil {
		return documentLegalization, err
	}

	return documentLegalization, nil
}

func (r *repository) UpdateStatusToRejected(ID int) (DocumentLegalization, error) {
	var documentLegalization DocumentLegalization
	err := r.db.Model(&documentLegalization).Where("id = ?", ID).Updates(DocumentLegalization{Status: "REJECTED"}).Error
	if err != nil {
		return documentLegalization, err
	}

	return documentLegalization, nil
}

func (r *repository) FindForKaprodi(usersID []int) ([]DocumentLegalization, error) {
	var documentLegalizations []DocumentLegalization

	err := r.db.Where("status = ? AND user_id IN ?", "APPROVED_BY_KARYAWAN", usersID).Preload("User").Preload("Subject").Find(&documentLegalizations).Error
	if err != nil {
		return documentLegalizations, err
	}

	return documentLegalizations, nil
}

func (r *repository) UpdateStatusToApprovedByKaprodi(ID int, approvedAt time.Time) (DocumentLegalization, error) {
	var documentLegalization DocumentLegalization
	err := r.db.Model(&documentLegalization).Where("id = ?", ID).Updates(DocumentLegalization{Status: "APPROVED_BY_KAPRODI", ApprovedByKaprodiAt: approvedAt}).Error
	if err != nil {
		return documentLegalization, err
	}

	return documentLegalization, nil
}

func (r *repository) FindForWadek() ([]DocumentLegalization, error) {
	var documentLegalizations []DocumentLegalization

	err := r.db.Where("status = ?", "APPROVED_BY_KAPRODI").Preload("User").Preload("Subject").Find(&documentLegalizations).Error
	if err != nil {
		return documentLegalizations, err
	}

	return documentLegalizations, nil
}

func (r *repository) UpdateStatusToSignedByWadek(ID int, msgDigest, signature string, signedAt, expiredAt time.Time) (DocumentLegalization, error) {
	var documentLegalization DocumentLegalization
	err := r.db.Model(&documentLegalization).Where("id = ?", ID).Updates(DocumentLegalization{Status: "SIGNED", MessageDigest: msgDigest, Signature: signature, SignedByWadekAt: signedAt, ExpiredAt: expiredAt}).Error
	if err != nil {
		return documentLegalization, err
	}

	return documentLegalization, nil
}

func (r *repository) FindByUUID(UUID string) (DocumentLegalization, error) {
	var documentLegalization DocumentLegalization

	err := r.db.Preload("User").Preload("Subject").Where("uuid = ?", UUID).Find(&documentLegalization).Error
	if err != nil {
		return documentLegalization, err
	}

	return documentLegalization, nil
}
