package users_route

import (
	"net/http"

	"github.com/LeonardsonCC/dinheiros/db"
	"github.com/LeonardsonCC/dinheiros/internal/domain"
	"github.com/LeonardsonCC/dinheiros/internal/repository"
	"github.com/LeonardsonCC/dinheiros/rest"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func CreateUserHandler(c *gin.Context) {
	db, err := db.GetConnection()
	if err != nil {
		rest.Err(c, "failed to connect to database", err)
		return
	}

	var u domain.User

	if err := c.ShouldBindJSON(&u); err != nil {
		rest.Err(c, "user invalid", err)
		return
	}

	repo := repository.UserRepository{DB: db}

	err = repo.Create(u)
	if err != nil {
		rest.Err(c, "failed to create user", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user created",
		"email":   u.Email,
	})
}
