package passwordless

import (
	"bytes"
	"fmt"
	"log"
	"net/smtp"
	"os"
	"strings"
	"text/template"
)

var (
	auth      smtp.Auth
	fromEmail string
)

const (
	EMAIL_FROM          = ""
	EMAIL_HOST_USER     = ""
	EMAIL_HOST_PASSWORD = ""
	EMAIL_HOST          = ""
	EMAIL_PORT          = ""
)

var emailTemplate = template.Must(template.New("emailTemplate").Parse(`From: {{.From}}
To: {{.To}}
Subject: {{.Subject}}

{{.Body}}`))

func SendMail(to []string, subject, message string) error {
	var doc bytes.Buffer

	ctx := struct {
		From    string
		To      string
		Subject string
		Body    string
	}{
		fromEmail,
		strings.Join(to, ", "),
		subject,
		message,
	}

	if err := emailTemplate.Execute(&doc, ctx); err != nil {
		return err
	}

	return smtp.SendMail(
		fmt.Sprintf("%v:%v", EMAIL_HOST, EMAIL_PORT),
		auth,
		fromEmail,
		to,
		doc.Bytes())
}

func init() {
	auth = smtp.PlainAuth("", EMAIL_HOST_USER, EMAIL_HOST_PASSWORD, EMAIL_HOST)
}
