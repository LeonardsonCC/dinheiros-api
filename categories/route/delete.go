package categories_route

import (
	"net/http"
	"strconv"

	categories_repo "github.com/LeonardsonCC/dinheiros/categories/repo"
	"github.com/LeonardsonCC/dinheiros/db"
	"github.com/LeonardsonCC/dinheiros/rest"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func DeleteCategoryHandler(c *gin.Context) {
	db, err := db.GetConnection()
	if err != nil {
		rest.Err(c, "failed to connect to database", err)
		return
	}

	categoryIDStr := c.Param("id")
	categoryID, err := strconv.Atoi(categoryIDStr)
	if err != nil {
		rest.Err(c, "invalid category id", err)
		return
	}

	repo := categories_repo.CategoryRepository{DB: db}

	err = repo.Delete(categoryID)
	if err != nil {
		rest.Err(c, "failed to delte category", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "category deleted",
	})
}
