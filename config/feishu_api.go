package config

import "github.com/YasyaKarasu/feishuapi"

func SetupFeishuApiClient(cli *feishuapi.AppClient) {
	cli.Conf = C.Feishu
}
