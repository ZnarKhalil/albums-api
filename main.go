package main

import (
	"log"

	"github.com/ZnarKhalil/albums-api/db"
	"github.com/ZnarKhalil/albums-api/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the database
	db.Init()
	defer db.DB.Close()

	// Initialize the Gin router
	router := gin.Default()

	router.GET("/albums", handlers.GetAlbums)
	router.GET("/albums/:id", handlers.GetAlbumById)
	router.POST("/albums", handlers.AddAlbum)

	// Start the server on port 8080
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
