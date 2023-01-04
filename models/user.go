package models

import (
	"goProject/database"
	"html"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username string
	Password string
}

func (u *User) SaveUser() (*User, error) {

	var err error
	err = database.DB.Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) BeforeSave() error {

	//turn password into hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)

	//remove spaces in username
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))

	return nil

}
