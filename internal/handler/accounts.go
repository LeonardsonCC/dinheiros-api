package handler

import (
	"net/http"
	"strconv"

	"github.com/LeonardsonCC/dinheiros/db"
	"github.com/LeonardsonCC/dinheiros/internal/domain"
	"github.com/LeonardsonCC/dinheiros/internal/repository"
	"github.com/LeonardsonCC/dinheiros/rest"
	"github.com/gin-gonic/gin"
)

func AccountsRoutes(r *gin.Engine) {
	g := r.Group("/account")

	g.POST("/", CreateAccountHandler)
	g.GET("/", GetAccountsHandler)
	g.DELETE("/:id", DeleteAccountHandler)
	g.PUT("/:id", UpdateAccountHandler)
}

func CreateAccountHandler(c *gin.Context) {
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

	var a domain.Account
	a.UserID = userID

	if err := c.ShouldBindJSON(&a); err != nil {
		rest.Err(c, "account invalid", err)
		return
	}

	repo := repository.AccountRepository{DB: db}

	err = repo.Create(a)
	if err != nil {
		rest.Err(c, "failed to create account", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "account created",
	})
}

func GetAccountsHandler(c *gin.Context) {
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

	repo := repository.AccountRepository{DB: db}

	accs, err := repo.Get(userID)
	if err != nil {
		rest.Err(c, "failed to get addresses", err)
		return
	}

	if accs == nil {
		accs = []domain.Account{}
	}

	c.JSON(http.StatusOK, accs)
}

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

	var a domain.Account
	a.ID = accountID
	a.UserID = userID

	if err := c.ShouldBindJSON(&a); err != nil {
		rest.Err(c, "account invalid", err)
		return
	}

	repo := repository.AccountRepository{DB: db}

	err = repo.Update(a)
	if err != nil {
		rest.Err(c, "failed to update address", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "address updated",
	})
}
