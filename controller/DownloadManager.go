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
	fmt.Fprint(w, "Health: OK")
}

func GenerateUUID() string {
	id := uuid.New()
	return id.String()
}

const limitThreads = 5

func Downloader(w http.ResponseWriter, r *http.Request) {
	var mapp model.Response
	// M := make(ResponseMap)
	requestBody, _ := ioutil.ReadAll(r.Body)
	var downloadRequest model.Download
	json.Unmarshal(requestBody, &downloadRequest)
	downloadID := model.DownloadID{"Id" + GenerateUUID()}
	mapp.ID = downloadID.ID
	mapp.Status = "QUEUED"
	mapp.StartTime = time.Now()

	if downloadRequest.Type == "serial" {
		for _, url := range downloadRequest.Urls {
			_ = Download(url)
		}
		mapp.DownloadType = "serial"
		mapp.Status = "SUCCESSFUL"
		mapp.EndTime = time.Now()
		mapp.Files = F
	} else if downloadRequest.Type == "concurrent" {
		mapp.DownloadType = "concurrent"
		var ch = make(chan string)
		for i := 0; i < limitThreads; i++ {
			go func() {
				for {
					url, ok := <-ch
					if !ok {
						mapp.Status = "SUCCESSFUL"
						mapp.EndTime = time.Now()
						mapp.Files = F
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
	// mapp.Status = "SUCCESSFUL"
	// mapp.EndTime = time.Now()
	// mapp.Files = F
	w.Header().Set("Content-type", "application/json")
	id, _ := json.Marshal(downloadID)
	// fmt.Println("map: ", mapp)
	M[downloadID.ID] = mapp
	w.Write(id)
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
func Status(w http.ResponseWriter, r *http.Request) {
	id := (mux.Vars(r)["id"])
	mapp, _ := json.Marshal(M[id])
	w.Write(mapp)
}
