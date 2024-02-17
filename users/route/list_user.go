package users_route

import (
	"net/http"

	"github.com/LeonardsonCC/dinheiros/db"
	"github.com/LeonardsonCC/dinheiros/internal/repository"
	"github.com/LeonardsonCC/dinheiros/rest"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func ListUserHandler(c *gin.Context) {
	db, err := db.GetConnection()
	if err != nil {
		rest.Err(c, "failed to connect to database", err)
		return
	}

	repo := repository.UserRepository{DB: db}

	users, err := repo.List()
	if err != nil {
		rest.Err(c, "failed to get user", err)
		return
	}

	c.JSON(http.StatusOK, users)
}
