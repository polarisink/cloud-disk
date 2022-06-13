package test

import (
	"cloud-disk/core/define"
	"crypto/tls"
	"github.com/jordan-wright/email"
	"net/smtp"
	"testing"
)

func TestSendMail(t *testing.T) {
	e := email.NewEmail()
	e.From = "Get <1952482944@qq.com>"
	e.To = []string{"lqswuhan1999@gmail.com"}
	e.Subject = "Awesome Subject"
	e.HTML = []byte("rx,你的验证码为: <h1>765241</h1>")
	err := e.SendWithTLS("smtp.qq.com:465", smtp.PlainAuth("", "1952482944@qq.com", define.MailPassword, "smtp.qq.com"), &tls.Config{
		InsecureSkipVerify: true, ServerName: "smtp.qq.com"})
	if err != nil {
		t.Fatal(err)
	}
}
