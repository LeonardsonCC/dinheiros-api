package transactions_route

import (
	"net/http"
	"strconv"

	"github.com/LeonardsonCC/dinheiros/db"
	"github.com/LeonardsonCC/dinheiros/internal/domain"
	"github.com/LeonardsonCC/dinheiros/internal/repository"
	"github.com/LeonardsonCC/dinheiros/rest"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func GetTransactionsHandler(c *gin.Context) {
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

	accountIDStr := c.Param("account_id")
	accountID, err := strconv.Atoi(accountIDStr)
	if err != nil {
		accountID = 0
	}

	repo := repository.TransactionsRepository{DB: db}
	repoCats := repository.CategoryRepository{DB: db}

	txs, err := repo.List(userID, accountID)
	if err != nil {
		rest.Err(c, "failed to get addresses", err)
		return
	}

	cats, err := repoCats.GetCategoriesFromAccount(userID, accountID)
	if err != nil {
		rest.Err(c, "failed to get addresses", err)
		return
	}

	ts := make([]domain.TransactionJson, len(txs))
	for i, tx := range txs {
		ts[i] = domain.MapDomainToJson(tx, cats[tx.ID])
	}

	c.JSON(http.StatusOK, ts)
}
