package transactions_route

import (
	"net/http"
	"strconv"

	"github.com/LeonardsonCC/dinheiros/db"
	"github.com/LeonardsonCC/dinheiros/rest"
	"github.com/LeonardsonCC/dinheiros/transactions"
	transactions_repo "github.com/LeonardsonCC/dinheiros/transactions/repo"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func GetTransactionsHandler(c *gin.Context) {
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

	accountIDStr := c.Param("account_id")
	accountID, err := strconv.Atoi(accountIDStr)
	if err != nil {
		rest.Err(c, "invalid account id", err)
		return
	}

	repo := transactions_repo.TransactionsRepository{DB: db}

	txs, err := repo.List(accountID)
	if err != nil {
		rest.Err(c, "failed to get addresses", err)
		return
	}

	ts := make([]transactions.TransactionJson, len(txs))
	for i, tx := range txs {
		ts[i] = transactions.MapDomainToJson(tx)
	}

	c.JSON(http.StatusOK, ts)
}
