package oauth

import (
	"aastar_dashboard_back/repository"
	"aastar_dashboard_back/rpc_server/middlewares"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PasswordRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// PasswordOauthLogin
// @Tags Oauth
// @Description PasswordOauthLogin
// @Accept json
// @Product json
// @Param passwordRequest body PasswordRequest true "PasswordRequest Model"
// @Router /oauth/password [post]
// @Success 200
func PasswordOauthLogin(ctx *gin.Context) {
	var req PasswordRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	user, err := repository.FindUserByEmail(req.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(400, gin.H{"error": "User not found"})
			return
		}
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if user.PassWord != req.Password {
		ctx.JSON(400, gin.H{"error": "Password not correct"})
		return
	}
	ctx.Set("user_id", "dylan")
	middlewares.GinJwtMiddleware().LoginHandler(ctx)
}

func PasswordSignUp() {

}
