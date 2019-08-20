package controller

import (
	"GoDownloader/model"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

var M = make(map[string]model.Response)
var F = make(map[string]string)

func Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	fmt.Println("Health:OK")
}
func GenerateUUID() string {
	id := uuid.New()
	return id.String()
}

const limitThreads = 5

func Downloader(w http.ResponseWriter, r *http.Request) {
	// M := make(ResponseMap)
	requestBody, _ := ioutil.ReadAll(r.Body)
	var downloadRequest model.Download
	json.Unmarshal(requestBody, &downloadRequest)
	downloadID := model.DownloadID{"Id" + GenerateUUID()}
	M[downloadID.ID] = model.Response{
		ID:        downloadID.ID,
		Status:    "QUEUED",
		StartTime: time.Now(),
	}
	if downloadRequest.Type == "serial" {
		for _, url := range downloadRequest.Urls {
			_ = Download(url)
		}
		M[downloadID.ID] = model.Response{
			DownloadType: "serial",
		}
	} else if downloadRequest.Type == "concurrent" {
		M[downloadID.ID] = model.Response{
			DownloadType: "concurrent",
		}
		var ch = make(chan string)
		for i := 0; i < limitThreads; i++ {
			go func() {
				for {
					url, ok := <-ch
					if !ok {
						return //channel is closed
					}
					_ = Download(url)
				}
			}()
		}
		go func() {
			for _, url := range downloadRequest.Urls {
				ch <- url
			}
			close(ch)
			return
		}()
	}
	M[downloadID.ID] = model.Response{
		Status:  "SUCCESSFUL",
		EndTime: time.Now(),
		Files:   F,
	}
	w.Header().Set("Content-type", "application/json")
	id, _ := json.Marshal(downloadID)
	fmt.Println("map: ", M[downloadID.ID])
	w.Write(id)
	// fmt.Fprint(w, M[downloadID.ID])

}
func Status(w http.ResponseWriter, r *http.Request) {
	id := (mux.Vars(r)["id"])
	fmt.Println("here")
	fmt.Fprint(w, M[id])
}
func Download(url string) error {
	filepath := "/Users/aayushchaturvedi/Desktop" + "/" + GenerateUUID() + ".png"
	F[url] = filepath
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, resp.Body)
	return err
}
