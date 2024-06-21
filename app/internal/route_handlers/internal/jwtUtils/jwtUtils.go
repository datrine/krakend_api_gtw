package jwtutils

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"os"
	"time"

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
	ttl := time.Now().Add(time.Hour * 24 * 30)
	claims := MyCustomClaims{
		TokenPayload: TokenPayload{
			Email:  payload.Email,
			UserId: payload.UserId,
			Scope:  payload.UserId,
			Role:   payload.UserId,
		}, RegisteredClaims: jwtV5.RegisteredClaims{
			ExpiresAt: jwtV5.NewNumericDate(ttl),
		},
	}
	token := jwtV5.NewWithClaims(sg, claims)
	jwt, err := token.SignedString(pK)
	if err != nil {
		panic(err.Error())
	}
	return jwt, err

}

func VerifyToken(tk string) (*TokenPayload, error) {
	payload := &TokenPayload{}
	var cl *MyCustomClaims
	var err error
	token, err := jwtV5.ParseWithClaims(tk, &MyCustomClaims{}, func(t *jwtV5.Token) (interface{}, error) {
		pK := GetPrivateKey()
		return &pK.PublicKey, nil
	})
	if err != nil {
		return nil, err
	}
	cl = token.Claims.(*MyCustomClaims)
	payload.Email = cl.Email
	payload.Role = cl.Role
	payload.Scope = cl.Scope
	payload.UserId = cl.UserId
	return payload, nil
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
