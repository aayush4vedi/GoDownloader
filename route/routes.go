package route

import (
	"net/http"

	"github.com/aayush4vedi/GoDownloader/controller"
)

func RouteRequest(server *http.ServeMux) {
	server.HandleFunc("/health", controller.homePage)
	server.HandleFunc("/downloads", controller.DownloadManager)
	// TODO: add the remaining routes
}
