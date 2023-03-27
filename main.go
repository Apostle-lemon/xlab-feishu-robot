package main

import (
	"fmt"
	"xlab-feishu-robot/docs"
	config "xlab-feishu-robot/internal/config"
	"xlab-feishu-robot/internal/log"

	"xlab-feishu-robot/internal/pkg"

	"xlab-feishu-robot/internal"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Init() {

}

func main() {
	config.ReadConfig()

	// log
	log.SetupLogrus()
	logrus.Info("Robot starts up")

	// feishu api client
	config.SetupFeishuApiClient(&pkg.Cli)
	pkg.Cli.StartTokenTimer()

	// robot server
	r := gin.Default()
	internal.Init(r)

	// api docs by swagger
	docs.SwaggerInfo.BasePath = "/"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.Run(":" + fmt.Sprint(config.C.Server.Port))
}
