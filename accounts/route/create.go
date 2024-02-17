package accounts_route

import (
	"net/http"
	"strconv"

	accounts_repo "github.com/LeonardsonCC/dinheiros/accounts/repo"
	"github.com/LeonardsonCC/dinheiros/db"
	"github.com/LeonardsonCC/dinheiros/internal/domain"
	"github.com/LeonardsonCC/dinheiros/rest"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func CreateAccountHandler(c *gin.Context) {
	db, err := db.GetConnection()
	if err != nil {
		rest.Err(c, "failed to connect to database", err)
		return
	}

	userIDStr := c.GetHeader("user")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		rest.Err(c, "invalid user id", err)
	}

	var a domain.Account
	a.UserID = userID

	if err := c.ShouldBindJSON(&a); err != nil {
		rest.Err(c, "account invalid", err)
		return
	}

	repo := accounts_repo.AccountRepository{DB: db}

	err = repo.Create(a)
	if err != nil {
		rest.Err(c, "failed to create account", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "account created",
	})
}
