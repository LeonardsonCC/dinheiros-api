package transactions_route

import "github.com/gin-gonic/gin"

func SetupRoutes(r *gin.Engine) {
	g := r.Group("/account/:account_id/transaction")

	g.POST("/", CreateTransactionHandler)
	g.GET("/", GetTransactionsHandler)
	g.GET("/:id", GetSingleTransactionHandler)
}
