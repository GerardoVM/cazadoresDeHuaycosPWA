package main

import (
	"bytes"
	"github.com/minio/minio-go"
	"github.com/gin-gonic/gin"
	"time"
	"net/url"
	"github.com/k0kubun/pp"
)

const CitaBucketName = "cita-huaicosycrecidas"

func SaveNewVideo(name string, contentType string, data []byte, c *gin.Context) error  {
	objectName := "videos/" + name

	buf := bytes.NewReader(data)

	//isDone := false
	//go func() {
	//	for !isDone {
	//		//c.SSEvent("eee", "sex sex!")
	//		c.Stream(func(w io.Writer) bool {
	//			w.Write([]byte("gogo\n"))
	//			log.Printf("upload percent: %.2f\n", float64(buf.Size() - int64(buf.Len()))/float64(buf.Size())*100)
	//			log.Printf("buf size: %d, buf len: %d\n", buf.Size(), buf.Len())
	//			time.Sleep(1*time.Second)
	//			return !isDone
	//		})
	//
	//	}
	//}()
	_, err := CitaStorage.PutObject(CitaBucketName,
		objectName,
		buf,
		int64(buf.Len()),
		minio.PutObjectOptions{
			ContentType: contentType,
		})
	//isDone = true


	if err != nil {
		return err
	}

	return nil
}

func GetPresignedVideoUrl(name string) (string, error){
	objectName := "videos/" + name

	uri, err := CitaStorage.PresignedGetObject(CitaBucketName, objectName, 1*time.Hour, url.Values{})
	if err != nil {
		return "", err
	}

	pp.Println(uri)

	return uri.Scheme + "://" + uri.Host + uri.RequestURI(), nil
}