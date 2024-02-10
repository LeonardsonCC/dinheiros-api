package categories_route

import (
	"net/http"
	"strconv"

	"github.com/LeonardsonCC/dinheiros/categories"
	categories_repo "github.com/LeonardsonCC/dinheiros/categories/repo"
	"github.com/LeonardsonCC/dinheiros/db"
	"github.com/LeonardsonCC/dinheiros/rest"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func UpdateCategoryHandler(c *gin.Context) {
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

	categoryIDStr := c.Params.ByName("id")
	categoryID, err := strconv.Atoi(categoryIDStr)
	if err != nil {
		rest.Err(c, "invalid account id id", err)
	}

	var cat categories.Category
	if err := c.ShouldBindJSON(&cat); err != nil {
		rest.Err(c, "category invalid", err)
		return
	}
	cat.UserID = userID
	cat.ID = categoryID

	repo := categories_repo.CategoryRepository{DB: db}

	err = repo.Update(cat)
	if err != nil {
		rest.Err(c, "failed to update category", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "category updated",
	})
}
