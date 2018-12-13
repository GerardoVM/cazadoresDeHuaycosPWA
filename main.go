package main

import (
	"log"
	//"github.com/pkg/errors"
	//"github.com/asaskevich/govalidator"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	//"io/ioutil"
	//"fmt"
)

//func main() {

	//r := gin.Default()
	//r.Use(cors.New(cors.Config{
		//AllowMethods:     []string{"PUT", "POST", "OPTIONS"},
		//AllowHeaders:     []string{"Origin", "Content-Type", "content-type"},
		//ExposeHeaders:    []string{"Content-Length"},
		//AllowCredentials: true,
		//AllowAllOrigins: true,
	//}))
	//LinkVideosApi(r)
	//LinkAdminAPI(r)
	//LinkOpenAPI(r)
	//LinkAuthJWT(r) // Making the auth context, all above this will be restricted
	//LinkUsersHelper(r)
	//log.Fatal(autotls.Run(r,"citapp.tk"))
//}

type SignupBody struct {
 	Dni   string
 	Email string
 	Password   string
 	Name   string
 }

 type User struct {
 	Username   string
 	Password   string `json:"-"`
 	IsAdmin bool
 	CreatedAt time.Time
 }

func HomeServer(w http.ResponseWriter, req *http.Request) {
    enableCors(&w)
    w.Header().Set("Content-Type", "text/plain")
    w.Write([]byte("You are right.\n"))
    // fmt.Fprintf(w, "This is an example server.\n")
    // io.WriteString(w, "This is an example server.\n")
}

func signUpServer(w http.ResponseWriter, req *http.Request) {

    enableCors(&w)

    user := User{}

    err := json.NewDecoder(r.Body).Decode(&user)

    if err != nil{
        panic(err)
    }

    user.CreatedAt = time.Now().Local()

    userJson, err := json.Marshal(user)

    if err != nil{
        panic(err)
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)

    w.Write(userJson)


    //Works in postman
    //w.Header().Set("content-type", "application/json")

    //bod := SignupBody{}

    //err := json.NewDecoder(req.Body).Decode(&bod)

    //if err != nil{
        //panic(err)
    //}

    //bodJson, err := json.Marshal(bod)

    //if err != nil{
        //panic(err)
    //}

    //w.Write(bodJson)

    //profile := Profile{"Alex", []string{"snowboarding", "programming"}}

      //js, err := json.Marshal(profile)
      //if err != nil {
        //http.Error(w, err.Error(), http.StatusInternalServerError)
        //return
      //}

    //w.Header().Set("Content-Type", "application/json")
    //w.Write(js)

}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func main() {

    r := mux.NewRouter()
    r.HandleFunc("/", HomeServer)
    r.HandleFunc("/signup", signUpServer)

    http.Handle("/", r)
    err := http.ListenAndServeTLS("citapp.tk:443", "cert.pem", "key.pem", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
