package model

import (
	"time"
)

type Download struct {
	Type string   `json:"type"`
	Urls []string `json:"urls"`
}

type Response struct {
	ID           string            `json:"id"`
	StartTime    time.Time         `json:"start_time"`
	EndTime      time.Time         `json:"end_time"`
	Status       string            `json:"status"`
	DownloadType string            `json:"download_type"`
	Files        map[string]string `json:files`
}
