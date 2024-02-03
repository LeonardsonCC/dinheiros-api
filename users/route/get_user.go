package users_route

import (
	"net/http"
	"strconv"

	"github.com/LeonardsonCC/dinheiros/db"
	"github.com/LeonardsonCC/dinheiros/rest"
	users_repo "github.com/LeonardsonCC/dinheiros/users/repo"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func GetUserHandler(c *gin.Context) {
	db, err := db.GetConnection()
	if err != nil {
		rest.Err(c, "failed to connect to database", err)
		return
	}

	userIDStr := c.Params.ByName("id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		rest.Err(c, "invalid user id", err)
		return
	}

	repo := users_repo.UserRepository{DB: db}

	u, err := repo.Get(userID)
	if err != nil {
		rest.Err(c, "failed to get user", err)
		return
	}

	c.JSON(http.StatusOK, u)
}
