package categories_route

import "github.com/gin-gonic/gin"

func SetupRoutes(r *gin.Engine) {
	g := r.Group("/category")
	g.POST("/", CreateCategoryHandler)
	g.PUT("/:id", UpdateCategoryHandler)
	g.GET("/", ListCategoriesHandler)
	g.GET("/:id", GetCategoryHandler)
	g.DELETE("/:id", DeleteCategoryHandler)
}
