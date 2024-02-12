package transactions_route

import (
	"net/http"
	"strconv"

	categories_repo "github.com/LeonardsonCC/dinheiros/categories/repo"
	"github.com/LeonardsonCC/dinheiros/db"
	"github.com/LeonardsonCC/dinheiros/rest"
	transactions_repo "github.com/LeonardsonCC/dinheiros/transactions/repo"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func DeleteTransactionHandler(c *gin.Context) {
	db, err := db.GetConnection()
	if err != nil {
		rest.Err(c, "failed to connect to database", err)
		return
	}

	// TODO: validate user id
	// userIDStr := c.GetHeader("user")
	// userID, err := strconv.Atoi(userIDStr)
	// if err != nil {
	// 	rest.Err(c, "invalid user id", err)
	// }

	transactionIDStr := c.Param("id")
	transactionID, err := strconv.Atoi(transactionIDStr)
	if err != nil {
		rest.Err(c, "invalid transaction id", err)
		return
	}

	repo := transactions_repo.TransactionsRepository{DB: db}
	catRepo := categories_repo.CategoryRepository{DB: db}

	err = catRepo.DeleteByTransaction(transactionID)
	if err != nil {
		rest.Err(c, "failed to delete transaction categories", err)
		return
	}

	err = repo.Delete(transactionID)
	if err != nil {
		rest.Err(c, "failed to delete transaction", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "transaction deleted",
	})
}
