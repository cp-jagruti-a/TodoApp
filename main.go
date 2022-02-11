package main

import (
	"db"
	"net/http"
    "todo"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
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

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})

	r.POST("/api/v1/todo", todoRepo.Add)
	r.PUT("/api/v1/todo/:id", todoRepo.Update)
	r.DELETE("/api/v1/todo/:id", todoRepo.Delete)
	r.GET("/api/v1/todo/all", todoRepo.Get)
	r.GET("/api/v1/todo/:id", todoRepo.GetTodoById)
	return r
}

