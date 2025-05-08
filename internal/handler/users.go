package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/LeonardsonCC/dinheiros/db"
	"github.com/LeonardsonCC/dinheiros/internal/domain"
	"github.com/LeonardsonCC/dinheiros/internal/logger"
	"github.com/LeonardsonCC/dinheiros/internal/repository"
	"github.com/LeonardsonCC/dinheiros/internal/telemetry"
	"github.com/LeonardsonCC/dinheiros/internal/telemetry/spans"
	"github.com/LeonardsonCC/dinheiros/rest"
)

func UsersRoutes(r *gin.Engine) {
	g := r.Group("/user")

	g.POST("/", createUserHandler)
	g.GET("/:email", getUserHandler)
}

func createUserHandler(c *gin.Context) {
	logger := logger.FromContext(c.Request.Context())

	ctx, sp := telemetry.GetAppTracer().Start(c.Request.Context(), spans.UserHandler)
	defer sp.End()

	db, err := db.GetConnection(ctx)
	if err != nil {
		rest.Err(c, "failed to connect to database", err)
		logger.Err(err).Msg("failed to connect to database")
		return
	}

	var u domain.User

	if err := c.ShouldBindJSON(&u); err != nil {
		rest.Err(c, "user invalid", err)
		logger.Err(err).Msg("user invalid")
		return
	}

	repo := repository.UserRepository{DB: db}

	err = repo.Create(ctx, u)
	if err != nil {
		rest.Err(c, "failed to create user", err)
		logger.Err(err).Msg("failed to create user")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user created",
		"email":   u.Email,
	})
}

func getUserHandler(c *gin.Context) {
	logger := logger.FromContext(c.Request.Context())

	ctx, sp := telemetry.GetAppTracer().Start(c.Request.Context(), spans.UserHandler)
	defer sp.End()

	db, err := db.GetConnection(ctx)
	if err != nil {
		rest.Err(c, "failed to connect to database", err)
		logger.Err(err).Msg("failed to connect to database")
		return
	}

	email := c.Params.ByName("email")

	repo := repository.UserRepository{DB: db}

	u, err := repo.Get(ctx, email)
	if err != nil {
		rest.Err(c, "failed to get user", err)
		logger.Err(err).Msg("failed to get user")
		return
	}

	c.JSON(http.StatusOK, u)
}
