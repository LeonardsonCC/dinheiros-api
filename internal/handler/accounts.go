package handler

import (
	"net/http"
	"strconv"

	"github.com/LeonardsonCC/dinheiros/db"
	"github.com/LeonardsonCC/dinheiros/internal/domain"
	"github.com/LeonardsonCC/dinheiros/internal/repository"
	"github.com/LeonardsonCC/dinheiros/internal/telemetry"
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
	ctx, sp := telemetry.GetAppTracer().Start(c.Request.Context(), "handler accounts")
	defer sp.End()

	db, err := db.GetConnection(ctx)
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

	var a domain.Account
	a.UserID = userID

	if err := c.ShouldBindJSON(&a); err != nil {
		rest.Err(c, "account invalid", err)
		return
	}

	repo := repository.AccountRepository{DB: db}

	err = repo.Create(ctx, a)
	if err != nil {
		rest.Err(c, "failed to create account", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "account created",
	})
}

func GetAccountsHandler(c *gin.Context) {
	ctx, sp := telemetry.GetAppTracer().Start(c.Request.Context(), "handler accounts")
	defer sp.End()

	db, err := db.GetConnection(ctx)
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

	repo := repository.AccountRepository{DB: db}

	accs, err := repo.Get(ctx, userID)
	if err != nil {
		rest.Err(c, "failed to get accounts", err)
		return
	}

	if accs == nil {
		accs = []domain.Account{}
	}

	c.JSON(http.StatusOK, accs)
}

func DeleteAccountHandler(c *gin.Context) {
	ctx, sp := telemetry.GetAppTracer().Start(c.Request.Context(), "handler accounts")
	defer sp.End()

	db, err := db.GetConnection(ctx)
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

	accountIDStr := c.Params.ByName("id")
	accountID, err := strconv.Atoi(accountIDStr)
	if err != nil {
		rest.Err(c, "invalid account id id", err)
	}

	repo := repository.AccountRepository{DB: db}

	err = repo.Delete(ctx, userID, accountID)
	if err != nil {
		rest.Err(c, "failed to delete account", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "account deleted",
	})
}

func UpdateAccountHandler(c *gin.Context) {
	ctx, sp := telemetry.GetAppTracer().Start(c.Request.Context(), "handler accounts")
	defer sp.End()

	db, err := db.GetConnection(ctx)
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

	accountIDStr := c.Params.ByName("id")
	accountID, err := strconv.Atoi(accountIDStr)
	if err != nil {
		rest.Err(c, "invalid account id", err)
	}

	var a domain.Account
	a.ID = accountID
	a.UserID = userID

	if err := c.ShouldBindJSON(&a); err != nil {
		rest.Err(c, "account invalid", err)
		return
	}

	repo := repository.AccountRepository{DB: db}

	err = repo.Update(ctx, a)
	if err != nil {
		rest.Err(c, "failed to update account", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "account updated",
	})
}
