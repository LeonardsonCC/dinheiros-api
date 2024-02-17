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

func CategoriesRoutes(r *gin.Engine) {
	g := r.Group("/category")
	g.POST("/", CreateCategoryHandler)
	g.PUT("/:id", UpdateCategoryHandler)
	g.GET("/", ListCategoriesHandler)
	g.GET("/:id", GetCategoryHandler)
	g.DELETE("/:id", DeleteCategoryHandler)
}

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

	repo := repository.CategoryRepository{DB: db}

	err = repo.Create(cat)
	if err != nil {
		rest.Err(c, "failed to create category", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "category created",
	})
}

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

	repo := repository.CategoryRepository{DB: db}

	err = repo.Delete(categoryID)
	if err != nil {
		rest.Err(c, "failed to delte category", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "category deleted",
	})
}

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

	repo := repository.CategoryRepository{DB: db}

	cats, err := repo.ListByUser(userID)
	if err != nil {
		rest.Err(c, "failed to get categories", err)
		return
	}

	if len(cats) == 0 {
		cats = []domain.Category{}
	}

	c.JSON(http.StatusOK, cats)
}

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

	var cat domain.Category
	if err := c.ShouldBindJSON(&cat); err != nil {
		rest.Err(c, "category invalid", err)
		return
	}
	cat.UserID = userID
	cat.ID = categoryID

	repo := repository.CategoryRepository{DB: db}

	err = repo.Update(cat)
	if err != nil {
		rest.Err(c, "failed to update category", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "category updated",
	})
}
