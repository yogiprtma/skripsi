package document_legalization

import (
	"project-skripsi/subject"
	"project-skripsi/user"
	"time"
)

type DataDocumentLegalizationFormatter struct {
	ID                 int
	UserID             int
	SubjectID          int
	Email              string
	FileNameDocument   string
	Status             string
	Signature          string
	CreatedAt          time.Time
	UpdatedAt          time.Time
	CreatedAtString    string
	UpdatedAtString    string
	ApprovedByKaryawan string
	ApprovedByKaprodi  string
	SignedByWadek      string
	Subject            subject.Subject
	User               user.User
	No                 int
}

func FormatDataDocumentLegalization(documentLegalization DocumentLegalization) DataDocumentLegalizationFormatter {
	formatter := DataDocumentLegalizationFormatter{
		ID:                 documentLegalization.ID,
		UserID:             documentLegalization.UserID,
		SubjectID:          documentLegalization.SubjectID,
		Email:              documentLegalization.Email,
		FileNameDocument:   documentLegalization.FileNameDocument,
		Signature:          documentLegalization.Signature,
		CreatedAt:          documentLegalization.CreatedAt,
		UpdatedAt:          documentLegalization.UpdatedAt,
		Status:             documentLegalization.Status,
		Subject:            documentLegalization.Subject,
		User:               documentLegalization.User,
		CreatedAtString:    documentLegalization.CreatedAt.Format("01/02/2006"),
		UpdatedAtString:    documentLegalization.UpdatedAt.Format("01/02/2006"),
		ApprovedByKaryawan: documentLegalization.ApprovedByKaryawanAkademik,
		ApprovedByKaprodi:  documentLegalization.ApprovedByKaprodi,
		SignedByWadek:      documentLegalization.SignedByWadek,
	}

	return formatter
}

func FormatDataDocumentLegalizations(documentLegalizations []DocumentLegalization) []DataDocumentLegalizationFormatter {
	documentLegalizationsFormatter := []DataDocumentLegalizationFormatter{}

	for index, documentLegalization := range documentLegalizations {
		documentLegalizationFormatter := FormatDataDocumentLegalization(documentLegalization)
		documentLegalizationFormatter.No = index + 1

		documentLegalizationsFormatter = append(documentLegalizationsFormatter, documentLegalizationFormatter)
	}

	return documentLegalizationsFormatter
}
