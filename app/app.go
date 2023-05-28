package app

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ramadhanalfarisi/go-simple-dockerizing/controller"
)

type App struct {
	DB     *sql.DB
	Routes *gin.Engine
}

func (a *App) CreateConnection(){
	connStr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", UNAMEDB, PASSDB, HOSTDB, DBNAME)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	a.DB = db
}

func (a *App) CreateRoutes() {
	routes := gin.Default()
	controller := controller.NewMangaController(a.DB)
	routes.GET("/manga", controller.GetManga)
	routes.POST("/manga", controller.InsertManga)
	a.Routes = routes
}

func (a *App) Run(){
	a.Routes.Run(":8080")
}
