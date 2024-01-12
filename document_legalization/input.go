package document_legalization

import (
	"math/big"
	"project-skripsi/subject"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type FormCreateAppDocumentLegalizationInput struct {
	Subjects  []subject.Subject
	UserID    int
	SubjectID int    `form:"perihal" binding:"required"`
	Email     string `form:"email" binding:"required,email"`
	Error     error
}

func (form FormCreateAppDocumentLegalizationInput) Validate() error {
	return validation.ValidateStruct(&form,
		validation.Field(&form.SubjectID, validation.Required.Error("Perihal tidak boleh kosong.")),
		validation.Field(&form.Email, validation.Required.Error("Email tidak boleh kosong."), is.Email.Error("Email harus format email.")),
	)
}

type FormUpdateDocumentLegalizationInput struct {
	ID          int
	NameUser    string
	NpmUser     string
	NameSubject string
	E           *big.Int
	N           *big.Int
	Error       error
}

type FormUpdateDocToApprovedByKaryawanInput struct {
	ID               int
	UUID             string
	FileNameDocument string
	Error            error
}

type FormFindSignatureDocumentLegalizationInput struct {
	NpmUser string `form:"npm" binding:"required"`
	Error   error
}

type FormRejectedDocumentInput struct {
	ID     int    `form:"id"`
	Reason string `form:"reason_for_rejected" binding:"required"`
}

func (form FormRejectedDocumentInput) Validate() error {
	return validation.ValidateStruct(&form,
		validation.Field(&form.Reason, validation.Required.Error("Alasan tidak boleh kosong.")),
	)
}

type FormUpdateDocToSignedByWadekInput struct {
	ID            int
	MessageDigest string
	Signature     string
	SignedAt      time.Time
	ExpiredAt     time.Time
	Error         error
}
