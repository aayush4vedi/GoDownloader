package main

import (
	"GoDownloader/route"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	route.Route_call(router)
	log.Fatal(http.ListenAndServe(":8081", router))
}
