package config

import (
	"github.com/YasyaKarasu/feishuapi"
	lark "github.com/larksuite/oapi-sdk-go/v3"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"xlab-feishu-robot/internal/pkg"
)

type Config struct {
	Feishu feishuapi.Config
	Server struct {
		Port int

		// add your configuration fields here
		ExampleField1 string
	}

	// add your configuration fields here
	ExampleField2 struct {
		ExampleField3 int
	}
}

var C Config

func ReadConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath("./config/")

	if err := viper.ReadInConfig(); err != nil {
		logrus.Panic(err)
	}

	if err := viper.Unmarshal(&C); err != nil {
		logrus.Error("Failed to unmarshal config")
	}

	logrus.Info("Configuration file loaded")
}

func SetupFeishuApiClient() {
	// WithEnableTokenCache(true): 自动获取、缓存tenant_access_token
	pkg.Cli = lark.NewClient(C.Feishu.AppId, C.Feishu.AppSecret, lark.WithEnableTokenCache(true))
	// Ref:
	// - tenant_access_token: https://open.feishu.cn/document/server-docs/api-call-guide/calling-process/get-access-token
	// - API Client: https://github.com/larksuite/oapi-sdk-go/blob/v3_main/README.md#%E9%85%8D%E7%BD%AEapi-client
}
