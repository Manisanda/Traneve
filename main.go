package main

import (
	// "fmt"
	"log"
	"net/http"
	"traneve/api/controller"
	// "traneve/api/middleware"
	"github.com/gorilla/mux"
)

func main(){
    router := mux.NewRouter()
    router.Use(mux.CORSMethodMiddleware(router))
    router.HandleFunc("/api/", controller.Health).Methods("GET")

    srv := &http.Server{
	Handler: router,
	Addr: "127.0.0.1:8080",
    }

    log.Fatal(srv.ListenAndServe())
}


