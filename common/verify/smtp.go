package verify

import (
	"net/smtp"

	"github.com/jordan-wright/email"
)

//smtp邮件发送
func SendMail(from, to, subject, body, pass string) error {
	e := email.NewEmail()
	e.From = from
	e.To = []string{to}
	e.Subject = subject
	e.Text = []byte(body)
	return e.Send("smtp.163.com:25", smtp.PlainAuth("", e.From, pass, "smtp.163.com"))
}
