package transactions_route

import (
	"net/http"
	"strconv"

	categories_repo "github.com/LeonardsonCC/dinheiros/categories/repo"
	"github.com/LeonardsonCC/dinheiros/db"
	"github.com/LeonardsonCC/dinheiros/internal/domain"
	"github.com/LeonardsonCC/dinheiros/rest"
	transactions_repo "github.com/LeonardsonCC/dinheiros/transactions/repo"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func GetSingleTransactionHandler(c *gin.Context) {
	db, err := db.GetConnection()
	if err != nil {
		rest.Err(c, "failed to connect to database", err)
		return
	}

	// TODO: VALIDATE USER IS THE OWNER OF ACCOUNT
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

	txs, err := repo.Get(transactionID)
	if err != nil {
		rest.Err(c, "failed to get transactions", err)
		return
	}

	cats, err := catRepo.GetCategoriesFromTransaction(transactionID)
	if err != nil {
		rest.Err(c, "failed to get transactions", err)
		return
	}

	ts := make([]domain.TransactionJson, len(txs))
	for i, tx := range txs {
		ts[i] = domain.MapDomainToJson(tx, cats)
	}

	c.JSON(http.StatusOK, ts)
}
