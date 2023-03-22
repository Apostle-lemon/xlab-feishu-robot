# XLab 飞书机器人模板

| Web Framework | Log Manager     | Config Manager | Api Documentation  | Feishu Api Client     |
|:-------------:|:---------------:|:--------------:|:------------------:|:---------------------:|
| gin-gonic/gin | sirupsen/logrus | spf13/viper    | swaggo/gin-swagger | YasyaKarasu/feishuapi |

## Usage

- `app/event_handler` 自定义飞书事件处理方法
- `app/controller` 自定义 service controller
- `app/router` 为自定义的 service controller 注册路由
- `config` 在 `Config` 类型定义中添加自定义的配置字段
- `config.yaml` 添加自定义的配置字段

## Architecture

- `app` 机器人主体部分
- `config` 机器人配置
- `docs` swagger 生成的 Api 文档
- `pkg/dispatcher` 飞书事件调度器
