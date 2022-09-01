package main

import (
	"fmt"
	"xlab-feishu-robot/app"
	"xlab-feishu-robot/config"
	"xlab-feishu-robot/docs"
	"xlab-feishu-robot/plugins"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	config.ReadConfig()
	config.SetupLogrus()

	docs.SwaggerInfo.BasePath = "/"

	logrus.Info("Robot starts up")

	r := gin.Default()

	app.Init(r)
	plugins.Init(r)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(":" + fmt.Sprint(config.C.Server.Port))
}
