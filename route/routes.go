package route

import (
	"net/http"

	"github.com/aayush4vedi/GoDownloader/controller"
)

func RouteRequest(server *http.ServeMux) {
	server.HandleFunc("/health", DownloadManager.Health)
	server.HandleFunc("/downloads", controller.Downloader)
	// TODO: add the remaining routes
}
