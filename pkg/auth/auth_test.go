package auth

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidateToken(t *testing.T) {
	jwtWrapper := TokenManager{
		SecretKey:       "verySecret",
		Issuer:          "AuthService",
		ExpirationHours: 24,
	}

	generatedToken, err := jwtWrapper.GenerateToken("jwt@email.com", 1)
	assert.NoError(t, err)

	claims, err := jwtWrapper.ValidateToken(generatedToken)
	assert.NoError(t, err)

	assert.Equal(t, "jwt@email.com", claims.Email)
	assert.Equal(t, "AuthService", claims.Issuer)
}
