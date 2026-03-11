package database

import (
	"testing"

	"github.com/JacksomGuilherme/GoExpert/APIs/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateUser(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.User{})
	user, _ := entity.NewUser("Jacksom", "jacksom@golang.com", "123456")

	repository := NewUserRepository(db)

	err = repository.Create(user)
	assert.Nil(t, err)

	var userFound entity.User
	err = db.First(&userFound, "id = ?", user.ID).Error
	assert.Nil(t, err)
	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.Email, userFound.Email)
	assert.Equal(t, user.Password, userFound.Password)
	assert.NotNil(t, user.Password)
}

func TestFindByEmail(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.User{})
	user, _ := entity.NewUser("Jacksom", "jacksom@golang.com", "123456")

	repository := NewUserRepository(db)

	err = repository.Create(user)
	assert.Nil(t, err)

	userFound, err := repository.FindByEmail(user.Email)
	assert.Nil(t, err)
	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.Email, userFound.Email)
	assert.Equal(t, user.Password, userFound.Password)
	assert.NotNil(t, user.Password)
}
