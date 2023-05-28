package controller

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/ramadhanalfarisi/go-simple-dockerizing/model"
	"github.com/ramadhanalfarisi/go-simple-dockerizing/repository"
)

type MangaController struct {
	DB *sql.DB
}

func NewMangaController(db *sql.DB) MangaControllerInterface {
	return &MangaController{DB: db}
}

// GetManga implements MangaControllerInterface
func (m *MangaController) GetManga(g *gin.Context) {
	db := m.DB
	repo_manga := repository.NewMangaRepository(db)
	get_manga := repo_manga.SelectManga()
	if get_manga != nil {
		g.JSON(200, gin.H{"status": "success", "data": get_manga, "msg": "get manga successfully"})
	} else {
		g.JSON(200, gin.H{"status": "success", "data": nil, "msg": "get manga successfully"})
	}
}

// InsertManga implements MangaControllerInterface
func (m *MangaController) InsertManga(g *gin.Context) {
	db := m.DB
	var post model.PostManga
	if err := g.ShouldBindJSON(&post); err == nil {
		repo_manga := repository.NewMangaRepository(db)
		insert := repo_manga.InsertManga(post)
		if insert {
			g.JSON(200, gin.H{"status": "success", "msg": "insert manga successfully"})
		} else {
			g.JSON(500, gin.H{"status": "failed", "msg": "insert manga failed"})
		}
	} else {
		g.JSON(400, gin.H{"status": "success", "msg": err})
	}
}
