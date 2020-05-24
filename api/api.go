package api

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/roccoriu/restapi/api/endpoints/user"
	"github.com/roccoriu/restapi/settings"
	"log"
	"net/http"
	/*	"go.mongodb.org/mongo-driver/bson"
		"go.mongodb.org/mongo-driver/bson/primitive"*/)

func Api() {
	// init router
	setting := settings.GetSettings()
	router := mux.NewRouter()

	// set api endpoints
	user.SetUserEndpoint(router)

	// setup server and listen to port 8000
	fmt.Println("starting server...")
	fmt.Printf("Running Server on %s%s\n", setting.ServerAddr, setting.TcpPort)
	log.Fatal(http.ListenAndServe(setting.ServerAddr+setting.TcpPort, router))

}
