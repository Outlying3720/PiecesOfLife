package models

import (
	"fmt"
	"math/rand"

	"gorm.io/gorm"
)

type Token struct {
	gorm.Model
	Token  string `gorm:"not null; unique" json:"token"`
	UserID uint   `gorm:"not null" json:"userid"`
	User   User   `gorm:"foreignKey:UserID;references:ID"`
}

func RandomAToken() string {
	rand := rand.Intn(32)
	return fmt.Sprintf("token-%d", rand)
}

func NewToken(user *User) (string, error) {
	newtoken := RandomAToken()
	token := Token{
		Token: newtoken,
		User:  *user,
	}

	fmt.Println(*user)
	fmt.Println(token)
	err := DB.Create(&token).Error
	if err != nil {
		return "", err
	}

	return newtoken, nil
}

func ExamToken(token string) (*User, error) {
	fmt.Println(token)
	found := &Token{}
	err := DB.Debug().Preload("User").Model(&Token{}).Where("token = ?", token).First(&found).Error
	if err != nil {
		return nil, err
	}

	fmt.Println(found)

	return &found.User, nil
}
