package main

import (
	"net/http"

	"github.com/aayush4vedi/GoDownloader/route"
)

func main() {
	server := http.NewServeMux()
	route.RouteRequest(server)
	_ = http.ListenAndServe(":3000", server)
}
