package models

import (
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// database storing users, where the key is the username and the value is the user type
type Database interface {
	GetUserFromID(id int) *User
	GetUserFromName(name string) *User
	Add(name string, password_raw string) (*User, error)
	DeleteUserFromID(id int) *User
	DeleteUserFromName(name string) *User
	Validate(id int, password string) (bool, error)
	ValidateByName(name string, password string) (bool, error)
}

type database struct {
	db         map[int]User
	name_to_id map[string]int
	next_id    int
}

func NewDB() Database {
	return &database{
		db:         make(map[int]User),
		name_to_id: make(map[string]int),
		next_id:    1,
	}
}

func (d *database) GetUserFromID(id int) *User {
	user, ok := d.db[id]
	if !ok {
		return nil
	}

	return &user
}

func (d *database) GetUserFromName(name string) *User {
	id, ok := d.name_to_id[name]
	if !ok {
		return nil
	}

	// both maps are synced, so no need for success check
	user := d.db[id]
	return &user
}

func (d *database) Add(name string, password_raw string) (*User, error) {
	password_encrypted, err := bcrypt.GenerateFromPassword([]byte(password_raw), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Error encrypting password of new user.")
		return nil, err
	}

	user := User{
		Name:          name,
		ID:            d.next_id,
		Password_Hash: string(password_encrypted),
	}

	d.db[d.next_id] = user
	d.name_to_id[name] = d.next_id

	d.next_id++

	return &user, nil
}

func (d *database) DeleteUserFromID(id int) *User {
	user, ok := d.db[id]
	if !ok {
		fmt.Println("Delete error: user does not exist.")
		return nil
	}

	delete(d.name_to_id, user.Name)
	delete(d.db, id)

	return &user
}

func (d *database) DeleteUserFromName(name string) *User {
	id, ok := d.name_to_id[name]
	if !ok {
		return nil
	}

	user := d.db[id]

	delete(d.name_to_id, user.Name)
	delete(d.db, id)

	return &user
}

func (d *database) Validate(id int, password string) (bool, error) {
	user, ok := d.db[id]
	if !ok {
		return false, errors.New("Validate Error: User does not exist")
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password_Hash), []byte(password))
	if err != nil {
		return false, errors.New("failed password check")
	}

	return true, nil
}

func (d *database) ValidateByName(name string, password string) (bool, error) {
	return d.Validate(d.name_to_id[name], password)
}
