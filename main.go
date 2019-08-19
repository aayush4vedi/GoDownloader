package main

import (
	"https://github.com/aayush4vedi/GoDownloader/route"
	"net/http"
)

func main() {
	server := http.NewServeMux()
	route.RouteRequest(server)
	_ = http.ListenAndServe(":3000", server)
}