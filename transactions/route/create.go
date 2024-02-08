package transactions_route

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/LeonardsonCC/dinheiros/db"
	"github.com/LeonardsonCC/dinheiros/rest"
	"github.com/LeonardsonCC/dinheiros/transactions"
	transactions_repo "github.com/LeonardsonCC/dinheiros/transactions/repo"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func CreateTransactionHandler(c *gin.Context) {
	db, err := db.GetConnection()
	if err != nil {
		rest.Err(c, "failed to connect to database", err)
		return
	}

	userIDStr := c.GetHeader("user")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		rest.Err(c, "invalid user id", err)
		return
	}

	var accID int
	accountIDStr := c.Param("account_id")
	if accountIDStr != "" {
		accountID, err := strconv.Atoi(accountIDStr)
		if err != nil {
			rest.Err(c, "invalid account id", err)
			return
		}
		accID = accountID
	}

	var t transactions.TransactionJson
	if err := c.ShouldBindJSON(&t); err != nil {
		rest.Err(c, "transaction invalid", err)
		return
	}

	tx, err := transactions.MapJsonToDomain(t)
	if err != nil {
		rest.Err(c, "transaction invalid 2", err)
		return
	}

	tx.UserID = userID
	if accID != 0 {
		tx.AccountID = accID
	}

	if tx.AccountID == 0 {
		rest.Err(c, "transaction without account id", fmt.Errorf("failed to get account id"))
		return
	}

	repo := transactions_repo.TransactionsRepository{DB: db}

	err = repo.Create(tx)
	if err != nil {
		rest.Err(c, "failed to create transaction", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "transaction created",
	})
}
