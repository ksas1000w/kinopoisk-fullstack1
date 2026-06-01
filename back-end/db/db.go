package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Init() {
	var err error

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Ошибка подключения к БД", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Бд не отвечает", err)
	}

	log.Println("Подключение к БД успешно")
	migrate()
}

func migrate() {
	queries := []string{

		`CREATE TABLE IF NOT EXISTS trailers (
			trailer_id SERIAL PRIMARY KEY,
			path TEXT NOT NULL
		);`,
		`CREATE TABLE IF NOT EXISTS logos (
			logo_id SERIAL PRIMARY KEY,	
			path TEXT NOT NULL
		);`,
		`CREATE TABLE IF NOT EXISTS films (
			film_id SERIAL PRIMARY KEY,
			title TEXT NOT NULL,
			is_serial BOOLEAN NOT NULL DEFAULT false,
			description TEXT,
			trailer_id INT REFERENCES trailers(trailer_id)
		);`,
		`CREATE TABLE IF NOT EXISTS film_cards (
			film_card_id SERIAL PRIMARY KEY,
			film_id INT REFERENCES films(film_id),
			path TEXT NOT NULL,
			is_horizontal BOOLEAN NOT NULL DEFAULT true 
		);`,
		`CREATE TABLE IF NOT EXISTS logos_films (
			logo_film_id SERIAL PRIMARY KEY,
			logo_id INT REFERENCES logos(logo_id),
			film_id INT REFERENCES films(film_id) 
		);`,
	}
	for _, query := range queries {
		_, err := DB.Exec(query)
		if err != nil {
			log.Fatal("Ошибка миграции:", err)
		}
	}

	log.Println("Миграция выполена успешно")
}

// INSERT INTO films (title, is_serial, description) VALUES ('Атака титанов', true, 'Люди сражаются с титанами, которые мечтают их съесть. Финал самого эпичного аниме современности'),
// ('Дандадан', true, 'Внучка медиума и юный уфолог внезапно встречают призраков и пришельцев. Хитовое аниме — безумное и смешное'),
// ('Железный человек', false, 'Он не просто металл — он живой! Идём в кино на высокотехнологичную легенду, соединяющую поколения.');
//INSERT INTO film_cards(film_id,path,is_horizontal) VALUES(1,'//avatars.mds.yandex.net/get-ott/212840/2a00000186a8b7a1951185c6175ed0f07fd0/2016x1134',true),(2,'//avatars.mds.yandex.net/get-ott/1652588/2a000001981bf0c61fd778cd9bc97146001c/2016x1134',true);

// INSERT INTO films (title, is_serial, description) VALUES ('Железный человек', false, 'Он не просто металл — он живой! Идём в кино на высокотехнологичную легенду, соединяющую поколения.')
