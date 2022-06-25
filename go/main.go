package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HelloWorldHandler(c *gin.Context) {
	c.String(http.StatusOK, "Hello, World!");
}

func main() {
	router := gin.Default()

	router.GET("/", HelloWorldHandler)

	router.Run()
}