package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"GoDownloader/model"

	"github.com/google/uuid"
)

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
	requestBody, _ := ioutil.ReadAll(r.Body)
	var downloadRequest model.Download
	json.Unmarshal(requestBody, &downloadRequest)
	if downloadRequest.Type == "serial" {
		for _, url := range downloadRequest.Urls {
			_ = Download(url)
		}

	} else if downloadRequest.Type == "concurrent" {
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
	downloadID := model.DownloadID{"Id" + GenerateUUID()}
	w.Header().Set("Content-type", "application/json")
	id, _ := json.Marshal(downloadID)
	w.Write(id)
}

func Download(url string) error {
	filepath := "/Users/aayushchaturvedi/Desktop" + "/" + GenerateUUID() + ".png"
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
