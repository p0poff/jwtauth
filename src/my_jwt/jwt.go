package my_jwt

import (
	jwt_git "github.com/dgrijalva/jwt-go"
	"time"
	"params"
	// "errors"
	"fmt"
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
			ExpiresAt: time.Now().Unix() + param.Expire,
			Issuer: "jwt",
			IssuedAt: time.Now().Unix(),
		},
	}
	token := jwt_git.NewWithClaims(jwt_git.SigningMethodHS256, claims)
	return token.SignedString(mySigningKey)
}

func (claims Claims) ParseToken(strToken string, param params.Init) (map[string]string, error) {
	token, err := jwt_git.ParseWithClaims(strToken, &Claims{}, func(token *jwt_git.Token) (interface{}, error) {
        return []byte(param.TokenKey), nil
    })

    if err != nil {
    	return map[string]string{}, err
    } 

    if claims, ok := token.Claims.(*Claims); ok && token.Valid {
        return map[string]string{"username": claims.Username}, nil
    } 

    return map[string]string{}, nil
}

func (claims Claims) IsValid(strToken string, param params.Init) (bool, error) {
	token, err := jwt_git.ParseWithClaims(strToken, &Claims{}, func(token *jwt_git.Token) (interface{}, error) {
        return []byte(param.TokenKey), nil
    })

    if err != nil {
    	return false, err
    } 

    if claims, ok := token.Claims.(*Claims); ok && token.Valid {
        fmt.Printf("%v %v", claims.Username, claims.StandardClaims.ExpiresAt)
        return true, nil
    } 

	return false, nil
}