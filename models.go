package main

import (
	"github.com/globalsign/mgo/bson"
	"github.com/go-bongo/bongo"
)

type User struct {
	bongo.DocumentBase `bson:",inline"`

	Type string `json:"type" bson:"type"`

	Name           string           `json:"name" bson:"name"`
	Email          string           `json:"email" bson:"email"`
	MobilePhone    string           `json:"mobile_phone" bson:"mobile_phone"`
	DNI            string           `json:"dni" bson:"dni"`
	PseudoPassword string           `json:"pseudo_password" bson:"pseudo_password"`
	Password       string           `json:"password" bson:"password"`
	SavedVideos    []*bson.ObjectId `json:"saved_videos" bson:"saved_videos"`
}

type Video struct {
	bongo.DocumentBase `bson:",inline"`

	OwnerID bson.ObjectId `json:"owner_id" bson:"owner_id"`
	Url     string        `json:"url" bson:"url"`

	Title       string `json:"title" bson:"title"`
	Description string `json:"description" bson:"description"`

	Location string `json:"location" bson:"location"`

	Latitude  string `json:"latitude" bson:"latitude"`
	Longitude string `json:"longitude" bson:"longitude"`

	State int `json:"state" bson:"state"`

	Thumbnail []byte `json:"thumbnail" bson:"thumbnail"`
}
