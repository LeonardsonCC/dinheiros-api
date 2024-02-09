package transactions_route

import "github.com/gin-gonic/gin"

func SetupRoutes(r *gin.Engine) {
	txGroup := r.Group("/account/:account_id/transactions")
	txGroup.POST("/", CreateTransactionHandler)
	txGroup.GET("/", GetTransactionsHandler)

	t := r.Group("/transaction")
	t.POST("/", CreateTransactionHandler)
	t.GET("/:id", GetSingleTransactionHandler)
	t.GET("/", GetTransactionsHandler)
	t.DELETE("/:id", DeleteTransactionHandler)
	t.PUT("/:id", UpdateTransactionHandler)
}
