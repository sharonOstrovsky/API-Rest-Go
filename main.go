package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type album struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Year   int    `json:"year"`
}

var albums = []album{
	{ID: "1", Title: "Familia", Artist: "Camila Cabello", Year: 2022},
	{ID: "2", Title: "21", Artist: "Adele", Year: 2011},
	{ID: "3", Title: "The Eminem Show", Artist: "Eminem", Year: 2002},
	{ID: "4", Title: "Meteora", Artist: "Linkin Park", Year: 2003},
	{ID: "5", Title: "25", Artist: "Adele", Year: 2015},
}

// handler
// devuelve al cliente toda la data un dato serializado de tipo json
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums) //devuelvo un estado
}

// c *gin.Context -> lo que nos envia el cliente
func postAlbums(c *gin.Context) {
	var newAlbum album

	//manejo de error
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)

	c.IndentedJSON(http.StatusCreated, albums)
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, album := range albums {
		if album.ID == id {
			c.IndentedJSON(http.StatusOK, album)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album no encontrado"})
}

func main() {
	fmt.Println("HOLA GO!")

	//inicializo las rutas
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.GET("/albums/:id", getAlbumByID)

	//para ejecutar en un servidor(localhost)
	router.Run("localhost:8080")

}
