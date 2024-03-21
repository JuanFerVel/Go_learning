package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Tittle string  `json:"tittle"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Tittle: "Nueva era", Artist: "Juan", Price: 45.22},
	{"2", "Nueva era 2", "Juan", 46.23},
	{"3", "Nueva era 3", "Juan", 47.24},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbum(c *gin.Context) {
	var newAlbum album
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	albums = append(albums, newAlbum)

	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumById(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album no encontrado"})
}

func updateAlbum(c *gin.Context) {
	id := c.Param("id")
	i, _ := strconv.Atoi(id)

	var updateAlbum album

	if err := c.BindJSON(&updateAlbum); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album no encontrado"})
		return
	}

	albums[i-1] = updateAlbum

	c.IndentedJSON(http.StatusCreated, updateAlbum)
}

func deleteAlbum(c *gin.Context) {
	id := c.Param("id")
	i, _ := strconv.Atoi(id)

	for _, a := range albums {
		if a.ID == id {
			albumDelete := albums[i-1]
			albums = append(albums[:i-1], albums[i:]...)
			c.IndentedJSON(http.StatusOK, albumDelete)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album no encontrado"})
}

func main() {
	router := gin.Default()
	// router.GET("/", func(ctx *gin.Context) {
	// 	ctx.JSON(200, gin.H{
	// 		"message": "Hola Mundo",
	// 	})
	// })

	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbum)
	router.GET("/albums/:id", getAlbumById)
	router.PUT("/albums/:id", updateAlbum)
	router.DELETE("/albums/:id", deleteAlbum)

	router.Run(":8080")
	// router.Run("localhost:8080")
}
