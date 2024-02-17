package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/LeonardsonCC/dinheiros/db"
	"github.com/LeonardsonCC/dinheiros/internal/domain"
	"github.com/LeonardsonCC/dinheiros/internal/repository"
	"github.com/LeonardsonCC/dinheiros/rest"
	"github.com/gin-gonic/gin"
)

func TransactionsRoutes(r *gin.Engine) {
	txGroup := r.Group("/account/:account_id/transactions")
	txGroup.POST("/", CreateTransactionHandler)
	txGroup.GET("/", GetTransactionsHandler)

	t := r.Group("/transaction")
	t.POST("/", CreateTransactionHandler)
	t.GET("/:id", GetSingleTransactionHandler)
	t.GET("/", GetTransactionsHandler)
	t.DELETE("/:id", DeleteTransactionHandler)
	t.PUT("/:id", UpdateTransactionHandler)
}

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

	var t domain.TransactionJson
	if err := c.ShouldBindJSON(&t); err != nil {
		rest.Err(c, "transaction invalid", err)
		return
	}

	tx, cats, err := domain.MapJsonToDomain(t)
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

	repo := repository.TransactionsRepository{DB: db}
	catRepo := repository.CategoryRepository{DB: db}

	transactionID, err := repo.Create(tx)
	if err != nil {
		rest.Err(c, "failed to create transaction", err)
		return
	}

	err = catRepo.AddCategoryToTransaction(transactionID, cats)
	if err != nil {
		rest.Err(c, "failed to add categories", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("created transaction %d", transactionID),
	})
}

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

	repo := repository.TransactionsRepository{DB: db}
	catRepo := repository.CategoryRepository{DB: db}

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

func UpdateTransactionHandler(c *gin.Context) {
	db, err := db.GetConnection()
	if err != nil {
		rest.Err(c, "failed to connect to database", err)
		return
	}

	transactionIDStr := c.Params.ByName("id")
	transactionID, err := strconv.Atoi(transactionIDStr)
	if err != nil {
		rest.Err(c, "invalid transaction id id", err)
	}

	var t domain.TransactionJson
	t.ID = transactionID

	if err := c.ShouldBindJSON(&t); err != nil {
		rest.Err(c, "account invalid", err)
		return
	}

	tx, cats, err := domain.MapJsonToDomain(t)
	if err != nil {
		rest.Err(c, "transaction invalid 2", err)
		return
	}

	if tx.AccountID == 0 {
		rest.Err(c, "transaction without account id", fmt.Errorf("failed to get account id"))
		return
	}

	repo := repository.TransactionsRepository{DB: db}
	repoCats := repository.CategoryRepository{DB: db}

	err = repo.Update(tx)
	if err != nil {
		rest.Err(c, "failed to update transaction", err)
		return
	}

	err = repoCats.AddCategoryToTransaction(tx.ID, cats)
	if err != nil {
		rest.Err(c, "failed to update categories", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "transaction updated",
	})
}
