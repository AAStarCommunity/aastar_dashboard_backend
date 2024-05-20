package oauth

import (
	"aastar_dashboard_back/rpc_server/middlewares"
	"github.com/gin-gonic/gin"
)

func Logout(ctx *gin.Context) {
	middlewares.GinJwtMiddleware().LogoutHandler(ctx)
}
