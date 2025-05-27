package email_test

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"github.com/Mona-bele/rote-notify/pkg/email"
	"github.com/joho/godotenv"
	gomail "gopkg.in/mail.v2"
	"os"
	"testing"
)

func init() {
	godotenv.Load("../../../../../.env")
}

func TestNewEmail(t *testing.T) {

	emailCredential, err := email.NewEmailCredential("", "password", 587, "smtp.gmail.com")
	if err != nil {
		t.Error(err)
	}

	email, err := email.NewEmail("to", "subject", "teste", emailCredential)
	if err != nil {
		t.Error(err)
	}

	t.Log(email)
}

var body bytes.Buffer

func TestSendEmail(t *testing.T) {
	m := gomail.NewMessage()
	m.SetHeader("From", "ula@gmail.com")
	m.SetHeader("To", "ulatest@gmail.com")
	m.SetHeader("Subject", "Hello!")

	body.WriteString("<h1>Hello!</h1>")
	m.SetBody("text/html", body.String())

	d := gomail.NewDialer("smtp.gmail.com", 587, "ula@gmail.com", "xxxxxxxx")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		t.Error(err)
	}
}

func TestExampleSendEmailWithTemplate(t *testing.T) {
	m := gomail.NewMessage()

	/*if err := email.TemplateFileName(&body, "<h1>Ola Mundo</h1>"); err != nil {
		t.Error(err)
	}*/

	m.SetHeader("From", os.Getenv("EMAIL_FROM"))
	m.SetHeader("To", "pl1745240@gmail.com")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/plain", "Hello Word")

	fmt.Println(os.Getenv("EMAIL_FROM"))

	d := gomail.NewDialer(os.Getenv("EMAIL_HOST"), 465, os.Getenv("EMAIL_FROM"), os.Getenv("EMAIL_PASSWORD"))
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		t.Error(err)
	}
}

func TestSendEmailPlain(t *testing.T) {
	emailCredential, err := email.NewEmailCredential(os.Getenv("EMAIL_FROM"), os.Getenv("EMAIL_PASSWORD"), 465, os.Getenv("EMAIL_HOST"))
	if err != nil {
		t.Error(err)
	}

	email, err := email.NewEmail("pl1745240@gmail.com", "Hello!", "Hello Word", emailCredential)
	if err != nil {
		t.Error(err)
	}

	email.SendEmailPlain()
}

func TestSendEmailHtml(t *testing.T) {
	emailCredential, err := email.NewEmailCredential("ula@gmail.com", "xxxx", 587, "smtp.gmail.com")
	if err != nil {
		t.Error(err)
	}

	email, err := email.NewEmail("ulatest@gmail.com", "Hello!", "<h1>Hello Word</h1>", emailCredential)
	if err != nil {
		t.Error(err)
	}

	email.SendEmailHtml()
}

func TestSendEmailTemplateHtml(t *testing.T) {
	emailCredential, err := email.NewEmailCredential("ula@gmail.com", "xxxx", 587, "smtp.gmail.com")
	if err != nil {
		t.Error(err)
	}

	email, err := email.NewEmail("ulatest@gmail.com", "Hello!", "ula.html", emailCredential)
	if err != nil {
		t.Error(err)
	}

	email.SendEmailTemplateHtml()
}

func TestContentTypeHtml(t *testing.T) {

	contentType := email.NewContentType()
	contentType.ContentTypeBody = "text/html"

	body := email.NewContentTypeHtml()
	content := email.ContentTypeBody(body)

	if content != contentType.ContentTypeBody {
		t.Error("Content type is not html")
	}

	t.Log(content)
}
