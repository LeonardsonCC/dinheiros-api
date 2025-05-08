package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/LeonardsonCC/dinheiros/db"
	"github.com/LeonardsonCC/dinheiros/internal/domain"
	"github.com/LeonardsonCC/dinheiros/internal/repository"
	"github.com/LeonardsonCC/dinheiros/internal/telemetry"
	"github.com/LeonardsonCC/dinheiros/rest"
)

func UsersRoutes(r *gin.Engine) {
	g := r.Group("/user")

	g.POST("/", createUserHandler)
	g.GET("/:email", getUserHandler)
}

func createUserHandler(c *gin.Context) {
	ctx, sp := telemetry.GetAppTracer().Start(c.Request.Context(), "handler user")
	defer sp.End()

	db, err := db.GetConnection(ctx)
	if err != nil {
		rest.Err(c, "failed to connect to database", err)
		return
	}

	var u domain.User

	if err := c.ShouldBindJSON(&u); err != nil {
		rest.Err(c, "user invalid", err)
		return
	}

	repo := repository.UserRepository{DB: db}

	err = repo.Create(ctx, u)
	if err != nil {
		rest.Err(c, "failed to create user", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user created",
		"email":   u.Email,
	})
}

func getUserHandler(c *gin.Context) {
	ctx, sp := telemetry.GetAppTracer().Start(c.Request.Context(), "handler user")
	defer sp.End()

	db, err := db.GetConnection(ctx)
	if err != nil {
		rest.Err(c, "failed to connect to database", err)
		return
	}

	email := c.Params.ByName("email")

	repo := repository.UserRepository{DB: db}

	u, err := repo.Get(ctx, email)
	if err != nil {
		rest.Err(c, "failed to get user", err)
		return
	}

	c.JSON(http.StatusOK, u)
}
