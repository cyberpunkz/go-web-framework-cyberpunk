package mail

import (
	"github.com/go-gomail/gomail"
)

type EmailSender interface {
	Send(email Email) error
}

type Email interface {
	SetHeader(field string, value string)
	SetBody(contentType string, body string)
	GetMessage() *gomail.Message
}

type generalEmail struct {
	message *gomail.Message
}

func NewGeneralEmail() Email {
	return &generalEmail{
		message: gomail.NewMessage(),
	}
}

func (email *generalEmail) SetHeader(field string, value string) {
	email.message.SetHeader(field, value)
}

func (email *generalEmail) SetBody(contentType string, body string) {
	email.message.SetBody(contentType, body)
}

func (email *generalEmail) GetMessage() *gomail.Message {
	return email.message
}
