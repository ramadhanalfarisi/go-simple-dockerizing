package repository

import (
	"database/sql"
	"log"

	"github.com/ramadhanalfarisi/go-simple-dockerizing/model"
)

type MangaRepository struct {
	DB *sql.DB
}

func NewMangaRepository(db *sql.DB) MangaRepositoryInterface {
	return &MangaRepository{DB: db}
}

// InsertManga implements MangaRepositoryInterface
func (m *MangaRepository) InsertManga(post model.PostManga) bool {
	stmt, err := m.DB.Prepare("INSERT INTO manga (title, genre, volumes, chapters, author) VALUES ($1, $2, $3, $4, $5)")
	if err != nil {
		log.Println(err)
		return false
	}
	defer stmt.Close()
	_, err2 := stmt.Exec(post.Title, post.Genre, post.Volumes, post.Chapters, post.Author)
	if err2 != nil {
		log.Println(err2)
		return false
	}
	return true
}

// SelectManga implements MangaRepositoryInterface
func (m *MangaRepository) SelectManga() []model.Manga {
	var result []model.Manga
	rows, err := m.DB.Query("SELECT * FROM manga")
	if err != nil {
		log.Println(err)
		return nil
	}
	for rows.Next() {
		var (
			id       uint
			title    string
			genre    string
			volumes  uint8
			chapters uint16
			author   string
		)
		err := rows.Scan(&id, &title, &genre, &volumes, &chapters, &author)
		if err != nil {
			log.Println(err)
		} else {
			manga := model.Manga{Id: id, Title: title, Genre: genre, Volumes: volumes, Chapters: chapters, Author: author}
			result = append(result, manga)
		}
	}
	return result
}
