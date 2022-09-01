package main

import (
	"fmt"
	"xlab-feishu-robot/app"
	"xlab-feishu-robot/config"
	"xlab-feishu-robot/docs"
	"xlab-feishu-robot/plugins"

	"github.com/YasyaKarasu/feishuapi"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	config.ReadConfig()

	// log
	config.SetupLogrus()
	logrus.Info("Robot starts up")

	// feishu api client
	var cli feishuapi.AppClient
	config.SetupFeishuApiClient(&cli)

	cli.StartTokenTimer()

	// robot server
	r := gin.Default()

	app.Init(r)
	plugins.Init(r)

	r.Run(":" + fmt.Sprint(config.C.Server.Port))

	// api docs by swagger
	docs.SwaggerInfo.BasePath = "/"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
