package handlers

import (
	"encoding/json"
	"fullstack/db"
	"fullstack/models"
	"log"
	"net/http"
)

func GetFilms(w http.ResponseWriter, r *http.Request) {
	query := `SELECT 
		f.film_id,
		f.title,
		f.is_serial,
		COALESCE(f.description, ''),
		t.trailer_id,
		t.path,
		fc.film_card_id,
		fc.film_id,
		fc.path,
		fc.is_horizontal,
		l.logo_id,
		l.path
	FROM films f
	LEFT JOIN trailers t     ON t.trailer_id = f.trailer_id
	LEFT JOIN film_cards fc  ON fc.film_id = f.film_id
	LEFT JOIN logos_films lf ON lf.film_id = f.film_id
	LEFT JOIN logos l 		 ON l.logo_id = lf.logo_id;
	`

	rows, err := db.DB.Query(query)
	if err != nil {
		log.Println("Ошибка запроса")
		return
	}
	defer rows.Close()

	var films []models.Films

	for rows.Next() {
		var f models.Films
		var t models.Trailer
		var fc models.FilmCard
		var l models.Logo

		var trailerID *int
		var trailerPath *string
		var cardID *int
		var cardFilmID *int
		var cardPath *string
		var cardIsHorizontal *bool
		var logoID *int
		var logoPath *string

		err := rows.Scan(
			&f.ID, &f.Title, &f.IsSerial, &f.Description,
			&trailerID, &trailerPath,
			&cardID, &cardFilmID, &cardPath, &cardIsHorizontal,
			&logoID, &logoPath,
		)
		if err != nil {
			log.Println("Ошибка чтения строки", err)
			continue
		}

		if trailerID != nil {
			t.ID = *trailerID
			t.Path = *trailerPath
			f.Trailer = &t
		}

		if cardID != nil {
			fc.ID = *cardID
			fc.FilmID = *cardFilmID
			fc.Path = *cardPath
			fc.IsHorizontal = *cardIsHorizontal
			f.Card = &fc
		}

		if logoID != nil {
			l.ID = *logoID
			l.Path = *logoPath
			f.Logo = &l
		}

		films = append(films, f)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(films)
}
