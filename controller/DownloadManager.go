package controller

import (
	"GoDownloader/model"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

var ResponseMap = make(map[string]model.Response)
var FileMap = make(map[string]string)

func Health(write http.ResponseWriter, read *http.Request) {
	write.WriteHeader(200)
	fmt.Fprint(write, "Health: OK")
}

func GenerateUUID() string {
	id := uuid.New()
	return id.String()
}

const limitThreads = 5

func Downloader(write http.ResponseWriter, read *http.Request) {
	var mapp model.Response
	requestBody, _ := ioutil.ReadAll(read.Body)
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
		mapp.Files = FileMap
	} else if downloadRequest.Type == "concurrent" {
		mapp.DownloadType = "concurrent"
		var ch = make(chan string)
		for lim := 0; lim < limitThreads; lim++ {
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
			mapp.Status = "SUCCESSFUL"
			mapp.EndTime = time.Now()
			mapp.Files = FileMap
			return
		}()
	} else {
		e := model.Error{
			InternalCode: "4001",
			Message:      "unknown type of download",
		}
		er, _ := json.Marshal(e)
		write.Write(er)
		return
	}
	write.Header().Set("Content-type", "application/json")
	id, _ := json.Marshal(downloadID)
	// fmt.Println("map: ", mapp)
	ResponseMap[downloadID.ID] = mapp
	write.Write(id)
}

func Download(url string) error {
	filepath := "/Users/aayushchaturvedi/Desktop" + "/" + GenerateUUID() + ".png"
	FileMap[url] = filepath
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
func Status(write http.ResponseWriter, read *http.Request) {
	id := (mux.Vars(read)["id"])
	if _, ok := ResponseMap[id]; ok {
		mapp, _ := json.Marshal(ResponseMap[id])
		write.Write(mapp)
	} else {
		e := model.Error{
			InternalCode: "4002",
			Message:      "unknown download ID",
		}
		er, _ := json.Marshal(e)
		write.Write(er)
	}
}
func Files(write http.ResponseWriter, read *http.Request) {
	write.Header().Set("Content-Type", "text/html; charset=utf-8")
	template, err := template.ParseFiles("template.html")
	if err != nil {
		fmt.Fprintf(write, "Unable to load template")
	}
	template.Execute(write, ResponseMap)
}
