package categories_route

import (
	"net/http"
	"strconv"

	"github.com/LeonardsonCC/dinheiros/db"
	"github.com/LeonardsonCC/dinheiros/internal/domain"
	"github.com/LeonardsonCC/dinheiros/internal/repository"
	"github.com/LeonardsonCC/dinheiros/rest"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func GetCategoryHandler(c *gin.Context) {
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

	repo := repository.CategoryRepository{DB: db}

	cats, err := repo.Get(categoryID)
	if err != nil {
		rest.Err(c, "failed to get category", err)
		return
	}

	if len(cats) == 0 {
		cats = []domain.Category{}
	}

	c.JSON(http.StatusOK, cats)
}
