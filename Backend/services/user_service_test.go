package services

import (
	"Crud_fiber_Go/models"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

var user = models.User{
	ID:       1234568,
	Name:     "Rohan",
	Email:    "rohan@rohan.in",
	Password: "password",
}

func TestCreateUser(t *testing.T) {
	models.ConnectDatabase()

	err := CreateUser(&user)

	assert.NoError(t, err)
}

func TestGetUSer(t *testing.T) {
	allUser, err := GetUsers()
	assert.NoError(t, err)
	assert.NotEmpty(t, allUser)
	name := allUser[0]
	assert.Equal(t, name.Email, user.Email)

}

func TestUpdateUser(t *testing.T) {
	err := UpdateUser(user.ID, &user)
	assert.NoError(t, err)
}

func TestDeleteUser(t *testing.T) {
	id := strconv.Itoa(int(user.ID))
	err := DeleteUser(id)
	assert.NoError(t, err)

}
