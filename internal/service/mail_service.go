package service

import (
	"bytes"
	"fmt"
	"net/smtp"
	"os"
	"text/template"
)

func ReportToAdmin() error {
	// Sender data.

	from := os.Getenv("EMAIL")
	password := os.Getenv("PASSWORD")

	// Receiver email address.
	to := []string{
		os.Getenv("ADMIN_EMAIL"),
	}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	t, err := template.ParseFiles("../../templates/report_to_admin/index.html")
	if err != nil {
		return err
	}
	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: This is a test subject \n%s\n\n", mimeHeaders)))

	t.Execute(&body, struct {
		Title string
	}{
		Title: "Xuất file thành công",
	})

	// Sending email.
	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, body.Bytes())
	if err != nil {
		return err
	}
	return err
}
