package route

import (
	"fmt"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	fmt.Println("Health:OK")
}
func RouteRequest(server *http.ServeMux) {
	fmt.Print("dasfs")
	//server.HandleFunc("/health", homePage)
	// server.HandleFunc("/downloads", controller.DownloadManager)
	// TODO: add the remaining routes
}
