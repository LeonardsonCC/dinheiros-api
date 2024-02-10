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

func ListCategoriesHandler(c *gin.Context) {
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

	repo := categories_repo.CategoryRepository{DB: db}

	cats, err := repo.ListByUser(userID)
	if err != nil {
		rest.Err(c, "failed to get categories", err)
		return
	}

	if len(cats) == 0 {
		cats = []categories.Category{}
	}

	c.JSON(http.StatusOK, cats)
}
