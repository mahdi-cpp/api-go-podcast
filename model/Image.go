package model

type Image struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Title     string `json:"title"`
	Title2    string `json:"title2"`
	AlbumName string `json:"album_name"`
	AlbumPath string `json:"album_path"`
}
