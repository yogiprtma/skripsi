package document_legalization

import (
	"project-skripsi/subject"
	"project-skripsi/user"
	"time"
)

type DocumentLegalization struct {
	ID                         int
	UUID                       string
	UserID                     int
	SubjectID                  int
	Email                      string
	FileNameDocument           string
	Status                     string
	MessageDigest              string
	Signature                  string
	ApprovedByKaryawanAkademik string
	ApprovedByKaprodi          string
	SignedByWadek              string
	ExpiredAt                  time.Time
	CreatedAt                  time.Time
	UpdatedAt                  time.Time
	Subject                    subject.Subject
	User                       user.User
}
