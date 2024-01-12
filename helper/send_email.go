package helper

import (
	"fmt"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"gopkg.in/gomail.v2"
)

const CONFIG_SMTP_HOST = "smtp.gmail.com"
const CONFIG_SMTP_PORT = 587

var CONFIG_SENDER_NAME string = fmt.Sprintf("Layanan Akademik Fakultas Teknik Universitas Bengkulu <%v>", os.Getenv("CONFIG_AUTH_EMAIL"))

var CONFIG_AUTH_EMAIL string = os.Getenv("CONFIG_AUTH_EMAIL")
var CONFIG_AUTH_PASSWORD string = os.Getenv("CONFIG_AUTH_PASSWORD")

func SendEmailToUserForSubmission(email, name, npm string) {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", CONFIG_SENDER_NAME)
	mailer.SetHeader("To", email)
	mailer.SetHeader("Subject", "Pengajuan Legalisasi Dokumen Transkrip Nilai")
	mailer.SetBody("text/html", fmt.Sprintf("Halo, %v (%v)<br><br>Pengajuan legalisasi transkrip nilai anda sedang diproses. Harap menunggu email balasan.<br><br>%v", name, npm, CONFIG_SENDER_NAME))

	dialer := gomail.NewDialer(
		CONFIG_SMTP_HOST,
		CONFIG_SMTP_PORT,
		CONFIG_AUTH_EMAIL,
		CONFIG_AUTH_PASSWORD,
	)

	err := dialer.DialAndSend(mailer)
	if err != nil {
		log.Println(err.Error())
	}
}

func SendEmailToUserForRejected(email, name, npm, message string) {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", CONFIG_SENDER_NAME)
	mailer.SetHeader("To", email)
	mailer.SetHeader("Subject", "Pengajuan Legalisasi Dokumen Transkrip Nilai")
	mailer.SetBody("text/html", fmt.Sprintf("Halo, %v (%v)<br><br>Pengajuan legalisasi transkrip nilai anda ditolak dengan alasan : %v<br><br>%v", name, npm, message, CONFIG_SENDER_NAME))

	dialer := gomail.NewDialer(
		CONFIG_SMTP_HOST,
		CONFIG_SMTP_PORT,
		CONFIG_AUTH_EMAIL,
		CONFIG_AUTH_PASSWORD,
	)

	err := dialer.DialAndSend(mailer)
	if err != nil {
		log.Println(err.Error())
	}
}

func SendEmailToUserForApproved(email, name, npm, filepath string) {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", CONFIG_SENDER_NAME)
	mailer.SetHeader("To", email)
	mailer.SetHeader("Subject", "Pengajuan Legalisasi Dokumen Transkrip Nilai")
	mailer.SetBody("text/html", fmt.Sprintf("Halo, %v (%v)<br><br>Pengajuan legalisasi transkrip nilai anda telah disetujui dan ditandatangani. Berikut lampiran dokumen transkrip nilai tersebut yang dapat anda unduh.<br><br>%v", name, npm, CONFIG_SENDER_NAME))
	mailer.Attach("./file_signed/" + filepath)

	dialer := gomail.NewDialer(
		CONFIG_SMTP_HOST,
		CONFIG_SMTP_PORT,
		CONFIG_AUTH_EMAIL,
		CONFIG_AUTH_PASSWORD,
	)

	err := dialer.DialAndSend(mailer)
	if err != nil {
		log.Println(err.Error())
	}
}
