package model

type Download struct {
	Type string   `json:"type"`
	Urls []string `json:"urls"`
}

type Response struct {
	ID string `json:"id"`
	// StartTime    time.Time         `json:"start_time"`
	// EndTime      time.Time         `json:"end_time"`
	// Status       string            `json:"status"`
	// DownloadType string            `json:"download_type"`
	// Files        map[string]string `json:files`
}
type Error struct {
	InternalCode error  `json:"internal_code"`
	Message      string `json:"message"`
}
