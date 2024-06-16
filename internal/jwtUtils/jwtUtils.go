package jwtutils

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"os"

	jwtV5 "github.com/golang-jwt/jwt/v5"
)

type TokenPayload struct {
	Email  string `json:"email"`
	UserId string `json:"user_id"`
	Scope  string `json:"scope"`
	Role   string `json:"role"`
}

type MyCustomClaims struct {
	jwtV5.RegisteredClaims
	TokenPayload
}

type JwtInterface interface {
}

func GenerateToken(payload *TokenPayload) (string, error) {
	pK := GetPrivateKey()
	sg := jwtV5.SigningMethodES256
	claims := MyCustomClaims{
		TokenPayload: TokenPayload{
			Email:  payload.Email,
			UserId: payload.UserId,
			Scope:  payload.UserId,
			Role:   payload.UserId,
		},
	}
	token := jwtV5.NewWithClaims(sg, claims)
	jwt, err := token.SignedString(pK)
	if err != nil {
		panic(err.Error())
	}
	return jwt, err

}

func GetPrivateKey() *ecdsa.PrivateKey {
	//esSigner := jwtV5.SigningMethodES256
	byt, err := os.ReadFile("privatekey.pem")
	if err == nil {
		var ecKey *ecdsa.PrivateKey
		key, err := x509.ParsePKCS8PrivateKey(byt)
		if err != nil {
			goto generate
		}
		ecKey = key.(*ecdsa.PrivateKey)
		return ecKey
	}
generate:
	pK, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(err.Error())
	}

	byt, err = x509.MarshalPKCS8PrivateKey(pK)
	if err != nil {
		panic(err.Error())
	}
	fl, err := os.Create("privatekey.pem")
	if err != nil {
		panic(err.Error())
	}
	_, err = fl.Write(byt)
	if err != nil {
		panic(err.Error())
	}
	return pK
}
