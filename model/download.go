package model

type Download struct {
	Type string   `json:"type"`
	Urls []string `json:"urls"`
}

type DownloadID struct {
	Id string
	// TODO: add other attributes
}
