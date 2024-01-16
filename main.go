package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var book = []string{
	"jaocb1",
	"jaocb2",
	"jaocb3",
}

func main() {
	app := gin.Default()
	app.GET("/ping", checkPing)
	app.GET("/books", getBooks)

	http.ListenAndServe(":6000", app)
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, book)
}

func checkPing(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "working",
	})

}
