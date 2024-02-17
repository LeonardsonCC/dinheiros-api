package categories_route

import (
	"net/http"
	"strconv"

	categories_repo "github.com/LeonardsonCC/dinheiros/categories/repo"
	"github.com/LeonardsonCC/dinheiros/db"
	"github.com/LeonardsonCC/dinheiros/internal/domain"
	"github.com/LeonardsonCC/dinheiros/rest"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func CreateCategoryHandler(c *gin.Context) {
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

	var cat domain.Category
	if err := c.ShouldBindJSON(&cat); err != nil {
		rest.Err(c, "category invalid", err)
		return
	}
	cat.UserID = userID

	repo := categories_repo.CategoryRepository{DB: db}

	err = repo.Create(cat)
	if err != nil {
		rest.Err(c, "failed to create category", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "category created",
	})
}
