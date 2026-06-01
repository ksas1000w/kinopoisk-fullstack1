package handlers

import (
	"encoding/json"
	"fullstack/db"
	"fullstack/models"
	"log"
	"net/http"
)

func AddFilm(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}
	var films models.Films
	log.Println("Метод одобрен,идем дальше")
	if err := json.NewDecoder(r.Body).Decode(&films); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	tx, err := db.DB.Begin()
	if err != nil {
		log.Panicln("Begin tx:", err)
		http.Error(w, "error server", http.StatusInternalServerError)
		return
	}
	defer tx.Rollback()

	log.Println("Получиши фильмы")
	err = tx.QueryRow(`INSERT INTO trailers(path) VALUES($1) RETURNING trailer_id;`, films.Trailer.Path).Scan(&films.Trailer.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Println("Загрузили трейлеры")

	err = tx.QueryRow("INSERT INTO films(title,description,is_serial,trailer_id) VALUES($1,$2,$3,$4) RETURNING film_id;", films.Title, films.Description, films.IsSerial, films.Trailer.ID).Scan(&films.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Println("Загрузили фильмы")
	err = tx.QueryRow("INSERT INTO film_cards(film_id,path,is_horizontal) VALUES($1,$2,$3) RETURNING film_card_id;", films.ID, films.Card.Path, films.Card.IsHorizontal).Scan(&films.Card.ID)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Println("Загрузили карточки")
	err = tx.QueryRow("INSERT INTO logos(path) VALUES ($1) RETURNING logo_id;", films.Logo.Path).Scan(&films.Logo.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Println("Загрузили логотипы")
	_, err = tx.Exec("INSERT INTO logos_films(film_id,logo_id) VALUES($1,$2);", films.ID, films.Logo.ID)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Println("Загрузили связь")

	defer tx.Rollback()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]int64{"ID": int64(films.ID)})
}
