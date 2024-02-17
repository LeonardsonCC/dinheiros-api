package main

import (
	"net/http"

	"github.com/LeonardsonCC/dinheiros/db"
	"github.com/LeonardsonCC/dinheiros/internal/handler"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	_ = godotenv.Load()

	// singleton so start here to be used for routes later
	_, err := db.GetConnection()
	if err != nil {
		panic(err)
	}

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	for _, route := range handler.Routes {
		route(r)
	}

	err = r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if err != nil {
		panic(err)
	}
}
