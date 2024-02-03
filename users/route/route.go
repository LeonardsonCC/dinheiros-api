package users_route

import "github.com/gin-gonic/gin"

func SetupRoutes(r *gin.Engine) {
	g := r.Group("/user")

	g.POST("/", CreateUserHandler)
	g.GET("/:email", GetUserHandler)
	g.GET("/", ListUserHandler)
}
