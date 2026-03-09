package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("Jacksom", "jacksom@golang.com", "123456")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "Jacksom", user.Name)
	assert.Equal(t, "jacksom@golang.com", user.Email)
}

func TestUser_ValidatePassword(t *testing.T) {
	user, err := NewUser("Jacksom", "jacksom@golang.com", "123456")
	assert.Nil(t, err)

	assert.True(t, user.ValidadePassword("123456"))
	assert.False(t, user.ValidadePassword("654321"))
	assert.NotEqual(t, "123456", user.Password)
}
