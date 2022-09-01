package config

import "github.com/YasyaKarasu/feishuapi"

type Config struct {
	Feishu feishuapi.Config
	Server struct {
		Port int
	}
}
