package main

import (
	"xlab-feishu-robot/app"
	"xlab-feishu-robot/config"
	"xlab-feishu-robot/plugins"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	config.ReadConfig()
	config.SetupLogrus()

	logrus.Info("Robot starts up")

	r := gin.Default()

	app.Init(r)
	plugins.Init(r)

	r.Run("127.0.0.1:" + viper.GetString("server.PORT"))
}
