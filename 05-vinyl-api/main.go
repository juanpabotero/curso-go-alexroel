package main

import (
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
	{ID: "1", Title: "The Dark Side of the Moon", Artist: "Pink Floyd", Price: 19.99},
	{ID: "2", Title: "Back in Black", Artist: "AC/DC", Price: 17.99},
	{ID: "3", Title: "Led Zeppelin IV", Artist: "Led Zeppelin", Price: 16.99},
}

// HANDLERS

// gin.Context es el contexto de la peticion
// c.IndentedJSON(200, gin.H{...}) es el metodo para responder con un JSON indentado
// http.StatusOK es una cosntante para el codigo de estado 200
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbum(c *gin.Context) {
	var newAlbum album
	// c.BindJSON(&newAlbum) es el metodo para mapear el JSON a la estructura
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumByID(c *gin.Context) {
	// c.Param("id") es el metodo para obtener el parametro de la URL
	id := c.Param("id")
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	// gin.H es un alias para map[string]interface{}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func updateAlbum(c *gin.Context) {
	id := c.Param("id")
	var updatedAlbum album
	if err := c.BindJSON(&updatedAlbum); err != nil {
		return
	}
	for i, a := range albums {
		if a.ID == id {
			albums[i] = updatedAlbum
			c.IndentedJSON(http.StatusOK, updatedAlbum)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func deleteAlbum(c *gin.Context) {
	id := c.Param("id")
	for i, a := range albums {
		if a.ID == id {
			albums = append(albums[:i], albums[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "album deleted"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func main() {
	// configuracion por defecto de gin
	router := gin.Default()

	// RUTAS

	// c.JSON(200, gin.H{...}) es el metodo para responder con un JSON
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Vinyl API",
		})
	})
	router.GET("/album", getAlbums)
	router.GET("/album/:id", getAlbumByID)
	router.POST("/album", postAlbum)
	router.PUT("/album/:id", updateAlbum)
	router.DELETE("/album/:id", deleteAlbum)

	// correr el servidor
	router.Run(":8080")
}
