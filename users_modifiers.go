package main

import (
	"errors"
	"github.com/globalsign/mgo/bson"
)

func CreateNewUser(user *User) (*User, error) {
	if user.SavedVideos == nil {
		user.SavedVideos = []*bson.ObjectId{}
	}
	err := UsersDB.Save(user)
	if err != nil {
		return nil, err
	}

	returnedUser := new(User)

	err = UsersDB.FindOne(bson.M{"dni": user.DNI}, returnedUser)
	if err != nil {
		return nil, err
	}

	return returnedUser, nil
}

func GetUserByDNI(dni string) (*User, error) {
	returnedUser := new(User)
	err := UsersDB.FindOne(bson.M{"dni": dni}, returnedUser)
	if err != nil {
		return nil, err
	}
	return returnedUser, nil

}

func DeleteUserByDNI(dni string) (*User, error) {
	returnedUser := new(User)
	err := UsersDB.FindOne(bson.M{"dni": dni}, returnedUser)
	if err != nil {
		return nil, err
	}

	err = UsersDB.DeleteOne(bson.M{"dni": dni})
	if err != nil {
		return nil, err
	}

	return returnedUser, nil
}

func UpdateUserByID(user *User) (*User, error) {

	if !user.Id.Valid() {
		return nil, errors.New("please, pass a valid id for user")
	}
	err := UsersDB.Save(user)
	if err != nil {
		return nil, err
	}

	returnedUser := new(User)

	err = UsersDB.FindOne(bson.M{"dni": user.DNI}, returnedUser)
	if err != nil {
		return nil, err
	}

	return returnedUser, nil

}


func AddVideoIdToUserByDNI(dni string, videoId *bson.ObjectId) (*User, error) {
	user, err := GetUserByDNI(dni)
	if err != nil {
	    return nil, err
	}

	if user.SavedVideos == nil {
		user.SavedVideos = []*bson.ObjectId{}
	}

	user.SavedVideos = append(user.SavedVideos, videoId)

	updatedUser, err := UpdateUserByID(user)
	if err != nil {
	    return nil, err
	}

	return updatedUser, nil
}