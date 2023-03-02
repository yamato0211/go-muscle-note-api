package repositories

import (
	"errors"
	"fiber-muscles/models"
	"fiber-muscles/schemas"
	"fiber-muscles/services"

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

func UserLogin(email string, password string) (t schemas.JwtTokenType, err error) {
	var u models.User
	if err = models.Psql.Where("email = ?", email).First(&u).Error; err != nil {
		err = errors.New("user not found")
		return
	}

	if err = bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password)); err != nil {
		err = errors.New("password is not correct")
		return
	}

	token, err := services.GenerateJwt(u.ID)
	if err != nil {
		return
	}

	t = schemas.JwtTokenType{Token: token}
	return
}

func GetUserByID(userID string) (u *models.User, err error) {
	if err = models.Psql.First(&u, "id = ?", userID).Error; err != nil {
		err = errors.New("this user id user not found")
		return
	}

	return
}
