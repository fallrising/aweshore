package main

import (
	"aweshore/internal/app/handler"
	"aweshore/pkg/db"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Initialize the database
	db.Init()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders: []string{"Origin", "Content-Type"},
	}))

	// Routes
	r.POST("/notes", handler.CreateNote)
	r.GET("/notes/:id", handler.GetNote)
	r.GET("/notes", handler.GetAllNotes)
	r.PUT("/notes/:id", handler.UpdateNote)
	r.DELETE("/notes/:id", handler.DeleteNote)

	r.Run() // listen and serve on 0.0.0.0:8080
}
