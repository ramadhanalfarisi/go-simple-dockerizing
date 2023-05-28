package repository

import "github.com/ramadhanalfarisi/go-simple-dockerizing/model"

type MangaRepositoryInterface interface {
	SelectManga() []model.Manga
	InsertManga(post model.PostManga) bool
}