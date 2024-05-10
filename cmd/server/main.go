package main

import (
	"aastar_dashboard_back/docs"
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
	docs.SwaggerInfo.BasePath = "/"

	Engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	Engine.GET("/health", Healthz)
	_ = Engine.Run(getPort())
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
