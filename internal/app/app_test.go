package app

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInitApp(t *testing.T) {
	err := InitApp()
	assert.NoError(t, err)
}
