package service

import (
	"bytes"
	"crypto/tls"
	"encoding/base32"
	gomail "gopkg.in/mail.v2"
	"log"
	"os"
	"text/template"
)

const (
	FrontEndURLVerify = "https://ceritanya-front-end.com/verify/"

	EmailHost = "smtp.gmail.com"
)

type subjectBody struct {
	subject string
	body    bytes.Buffer
}

func SendEmailVerification(emailSender, passSender, hashedEmailReceiver string) {
	log.Printf("[INFO] Sending email verification from %s", emailSender)

	email, err := base32.StdEncoding.DecodeString(hashedEmailReceiver)
	if err != nil {
		log.Printf("[ERROR] Failed to decode email receiver: %s", err)
		return
	}

	sb := subjectBody{
		subject: "Please verify your email",
		body:    bytes.Buffer{},
	}

	t, err := getTemplate("email-verification.html")
	if err != nil {
		log.Printf("[ERROR] Failed to get template: %s", err)
		return
	}

	err = t.Execute(&sb.body, struct {
		URL string
	}{
		URL: FrontEndURLVerify + hashedEmailReceiver,
	})

	if err != nil {
		log.Fatalf("[ERROR] Failed to execute template: %v", err)
	}

	if err := sendEmail(emailSender, passSender, string(email), sb); err != nil {
		log.Printf("[ERROR] Failed to send email: %s", err)
		return
	}

	log.Printf("[INFO] Email verification sent to %s", string(email))

}

func getTemplate(htmlFile string) (t *template.Template, err error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	wd = wd + "/internal/template/"

	t, err = template.ParseFiles(wd + htmlFile)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func sendEmail(emailSender, passSender, emailReceiver string, sb subjectBody) error {
	m := gomail.NewMessage()

	m.SetHeader("From", emailSender)
	m.SetHeader("To", emailReceiver)
	m.SetHeader("Subject", sb.subject)
	m.SetBody("text/html", string(sb.body.Bytes()))

	d := gomail.NewDialer(EmailHost, 587, emailSender, passSender)

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
