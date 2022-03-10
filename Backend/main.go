package main

import (
	"db"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"net/http"
	"todo"
)

func main() {
	sqlDb := db.NewSql()
	defer sqlDb.Close()

	r := setupRouter(sqlDb)
	_ = r.Run(":8080")
}

func setupRouter(sqlDb *sqlx.DB) *gin.Engine {

	r := gin.Default()

	todoRepo := todo.New(sqlDb)

	r.Use(gzip.Gzip(gzip.DefaultCompression))
	r.Use(cors.New(corsConfig()))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})

	r.POST("/todo", todoRepo.Add)
	r.PUT("/todo/:id", todoRepo.Update)
	r.DELETE("/todo/:id", todoRepo.Delete)
	r.GET("/todo", todoRepo.Get)
	r.GET("/todo/:id", todoRepo.GetTodoById)
	return r
}
func corsConfig() cors.Config {
	defaultCors := cors.DefaultConfig()
	defaultCors.AddExposeHeaders("access-token", "refresh-token")
	defaultCors.AllowAllOrigins = true
	defaultCors.AllowHeaders = []string{"*"}
	return defaultCors
}
