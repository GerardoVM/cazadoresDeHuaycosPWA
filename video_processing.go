package main

import (
	"os/exec"
	"fmt"
	"bytes"
	"image/jpeg"
	"image"
	"io"
	"github.com/disintegration/imaging"
	"io/ioutil"
	"os"
)

func extractFrameOfVideo(filename string, width int, height int) (image.Image, error) {
	cmd := exec.Command(
		"ffmpeg",
		"-i",
		filename,
		"-vframes", "1", "-s",
		fmt.Sprintf("%dx%d", width, height),
		"-f", "singlejpeg",
		"-",
	)
	var buffer bytes.Buffer
	cmd.Stdout = &buffer

	err := cmd.Run()

	if err != nil {
		return nil, err
	}

	thumb, err := jpeg.Decode(&buffer)
	if err != nil {
	    return nil, err
	}

	return thumb, nil
}

func GetThumbnailFromVideo(filename string) (io.Reader, error) {
	image, err := extractFrameOfVideo(filename, 960, 1280)
	if err != nil {
		return nil, err
	}
	w := 200.0
	r := 2./3.
	newImage := imaging.Fill(image, int(w), int(w*r), imaging.Center, imaging.Lanczos)
	buf := bytes.NewBuffer([]byte{})
	err = imaging.Encode(buf, newImage,imaging.JPEG)

	if err != nil {
		return nil, err
	}

	return buf, nil

}

func GetThumbnailFromVideoBytes(data []byte) (io.Reader, error) {
	f, err := ioutil.TempFile(os.TempDir(), "citatemp")
	if err != nil {
	    return nil, err
	}

	_, err = f.Write(data)
	if err != nil {
		return nil, err
	}

	filename := f.Name()

	thumb ,err := GetThumbnailFromVideo(filename)

	if err != nil {
		return nil, err
	}

	return thumb, err
}