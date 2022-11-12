package helper

import (
	"cloud-disk/core/define"
	"crypto/md5"
	"crypto/tls"
	"fmt"
	"net/smtp"

	"github.com/golang-jwt/jwt/v4"
	"github.com/jordan-wright/email"
)

func MD5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

func GenerateToken(id uint64, identity, name string) (string, error) {
	uc := define.UserCalim{
		Id:       id,
		Identity: identity,
		Name:     name,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)

	signedToken, error := token.SignedString([]byte(define.JWTKey))

	if error != nil {
		return "", error
	}

	return signedToken, error
}

func MailSend(mail, code string) error {
	e := email.NewEmail()

	e.From = define.MailFromName
	e.To = []string{mail}
	e.Subject = "验证码发送测试"
	e.HTML = []byte("您的验证码为: <h2>" + code + "</h2>")
	err := e.SendWithTLS("smtp.yeah.net:465",
		smtp.PlainAuth("", define.MailName, define.MailPassword, "smtp.yeah.net"),
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.yeah.net"})
	if err != nil {
		return err
	}

	return nil
}
