package accounts_route

import (
	"net/http"
	"strconv"

	"github.com/LeonardsonCC/dinheiros/db"
	"github.com/LeonardsonCC/dinheiros/internal/repository"
	"github.com/LeonardsonCC/dinheiros/rest"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func DeleteAccountHandler(c *gin.Context) {
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

	accountIDStr := c.Params.ByName("id")
	accountID, err := strconv.Atoi(accountIDStr)
	if err != nil {
		rest.Err(c, "invalid account id id", err)
	}

	repo := repository.AccountRepository{DB: db}

	err = repo.Delete(userID, accountID)
	if err != nil {
		rest.Err(c, "failed to delete address", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "address deleted",
	})
}
