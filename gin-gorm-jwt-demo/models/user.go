package models

import (
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"not null; unique" json:"username"`
	Password string `gorm:"size:255; not null" json:"password"`
}

func (u *User) SaveUser() error {
	err := DB.Create(u).Error
	return err
}

func examPassword(password, truepassword string) bool {
	return password == truepassword
}

func ExamUser(username string, password string) (string, error) {
	u := &User{}
	err := DB.Where(&User{Username: username}).Take(&u).Error
	if err != nil {
		return "", fmt.Errorf("user not found, %s", err.Error())
	}

	if !examPassword(password, u.Password) {
		return "", fmt.Errorf("password not match")
	}

	token, err := NewToken(u)
	if err != nil {
		return "", fmt.Errorf("get token with err, %s", err.Error())
	}

	return token, nil
}
