package accounts_route

import "github.com/gin-gonic/gin"

func SetupRoutes(r *gin.Engine) {
	g := r.Group("/account")

	g.POST("/", CreateAccountHandler)
	g.GET("/", GetAccountsHandler)
	g.DELETE("/:id", DeleteAccountHandler)
	g.PUT("/:id", UpdateAccountHandler)
}
