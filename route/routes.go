package route

import (
	"net/http"

	"GoDownloader/controller"
)

func RouteRequest(server *http.ServeMux) {
	server.HandleFunc("/health", controller.Health)
	server.HandleFunc("/downloads", controller.Downloader)
	// TODO: add the remaining routes
}
