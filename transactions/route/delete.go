package transactions_route

import (
	"net/http"
	"strconv"

	"github.com/LeonardsonCC/dinheiros/db"
	"github.com/LeonardsonCC/dinheiros/internal/repository"
	"github.com/LeonardsonCC/dinheiros/rest"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func DeleteTransactionHandler(c *gin.Context) {
	db, err := db.GetConnection()
	if err != nil {
		rest.Err(c, "failed to connect to database", err)
		return
	}

	transactionIDStr := c.Param("id")
	transactionID, err := strconv.Atoi(transactionIDStr)
	if err != nil {
		rest.Err(c, "invalid transaction id", err)
		return
	}

	repo := repository.TransactionsRepository{DB: db}
	catRepo := repository.CategoryRepository{DB: db}

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
