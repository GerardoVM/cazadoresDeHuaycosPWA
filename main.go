package main

import (
	"log"
	//"github.com/pkg/errors"
	//"github.com/asaskevich/govalidator"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
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

type signupBody struct {
    dni int
    email string
    password int
    name string
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

    data := json.NewDecoder(req.Body)

    var requestBody signupBody

    log.Println(req.Body)

    err := data.Decode(&requestBody)
      if err != nil {
        log.Println(req.Body)
          panic(err)
      }

    log.Println(requestBody.dni)

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
