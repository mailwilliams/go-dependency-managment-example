package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func BaseHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello world",
	})
}

func getNameHandler(c *gin.Context) {
	name := c.Params.ByName("name")
	c.JSON(http.StatusOK, gin.H{
		"message": "hello " + name,
	})
}

func main() {
	router := gin.Default()
	router.GET("/", BaseHandler)
	router.GET("/:name", getNameHandler)
	_ = http.ListenAndServe(":8080", router)
}
