# XLab 飞书机器人模板

| Web Framework | Log Manager     | Config Manager | Api Documentation  | Feishu Api Client     |
|:-------------:|:---------------:|:--------------:|:------------------:|:---------------------:|
| gin-gonic/gin | sirupsen/logrus | spf13/viper    | swaggo/gin-swagger | YasyaKarasu/feishuapi |

## Usage

- `internal/router` 为自定义的 service controller 注册路由
- `internal/dispatcher` 将飞书发送的 event 分发到对应的 event handler
- `internal/event_handler` 自定义飞书事件处理方法
- `internal/controller` 自定义 service controller
- `internal/config` 在 `Config` 类型定义中添加自定义的配置字段
- `config/config.yaml` 添加自定义的配置字段

## Architecture

- `internal` 机器人主体部分
- `/config` 机器人配置
- `docs` swagger 生成的 Api 文档
- `docs/devel` 开发文档
- `docs/guide` 使用文档
- `deployments` 部署相关