package main

import (
	"errors"
	valid "github.com/asaskevich/govalidator"
	"github.com/globalsign/mgo/bson"
	"golang.org/x/crypto/bcrypt"
)

type NewUser struct {
	Name     string `valid:"required" json:"name"`
	Email    string `valid:"email,required" json:"email"`
	DNI      string `valid:"numeric,length(8|8),required" json:"dni"`
	Password string `valid:"length(4|4),required" json:"password"`
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func SignUpUser(data *NewUser) (*User, error) {

	isCorrect, err := valid.ValidateStruct(data)
	if err != nil {
		return nil, err
	}

	if !isCorrect {
		return nil, errors.New("invalid parameters for new user")
	}

	hash, err := hashPassword(data.Password)
	if err != nil {
		return nil, err
	}

	user := &User{
		Name:           data.Name,
		DNI:            data.DNI,
		PseudoPassword: hash,
		SavedVideos:    []*bson.ObjectId{},
		Email:          data.Email,
		MobilePhone:    "",
		Password:       data.Password,
		Type:           "user",
	}

	rUser, err := CreateNewUser(user)
	if err != nil {
		return nil, err
	}

	return rUser, nil
}

func CheckCredentials(dni string, password string) error {
	if valid.IsNumeric(dni) {
		user, err := GetUserByDNI(dni)
		if err != nil {
			return err
		}

		if !checkPasswordHash(password, user.PseudoPassword) {
			return errors.New("password not match")
		}

		return nil

	}

	return errors.New("invalid dni")
}

func GetListOfVideoOfUserByDni(dni string) ([]*Video, error) {
	user, err := GetUserByDNI(dni)
	if err != nil {
			return nil, err
	}
	videos := make([]*Video, 0)

	for _, v := range user.SavedVideos {
		video, err := GetVideoById(*v)
		if err != nil {
			return nil, err
		}

		videos = append(videos, video)
	}

	return videos, nil
}

