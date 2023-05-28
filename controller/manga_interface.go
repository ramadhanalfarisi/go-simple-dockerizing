package controller

import "github.com/gin-gonic/gin"

type MangaControllerInterface interface {
	InsertManga(g *gin.Context)
	GetManga(g *gin.Context)
}