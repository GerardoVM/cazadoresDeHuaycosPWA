package main

import (
	"github.com/globalsign/mgo/bson"
	"github.com/pkg/errors"
)

func CreateNewVideo(video *Video) (*Video, error) {
	if !video.OwnerID.Valid() {
		return nil, errors.New("invalid owner, please link a valid owner")
	}

	err := VideosDB.Save(video)
	if err != nil {
		return nil, err
	}

	returnedVideo := new(Video)

	err = VideosDB.FindById(video.Id, returnedVideo)
	if err != nil {
		return nil, err
	}

	return returnedVideo, nil
}

func GetVideoById(id bson.ObjectId) (*Video, error) {
	returnedVideo := new(Video)

	err := VideosDB.FindById(id, returnedVideo)
	if err != nil {
		return nil, err
	}

	return returnedVideo, nil

}

func GetAllVideos() ([]Video, error) {
	returnedVideos := make([]Video, 0)

	result := VideosDB.Find(bson.M{})

	if result.Error != nil {
		return nil, result.Error
	}

	v := &Video{}
	for result.Next(v) {
		returnedVideos = append(returnedVideos, *v)
	}

	return returnedVideos, nil

}

func DeleteVideoById(id bson.ObjectId) (*Video, error) {
	returnedVideo := new(Video)

	err := VideosDB.FindById(id, returnedVideo)
	if err != nil {
		return nil, err
	}


	err = VideosDB.DeleteDocument(returnedVideo)
	if err != nil {
		return nil, err
	}

	return returnedVideo, nil
}

func UpdateVideoById(video *Video) (*Video, error) {

	if !video.Id.Valid() {
		return nil, errors.New("please, pass a valid id for this one video")
	}
	err := VideosDB.Save(video)
	if err != nil {
		return nil, err
	}

	returnedVideo := new(Video)

	err = VideosDB.FindById(video.Id, returnedVideo)
	if err != nil {
		return nil, err
	}

	return returnedVideo, nil

}
