package common_services

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHashPassword(t *testing.T) {
	hasher := AppHasher{}
	_, err := hasher.HashPassword("secret")
	assert.NoError(t, err)
}

func TestCheckPassword(t *testing.T) {
	hasher := AppHasher{}
	hashedPass, err := hasher.HashPassword("secret")
	assert.NoError(t, err)

	err = hasher.CheckPassword(hashedPass, "secret")
	assert.NoError(t, err)
}
