package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/aayush4vedi/GoDownloader/model"
	"github.com/google/uuid"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	fmt.Println("Health:OK")
}
func generateUUID() string {
	id := uuid.New()
	return id.String()
}

func DownloadManager(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var downloadRequest model.Download
	json.Unmarshal(requestBody, &downloadRequest)
	if downloadRequest.Type == "serial" {
		for _, url := range downloadRequest.Urls {
			_ = Download(url)
		}
		downloadID := model.Response{"Id" + generateUUID()}
		w.Header().Set("Content-type", "application/json")
		id, _ := json.Marshal(downloadID)
		w.Write(id)
	}
}

func Download(url string) error {
	filepath := "/tmp" + "/" + generateUUID()
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
