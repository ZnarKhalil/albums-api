package handlers

import (
	"net/http"
	"strconv"

	"github.com/ZnarKhalil/albums-api/db"
	"github.com/ZnarKhalil/albums-api/models"
	"github.com/gin-gonic/gin"
)

// GetAlbums handles GET requests to /albums
func GetAlbums(c *gin.Context) {
	artist := c.Query("artist") // Retrieve the artist query parameter, if present

	var albums []models.Album
	var err error

	if artist != "" {
		albums, err = models.AlbumsByArtist(db.DB, artist)
	} else {
		albums, err = models.GetAlbums(db.DB)
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, albums)
}

// GetAlbumsByArtist handles GET requests to /albums/:artist
func GetAlbumById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid album ID"})
		return
	}

	albums, err := models.AlbumByID(db.DB, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, albums)
}

// AddAlbum handles POST requests to add a new album
func AddAlbum(c *gin.Context) {
	var newAlbum models.Album

	if err := c.ShouldBindJSON(&newAlbum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	albumID, err := models.AddAlbum(db.DB, newAlbum)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"album_id": albumID})
}
