package route

import (
	"GoDownloader/controller"

	"github.com/gorilla/mux"
)

func Route_call(router *mux.Router) {
	router.HandleFunc("/health", controller.Health)
	router.HandleFunc("/download/{id}", controller.Status).Methods("GET")
	router.HandleFunc("/downloads", controller.Downloader).Methods("GET")
}
