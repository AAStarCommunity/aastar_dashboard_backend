package main

import (
	"aastar_dashboard_back/config"
	"aastar_dashboard_back/docs"
	"aastar_dashboard_back/repository"
	"aastar_dashboard_back/rpc_server/controller"
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"io"
	"os"
	"strings"
	"time"
)

var Engine *gin.Engine
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

func main() {

	Engine = gin.New()

	gin.SetMode(gin.DebugMode)
	logrus.SetLevel(logrus.DebugLevel)
	buildSwagger(Engine)
	config.Init("config/config.json")
	logrus.Infof("Config loaded successfully Env: %s", config.Environment.Name)
	logrus.Infof("System Config: %v", config.SystemConfigViper.AllSettings())
	dsn := config.GetDsn()
	logrus.Infof("DSN : %s", dsn)
	repository.Init(dsn)
	//DB Init
	buildMod(Engine)
	buildRouter()
	_ = Engine.Run(getPort())
}
func buildRouter() {
	Engine.GET("/api/healthz", Healthz)

	Engine.GET("/api/v1/paymaster_strategy/list", controller.GetStrategyList)
	Engine.GET("/api/v1/paymaster_strategy", controller.GetStrategy)
	Engine.PUT("/api/v1/paymaster_strategy", controller.UpdateStrategy)
	Engine.POST("/api/v1/paymaster_strategy", controller.AddStrategy)
	Engine.DELETE("/api/v1/paymaster_strategy", controller.DeleteStrategy)

	Engine.GET("/api/v1/api_key/list", controller.GetApiKeyList)
	Engine.GET("/api/v1/api_key", controller.GetApiKey)
	Engine.PUT("/api/v1/api_key", controller.UpdateApiKey)
	Engine.DELETE("/api/v1/api_key", controller.DeleteApiKey)
	Engine.POST("/api/v1/api_key", controller.AddApiKey)

}

// buildMod set Mode by Environment
func buildMod(routers *gin.Engine) {

	// prod mode
	if config.Environment.IsProduction() {
		gin.SetMode(gin.ReleaseMode)
		gin.DefaultWriter = io.Discard // disable gin log
		return
	}

	// dev mod
	if config.Environment.IsDevelopment() {
		gin.SetMode(gin.DebugMode)

		return
	}
}
func buildSwagger(router *gin.Engine) {
	docs.SwaggerInfo.BasePath = "/"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

// Healthz
// @Tags Healthz
// @Description Get Healthz
// @Accept json
// @Product json
// @Router /api/healthz [get]
// @Success 200
func Healthz(c *gin.Context) {
	logrus.Debug("In the Healthz")
	c.JSON(200, gin.H{
		"hello":   "AAStar Dashboard",
		"time":    time.Now(),
		"version": "v1.0.0",
	})
}
