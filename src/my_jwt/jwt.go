package my_jwt

import (
	jwt_git "github.com/dgrijalva/jwt-go"
	"time"
	"params"
)

type Claims struct {
	Username string
	jwt_git.StandardClaims
}

func (claims Claims) GenToken(username string, param params.Init) (string, error) {
	mySigningKey := []byte(param.TokenKey)
	claims = Claims{
		username,
		jwt_git.StandardClaims{
			ExpiresAt: time.Now().Add(param.Expire).Unix(),
			Issuer: "jwt",
			IssuedAt: time.Now().Unix(),
		},
	}
	token := jwt_git.NewWithClaims(jwt_git.SigningMethodHS256, claims)
	return token.SignedString(mySigningKey)
}