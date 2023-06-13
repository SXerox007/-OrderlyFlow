package main

import (
	"log"

	db "github.com/SXerox007/-OrderlyFlow/base/db/postgres"
	"github.com/SXerox007/-OrderlyFlow/base/environment"
	"github.com/SXerox007/-OrderlyFlow/base/router"
	"github.com/SXerox007/-OrderlyFlow/base/router/server"
)

// init
func Init() {
	env := environment.GetEnv()
	port := environment.GetPort()
	//db.DBConnecting()
	router.InitRouter()
	setupRouter(env, port)
}

func main() {
	pg := db.InitPG()
	defer pg.Close()
	Init()
}

func setupRouter(env, port string) {
	//templateMux := router.SubRouter("/orderlyflow")
	//templateMux.HandleFunc("/test", TestHandler()).Methods("GET")

	//
	brankasMux := router.SubRouter("/{api}/{version}/order")
	brankasMux.Use(AuthMiddleware)
	//brankasMux.HandleFunc("/create", UploadFile()).Methods("POST")

	log.Println("Server serve at", env+":"+port)
	server.StartServer(port)
}
