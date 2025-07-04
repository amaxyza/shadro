package services

import (
	"errors"
	"fmt"

	"github.com/amaxyza/shadro/models"
)

var users = models.NewDB()

func AddUser(name string, password string) (*models.User, error) {
	user, err := users.Add(name, password)
	if err != nil {
		return nil, errors.New("unable to create new account")
	}

	return user, nil
}

func DeleteUserByName(name string) error {
	user := users.DeleteUserFromName(name)
	if user == nil {
		return fmt.Errorf("error deleting user with name %v as it may not exist", name)
	}

	return nil
}

func DeleteUserByID(id int) error {
	user := users.DeleteUserFromID(id)
	if user == nil {
		return fmt.Errorf("error deleting user with id %v as it may not exist", id)
	}

	return nil
}

func ValidateUser(name string, password string) (bool, error) {
	return users.ValidateByName(name, password)
}

func GetUserByID(id int) (*models.User, error) {
	user := users.GetUserFromID(id)
	if user == nil {
		return user, fmt.Errorf("user with id %v does not exist", id)
	}

	return user, nil
}

func GetUserByName(name string) (*models.User, error) {
	user := users.GetUserFromName(name)
	if user == nil {
		return user, fmt.Errorf("user with name %v does not exist", name)
	}

	return user, nil
}

func GetUserList() []models.PublicUser {
	user_list := users.GetUsers()
	var public_user_list []models.PublicUser

	for _, v := range user_list {
		public_user_list = append(public_user_list, *models.Publicize(&v))
	}

	return public_user_list
}
