package models

type Episode struct {
	Title    string `json:"title"`
	VideoURL string `json:"video-url"`
}
type Trailer struct {
	ID   int    `json:"id"`
	Path string `json:"path"`
}
type Logo struct {
	ID   int    `json:"id"`
	Path string `json:"path"`
}
type FilmCard struct {
	ID           int    `json:"id"`
	FilmID       int    `json:"film_id"`
	Path         string `json:"path"`
	IsHorizontal bool   `json:"is_horizontal"`
}
type Films struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	IsSerial    bool      `json:"is_serial"`
	Description string    `json:"description"`
	Trailer     *Trailer  `json:"trailer"`
	Card        *FilmCard `json:"card"`
	Logo        *Logo     `json:"logo"`
}
