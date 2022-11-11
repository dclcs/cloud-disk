package define

import "github.com/golang-jwt/jwt/v4"

type UserCalim struct {
	Id       uint64
	Identity string
	Name     string
	jwt.StandardClaims
}

var JWTKey = "cloud-disk-key"
