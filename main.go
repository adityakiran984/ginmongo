package main

import (
	// "fmt"
	// "example/learnginmongo/configs"
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Starboy", Artist: "The Weeknd", Price: 17.99},
	{ID: "3", Title: "Currents", Artist: "Tame Impala", Price: 20.99},
}

func main() {
	router := gin.Default()

	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbum)
	router.PUT("/albums/:id", updateAlbum)
	router.DELETE("/albums/:id", deleteAlbum)

	router.Run("localhost:8080")
}

// Get all albums
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// Get album by ID
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// Add new album
func postAlbum(c *gin.Context) {
	var newAlbum album
	if err := c.BindJSON(&newAlbum); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"Error message": err.Error() + " couldn't add the album"})
		return
	}
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusOK, gin.H{"message": "album added successfully"})
	// albums = append(albums, newAlbum)
}

// Update an existing album
func updateAlbum(c *gin.Context) {
	id := c.Param("id")
	var updated album

	if err := c.BindJSON(&updated); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, a := range albums {
		if a.ID == id {
			// Updating fields
			albums[i].Title = updated.Title
			albums[i].Artist = updated.Artist
			albums[i].Price = updated.Price

			c.IndentedJSON(http.StatusOK, albums[i])
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// Delete an album
func deleteAlbum(c *gin.Context) {
	id := c.Param("id")

	for i, a := range albums {
		if a.ID == id {
			albums = append(albums[:i], albums[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "album deleted successfully"})
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not present"})
}
