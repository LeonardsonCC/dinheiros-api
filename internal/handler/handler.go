package handler

import "github.com/gin-gonic/gin"

type RouteSetup func(r *gin.Engine)

var Routes = []RouteSetup{
	UsersRoutes,
	AccountsRoutes,
	TransactionsRoutes,
	CategoriesRoutes,
}
