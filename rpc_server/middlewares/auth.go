package middlewares

import (
	"aastar_dashboard_back/config"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"time"
)

var jwtMiddleware *jwt.GinJWTMiddleware

func GinJwtMiddleware() *jwt.GinJWTMiddleware {
	return jwtMiddleware
}

func AuthHandler() gin.HandlerFunc {
	m, _ := jwt.New(&jwt.GinJWTMiddleware{
		Realm:      config.GetSystemConfigByKey(config.KeyJwtRealm),
		Key:        []byte(config.GetSystemConfigByKey(config.KeyJwtSecret)),
		Timeout:    time.Hour * 48,
		MaxRefresh: time.Hour / 2,
		IdentityHandler: func(c *gin.Context) interface{} {
			payload := jwt.ExtractClaims(c)
			return payload["user_id"]
		},
		IdentityKey: "user_id",
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(string); ok {
				return jwt.MapClaims{
					"user_id": v,
				}
			}
			return jwt.MapClaims{}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			userId := c.GetString("user_id")
			if userId == "" {
				return nil, jwt.ErrMissingLoginValues
			}
			return userId, nil
		},

		Authorizator: func(data interface{}, c *gin.Context) bool {
			// always return true unless the permission feature started
			return true
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		TokenLookup:   "header: Authorization",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
		HTTPStatusMessageFunc: func(e error, c *gin.Context) string {
			return "401 Unauthorized"
		},
	})

	jwtMiddleware = m

	return m.MiddlewareFunc()
}
