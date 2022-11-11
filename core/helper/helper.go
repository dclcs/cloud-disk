package helper

import (
	"cloud-disk/core/define"
	"crypto/md5"
	"fmt"

	"github.com/golang-jwt/jwt/v4"
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
