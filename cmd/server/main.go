package main

import (
	"aweshore/internal/app/handler"
	"aweshore/pkg/db"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Initialize the database
	db.Init()

	// Routes
	r.POST("/notes", handler.CreateNote)
	r.GET("/notes/:id", handler.GetNote)
	r.GET("/notes", handler.GetAllNotes)
	r.PUT("/notes/:id", handler.UpdateNote)
	r.DELETE("/notes/:id", handler.DeleteNote)

	r.Run() // listen and serve on 0.0.0.0:8080
}
