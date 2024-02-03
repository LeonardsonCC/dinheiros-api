package main

import (
	"net/http"

	accounts_route "github.com/LeonardsonCC/dinheiros/accounts/route"
	"github.com/LeonardsonCC/dinheiros/db"
	users_route "github.com/LeonardsonCC/dinheiros/users/route"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("error loading .env file")
	}

	// singleton so start here to be used for routes later
	db.GetConnection()

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	users_route.SetupRoutes(r)
	accounts_route.SetupRoutes(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
