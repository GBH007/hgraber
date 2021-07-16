package db

import (
	"log"
	"time"
)

// type Title struct {
// 	ID               int
// 	Name             string
// 	URL              string
// 	PageCount        int
// 	CreationTime     time.Time
// 	Loaded           bool
// 	ParsedPages      bool
// 	ParsedTags       bool
// 	ParsedAuthors    bool
// 	ParsedCharacters bool
// }

// InsertTitle добавляет тайтл
func InsertTitle(name, URL string, loaded bool) (int, error) {
	result, err := _db.Exec(
		`INSERT INTO titles(name, url, creation_time, loaded) VALUES(?, ?, ?, ?)`,
		name, URL, time.Now(), loaded,
	)
	if err != nil {
		log.Println(err)
		return -1, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Println(err)
		return -1, err
	}
	return int(id), nil
}

// UpdateTitleParsedPage обновляет информацию об обработанных страницах в тайтле
func UpdateTitleParsedPage(id, count int, success bool) error {
	_, err := _db.Exec(`UPDATE titles SET parsed_pages = ?, page_count = ? WHERE id = ?`, success, count, id)
	if err != nil {
		log.Println(err)
	}
	return err
}

// InsertPage добавляет страницу тайтла
func InsertPage(id int, name, URL string, page_number int) error {
	_, err := _db.Exec(
		`INSERT INTO pages(title_id, name, url, page_number, success) VALUES(?, ?, ?, ?, ?)
		ON CONFLICT(title_id, page_number) DO UPDATE SET name = excluded.name, url = excluded.url, success = false`,
		id, name, URL, page_number, false,
	)
	if err != nil {
		log.Println(err)
	}
	return err
}
