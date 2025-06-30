package models

type User struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Password_Hash string `json:"password_hash"`
}

type PublicUser struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func Publicize(user *User) *PublicUser {
	return &PublicUser{
		ID:   user.ID,
		Name: user.Name,
	}
}
