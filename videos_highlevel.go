package main

import (
	"log"
)

func GetVideoFromBaseData(
	title string,
	latitude string,
	longitude string,
	location string,
	url string,
	thumb []byte,
) *Video {
	return &Video{
		Title:       title,
		State:       0,
		Description: "",
		Latitude:    latitude,
		Longitude:   longitude,
		Location:    location,
		Url:         url,
		Thumbnail: thumb,
	}

}

func CreateAndLinkVideoWithUser(video *Video, user *User) error {
	video.OwnerID = user.Id

	if !video.Id.Valid() {
		var err error

		video, err = CreateNewVideo(video)
		if err != nil {
			return err
		}
	}

	_, err := AddVideoIdToUserByDNI(user.DNI, &video.Id)
	if err != nil {
		return err
	}

	return nil
}

func CreateAndLinkVideoWithUserByDNI(video *Video, userDni string) error {
	user, err := GetUserByDNI(userDni)
	if err != nil {
	    return err
	}
	log.Println(user.DNI)
	err = CreateAndLinkVideoWithUser(video, user)
	return err
}