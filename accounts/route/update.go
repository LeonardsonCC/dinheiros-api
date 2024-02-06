package accounts_route

import (
	"net/http"
	"strconv"

	"github.com/LeonardsonCC/dinheiros/accounts"
	accounts_repo "github.com/LeonardsonCC/dinheiros/accounts/repo"
	"github.com/LeonardsonCC/dinheiros/db"
	"github.com/LeonardsonCC/dinheiros/rest"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func UpdateAccountHandler(c *gin.Context) {
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

	var a accounts.Account
	a.ID = accountID
	a.UserID = userID

	if err := c.ShouldBindJSON(&a); err != nil {
		rest.Err(c, "account invalid", err)
		return
	}

	repo := accounts_repo.AccountRepository{DB: db}

	err = repo.Update(a)
	if err != nil {
		rest.Err(c, "failed to update address", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "address updated",
	})
}
