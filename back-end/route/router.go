package route

import (
	"encoding/json"
	"fullstack/handlers"
	"net/http"
)

type ResponseW struct {
	Status string `json:"status"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	var resp ResponseW
	resp.Status = "ok"
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
func SetupRouter() {
	upload := http.FileServer(http.Dir("/uploads"))
	http.Handle("/uploads/", http.StripPrefix("/uploads/", upload))
	http.HandleFunc("/", handlers.MainPage)
	http.HandleFunc("/watch/film/", handlers.WatchFilmPage) //плеер
	http.HandleFunc("/cinema", handlers.CinemaPage)         //кинотеатр
	http.HandleFunc("/admin", handlers.AdminPage)
	http.HandleFunc("/add", handlers.AddPage)
	http.HandleFunc("/api/add", handlers.AddFilm) //добавить кино
	http.HandleFunc("/api/releases", handlers.AddFilm)
	http.HandleFunc("/api/upload", handlers.AddFilm)
	http.HandleFunc("/api/films", handlers.GetFilms) //список всех фильмов
	// /film/:id/:seria
}
