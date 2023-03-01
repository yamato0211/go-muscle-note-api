package repositories

import (
	"errors"
	"fiber-muscles/models"

	"golang.org/x/crypto/bcrypt"
)

func CreateUser(name string, email string, password string) (u models.User, err error) {
	if err = models.Psql.Where("email = ?", email).First(&models.User{}).Error; err == nil {
		err = errors.New("this email user is already existed")
		return
	}
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), 10)
	u = models.User{Name: name, Email: email, PasswordHash: string(hashed)}
	if err = models.Psql.Create(&u).Error; err != nil {
		err = errors.New("failed create user")
		return
	}

	return
}
