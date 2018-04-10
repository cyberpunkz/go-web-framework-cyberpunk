package mail

import (
	"crypto/tls"

	"github.com/go-gomail/gomail"
)

type SMTPSender struct {
	dialer *gomail.Dialer
}

func NewSMTPSender(
	host string,
	port int,
	user string,
	pass string,
	skipVerify bool,
) EmailSender {
	dialer := gomail.NewPlainDialer(host, port, user, pass)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: skipVerify}

	return &SMTPSender{
		dialer: dialer,
	}
}

func (emailSender SMTPSender) Send(email Email) error {
	return emailSender.dialer.DialAndSend(
		email.GetMessage(),
	)
}
