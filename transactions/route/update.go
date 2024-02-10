package transactions_route

import (
	"fmt"
	"net/http"
	"strconv"

	categories_repo "github.com/LeonardsonCC/dinheiros/categories/repo"
	"github.com/LeonardsonCC/dinheiros/db"
	"github.com/LeonardsonCC/dinheiros/rest"
	"github.com/LeonardsonCC/dinheiros/transactions"
	transactions_repo "github.com/LeonardsonCC/dinheiros/transactions/repo"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func UpdateTransactionHandler(c *gin.Context) {
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

	transactionIDStr := c.Params.ByName("id")
	transactionID, err := strconv.Atoi(transactionIDStr)
	if err != nil {
		rest.Err(c, "invalid transaction id id", err)
	}

	var t transactions.TransactionJson
	t.ID = transactionID

	if err := c.ShouldBindJSON(&t); err != nil {
		rest.Err(c, "account invalid", err)
		return
	}

	tx, err := transactions.MapJsonToDomain(t)
	if err != nil {
		rest.Err(c, "transaction invalid 2", err)
		return
	}

	if tx.AccountID == 0 {
		rest.Err(c, "transaction without account id", fmt.Errorf("failed to get account id"))
		return
	}

	repo := transactions_repo.TransactionsRepository{DB: db}
	categoriesRepo := categories_repo.CategoriesRepository{DB: db}

	err = repo.Update(tx)
	if err != nil {
		rest.Err(c, "failed to update transaction", err)
		return
	}

	err = categoriesRepo.Save(userID, transactionID, t.Categories)
	if err != nil {
		rest.Err(c, "failed to update categories", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "transaction updated",
	})
}
