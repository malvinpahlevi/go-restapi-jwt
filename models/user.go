package models

import (
	"errors"
	"go-restapi-jwt/utils/token"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"html"
	"strings"
)

type User struct {
	gorm.Model
	Username string `gorm:"size:255;not null; unique" json:"username"`
	Password string `gorm:"size:255;not null;" json:"password"`
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginCheck(username string, password string) (string, error) {

	var err error

	u := &User{}

	err = DB.Model(&User{}).Where("username = ?", username).Take(&u).Error
	if err != nil {
		return "", err
	}

	err = VerifyPassword(u.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := token.GenerateToken(u.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (u *User) SaveUser() (*User, error) {

	var err error

	err = DB.Create(&u).Error
	if err != nil {
		return &User{}, err
	}

	return u, nil
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {

	pw, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(pw)
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))

	return
}

func GetUserById(uid uint) (User, error) {

	var user User
	if err := DB.First(&user, uid).Error; err != nil {
		return user, errors.New("User not found!")
	}

	user.PrepareGive()

	return user, nil
}

func (u *User) PrepareGive() {
	u.Password = ""
}
