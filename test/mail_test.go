package test

import (
	"cloud-disk/core/define"
	"crypto/tls"
	"net/smtp"
	"testing"

	"github.com/jordan-wright/email"
)

func TestSendMail(t *testing.T) {
	e := email.NewEmail()

	e.From = "yeahdcl<dclan19128961652@yeah.net>"
	e.To = []string{"dcl_1016@outlook.com"}
	e.Subject = "验证码发送测试"
	e.HTML = []byte("您的验证码为: <h2>testChar</h2>")
	err := e.SendWithTLS("smtp.yeah.net:465",
		smtp.PlainAuth("", "dclan19128961652@yeah.net", define.MailPassword, "smtp.yeah.net"),
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.yeah.net"})
	if err != nil {
		t.Fatal(err)
	}
}
