package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUe(t *testing.T) {
	user, err := NewUser("Fillipi Nascimento", "fillipi@gmail.com", "mudar123")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "Fillipi Nascimento", user.Name)
	assert.Equal(t, "fillipi@gmail.com", user.Email)
}

func TestUserValidatePassword(t *testing.T) {
	user, _ := NewUser("Fillipi Nascimento", "fillipi@gmail.com", "mudar123")
	assert.True(t, user.ValidatePassword("mudar123"))
	assert.False(t, user.ValidatePassword("mudar1234"))
}
