package main

import (
	"GoDownloader/route"
	"net/http"
)

func main() {
	server := http.NewServeMux()
	// fmt.Print("asdf")
	route.RouteRequest(server)
	_ = http.ListenAndServe(":8081", server)
}
