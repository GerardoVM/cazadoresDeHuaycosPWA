package main

import (
	"github.com/go-bongo/bongo"
	"log"
	"strings"
	"github.com/minio/minio-go"
)

var Connection *bongo.Connection
var UsersDB *bongo.Collection
var VideosDB *bongo.Collection

var CitaStorage *minio.Client


func init() {
	uri := "mongodb://<USERNAME>:<PASSWORD>@cluster0-shard-00-00-576ix." +
		"mongodb.net:27017,cluster0-shard-00-01-576ix.mongodb.net:27017," +
		"cluster0-shard-00-02-576ix.mongodb.net:27017/test?ssl=true&" +
		"replicaSet=Cluster0-shard-0&authSource=admin"

	uri = strings.Replace(uri, "<USERNAME>", "bregymr", -1)
	uri = strings.Replace(uri, "<PASSWORD>", "Ql0SIpf3WWMYLlEG", -1)

	log.Println("Connecting to DB...")
	log.Println("Using " + uri + "...")

	config := &bongo.Config{
		ConnectionString: uri,
		Database:         "cita",
	}

	var err error
	Connection, err = bongo.Connect(config)

	if err != nil {
		log.Fatal(err)
	}

	UsersDB = Connection.Collection(UsersCollection)
	VideosDB = Connection.Collection(VideosCollection)

	log.Println("Connection task successful\n")


	log.Println("Init Storage connection...")
	endpoint := "nyc3.digitaloceanspaces.com"
	accessKeyID := "SKTWHTY7CJEABRAG7U5R"
	secretAccessKey := "mw2sqO0N4pBRsrj5cU0CQ0CQ1YX7mGNSGUDpxWzNw1E"
	useSSL := true

	CitaStorage, err = minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Connection task successful\n")

}
