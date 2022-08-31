package main

import (
	robotConfig "xlab-feishu-robot/config"
	// feishuApi "xlab-feishu-robot/feishu_api"
	robotLog "xlab-feishu-robot/log"
	robotServer "xlab-feishu-robot/robot_server"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	robotLog.Setup()
	logrus.Info("Robot starts up")

	robotConfig.Init()

	// feishuApi.Client.StartTokenTimer()

	router := gin.Default()
	robotServer.Register(router)

	router.Run("127.0.0.1:" + viper.GetString("server.PORT"))
}
