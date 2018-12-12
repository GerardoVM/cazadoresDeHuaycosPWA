package main

import (
	"log"
	//"github.com/gin-gonic/gin"
	//"github.com/gin-contrib/cors"
	//"github.com/gin-gonic/autotls"
	"net/http"
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

func HelloServer(w http.ResponseWriter, req *http.Request) {
    enableCors(&w)
    w.Header().Set("Content-Type", "text/plain")
    w.Write([]byte("This is an example server.\n"))
    // fmt.Fprintf(w, "This is an example server.\n")
    // io.WriteString(w, "This is an example server.\n")
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func main() {
    http.HandleFunc("/", HelloServer)
    err := http.ListenAndServeTLS("citapp.tk:443", "cert.pem", "key.pem", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
