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

/////////////////////////////
// package models
// import (
//     “io”
//     “net/http”
//     “os”
//     “github.com/gin-gonic/gin”
// )
// type Concurrent struct {
//     Urls []string
// }
// const limitThreads = 5
// func (files *Concurrent) Download(r *gin.Context) {
//     var ch = make(chan string)
//     for i := 0; i < limitThreads; i++ {
//         go func() {
//             for {
//                 url, ok := <-ch
//                 if !ok {
//                     return //close go routine when channel is closed
//                 }
//                 resp, _ := http.Get(url)
//                 nameId := uuid()
//                 out, _ := os.Create("/Users/punitlakshwani/go/src/github.com/punit1997/DownloadManager/downloads/" + nameId)
//                 io.Copy(out, resp.Body)
//                 resp.Body.Close()
//             }
//         }()
//     }
//     go func() {
//         for _, url := range files.Urls {
//             ch <- url
//         }
//         close(ch)
//         return
//     }()
//     ResponseId := uuid()
//     r.JSON(200, gin.H{
//         “Id”: ResponseId,
//     })
// }
