package main

import (
	"aastar_dashboard_back/config"
	"aastar_dashboard_back/data_view_repository"
	"aastar_dashboard_back/docs"
	"aastar_dashboard_back/env"
	"aastar_dashboard_back/model"
	"aastar_dashboard_back/repository"
	"aastar_dashboard_back/rpc_server/controller"
	"aastar_dashboard_back/rpc_server/controller/oauth"
	"aastar_dashboard_back/rpc_server/middlewares"
	"flag"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var engine *gin.Engine
var aPort = flag.String("port", "", "Port")

func getPort() string {
	flag.Parse()

	if len(*aPort) == 0 {
		*aPort = os.Getenv("port")
	}

	if len(*aPort) == 0 {
		*aPort = ":80"
	}

	if !strings.HasPrefix(*aPort, ":") {
		*aPort = ":" + *aPort
	}

	return *aPort
}

// @contact.name   AAStar BackEndDashBoard
// @contact.url    https://aastar.xyz
// @BasePath /api
// @title AAStar BackEndDashBoard API
// @version v1
// @securityDefinitions.apikey JWT
// @in header
// @name Authorization
// @description Type 'Bearer \<TOKEN\>' to correctly set the AccessToken
// @BasePath /api
func main() {

	engine = gin.New()
	configPath := getConfigPath()
	config.Init(configPath)
	env.Init()
	buildMod(engine)
	buildSwagger(engine)

	engine.Use(middlewares.GenericRecoveryHandler())
	engine.Use(middlewares.LogHandler())
	oauth.Init()
	repository.Init()
	data_view_repository.Init()
	engine.GET("/api/healthz", Healthz)
	buildOAuth()
	buildMid()
	buildRouter()

	engine.NoRoute(func(ctx *gin.Context) {
		model.GetResponse().SetHttpCode(http.StatusNotFound).FailCode(ctx, http.StatusNotFound)
	})

	_ = engine.Run(getPort())
}

func getConfigPath() string {
	configPath := os.Getenv("CONFIG_PATH")
	logrus.Infof("Config Path:[%s]", configPath)
	if configPath == "" {
		configPath = "config/config.json"
	}
	return configPath
}
func buildMid() {
	engine.Use(middlewares.CorsHandler())
	engine.Use(middlewares.AuthHandler())
}
func buildRouter() {

	engine.GET("/api/v1/paymaster_strategy/list", controller.GetStrategyList)
	engine.GET("/api/v1/paymaster_strategy", controller.GetStrategy)
	engine.PUT("/api/v1/paymaster_strategy", controller.UpdateStrategy)
	engine.POST("/api/v1/paymaster_strategy", controller.AddStrategy)
	engine.DELETE("/api/v1/paymaster_strategy", controller.DeleteStrategy)
	engine.PUT("/api/v1/paymaster_strategy/switch_status", controller.SwitchStrategyStatus)
	engine.GET("/api/v1/api_key/list", controller.GetApiKeyList)
	engine.GET("/api/v1/api_key", controller.GetApiKey)
	engine.PUT("/api/v1/api_key", controller.UpdateApiKey)
	engine.DELETE("/api/v1/api_key", controller.DeleteApiKey)
	engine.POST("/api/v1/api_key/apply", controller.ApplyApiKey)

	engine.GET("/api/v1/user", controller.GetUserInfo)

	engine.POST("/api/v1/sponsor/deposit", controller.SponsorDeposit)
	engine.POST("/api/v1/sponsor/withdraw", controller.SponsorWithdraw)
	engine.GET("/api/v1/data/sponsor_transactions", controller.DataViewGetSponsorTransactionList)
	engine.GET("/api/v1/data/paymaster_requests", controller.DataViewApiKeyPaymasterRecallDetailList)
	engine.GET("/api/v1/data/api_detail_data/apikey_request_total_data", controller.DataViewApiKeyPaymasterRecallDetailList)
	engine.GET("/api/v1/data/request_health_list", controller.DataViewRequestHealth)
	engine.GET("/api/v1/data/request_health_one", controller.DataViewRequestHealthOneByApiKey)
	engine.GET("/api/v1/data/balance", controller.DataViewGetBalance)
	engine.GET("/api/v1/data/sponsored_metrics", controller.DataViewSponsoredMetrics)
	engine.GET("/api/v1/data/api_keys_data_overview", controller.DataViewApiKeysOverView)
	engine.GET("/api/v1/data/paymaster_pay_type_metrics", controller.DataViewPaymasterPayTypeMetrics)
}

// buildMod set Mode by Environment
func buildMod(routers *gin.Engine) {

	// prod mode
	if env.Environment.IsProduction() {
		logrus.SetLevel(logrus.InfoLevel)
		gin.SetMode(gin.ReleaseMode)
		gin.DefaultWriter = io.Discard // disable gin log
		logrus.Infof("Build Release Mode")
		return
	}

	// dev mod
	if env.Environment.IsDevelopment() {
		gin.SetMode(gin.DebugMode)
		logrus.SetLevel(logrus.DebugLevel)
		return
	}

}
func buildSwagger(router *gin.Engine) {
	docs.SwaggerInfo.BasePath = "/"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

// buildOAuth supports 3rd party login via OAuth
func buildOAuth() {
	engine.GET("/oauth/github", oauth.GithubOAuthLogin)
	engine.GET("/oauth/email", oauth.EmailOauthLogin)
	engine.POST("/oauth/password", oauth.PasswordOauthLogin)
	engine.POST("/oauth/logOut", oauth.Logout)
}

// Healthz
// @Tags Healthz
// @Description Get Healthz
// @Accept json
// @Product json
// @Router /api/healthz [get]
// @Success 200
func Healthz(c *gin.Context) {
	c.JSON(200, gin.H{
		"hello":   "AAStar Dashboard",
		"time":    time.Now(),
		"version": "v1.0.0",
	})
}
