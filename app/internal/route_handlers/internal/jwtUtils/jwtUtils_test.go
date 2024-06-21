package jwtutils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuth(t *testing.T) {
	_ = GetPrivateKey()
	assert.FileExists(t, "privatekey.pem")
	//ty := &ecdsa.PrivateKey{}
	//assert.IsType(t, ty, ky)

	jwt, err := GenerateToken(&TokenPayload{
		Email:  "trinitietp@gmail.com",
		UserId: "trinitietp",
	})
	fmt.Printf("%s\n", jwt)
	assert.NoError(t, err)

	payload, err := VerifyToken(jwt)
	assert.NoError(t, err)
	assert.Equal(t, "trinitietp@gmail.com", payload.Email)
}
