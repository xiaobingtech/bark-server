# Bark

<img src="https://wx3.sinaimg.cn/mw690/0060lm7Tly1g0nfnjjxbbj30sg0sg757.jpg" width=200px height=200px />

[Bark](https://github.com/Finb/Bark) 是一个 iOS App，用于把自定义通知推送到 iPhone。本仓库是 Bark 的推送服务端实现，基于 Go + Fiber，内置 APNs 推送能力，并支持传统 API 与 MCP 调用方式。

## 功能逻辑

1. 进程启动后由 main.go 解析命令行参数并创建 Fiber 实例  
2. 通过 routerAuth 启用可选 Basic Auth，routerSetup 挂载所有路由  
3. 根据参数选择数据库实现：Bbolt 默认、本地 MySQL、或 Serverless 环境变量模式  
4. /register 负责注册 device_key 与 device_token  
5. /push 或兼容路由根据参数组装 PushMessage，查询 device_token 并调用 APNs  
6. /mcp 提供 MCP HTTP 端点，将工具调用转为内部 push  
7. 退出时触发优雅关停，关闭 Fiber 与数据库连接

## 路由与行为概览

- /register：注册设备并返回 device_key  
- /register/:device_key：检查设备是否已注册  
- /push：V2 JSON 推送与批量推送  
- /:device_key 等兼容路由：V1 查询/表单/路径参数推送  
- /mcp 与 /mcp/:device_key：MCP Streamable HTTP 端点  
- /ping、/healthz、/info：健康检查与版本信息

## 安装与运行

### Docker

![Docker Automated build](https://img.shields.io/docker/automated/finab/bark-server.svg) ![Image Size](https://img.shields.io/docker/image-size/finab/bark-server?sort=date) ![License](https://img.shields.io/github/license/finb/bark-server)

``` sh
docker run -dt --name bark -p 8080:8080 -v `pwd`/bark-data:/data finab/bark-server
```

``` sh
docker run -dt --name bark -p 8080:8080 -v `pwd`/bark-data:/data ghcr.io/finb/bark-server
```

``` sh
mkdir bark-server && cd bark-server
curl -sL https://github.com/Finb/bark-server/raw/master/deploy/docker-compose.yaml > docker-compose.yaml
docker compose up -d
```

### 二进制运行

- 1、从 [releases](https://github.com/Finb/bark-server/releases) 下载预编译版本  
- 2、`chmod +x bark-server`  
- 3、`./bark-server --addr 0.0.0.0:8080 --data ./bark-data`  
- 4、`curl localhost:8080/ping`  

默认数据目录是 /data，可用 -d 参数调整路径。

### 开发者编译

- Golang 1.18+  
- GO111MODULE=on  
- GOPROXY=https://goproxy.cn  
- 安装 [go-task](https://taskfile.dev/installation/)  

```sh
task
task linux_amd64
task linux_amd64_v3
```

### MySQL 模式

使用 `-dsn=user:pass@tcp(mysql_host)/bark` 启用 MySQL 存储。

## 目录与文件说明

**根目录文件**
- README.md：项目说明  
- .gitignore：Git 忽略规则  
- LICENSE：授权协议  
- Procfile：PaaS 启动配置  
- Taskfile.yaml：go-task 构建与跨平台打包脚本  
- app.json：Heroku/平台配置  
- go.mod：Go 模块依赖  
- go.sum：Go 依赖校验  
- main.go：程序入口，参数解析、初始化与启动  
- router.go：路由注册、通用响应结构、统一中间件  
- route_auth.go：Basic Auth 认证逻辑  
- route_misc.go：根路径、健康检查、版本信息  
- route_register.go：设备注册与校验  
- route_push.go：V1/V2 推送、批量推送、参数解析  
- route_mcp.go：MCP Streamable HTTP 端点与工具定义  
- util.go：随机字符串工具  
- push_test.go：推送与注册的集成测试  

**apns/**
- apns.go：APNs 客户端池与推送实现  
- apns_certs.go：APNs 私钥与根证书集合  

**database/**
- database.go：数据库接口定义  
- bbolt.go：Bbolt 实现与本地文件存储  
- mysql.go：MySQL 实现与 TLS 连接  
- membase.go：内存数据库实现（测试使用）  
- envbase.go：Serverless 环境变量实现  

**docs/**
- API_V2.md：V2 API 文档  
- MCP.md：MCP 说明  

**.github/workflows/**
- ci.yaml：标签发布时构建并推送多架构镜像  

**deploy/**
- Dockerfile：镜像构建  
- docker-compose.yaml：本地容器运行  
- entrypoint.sh：容器入口脚本  
- bark-server.service：systemd 服务  
- AuthKey_LH4T9V5U4R_5U8LBRXG3A.p8：APNs AuthKey  
- AuthKey_LH4T9V5U4R_5U8LBRXG3A.pem：APNs 私钥 PEM  

**deploy/helm-chart/**
- Chart.yaml：Helm Chart 元数据  
- values.yaml：默认配置  
- .helmignore：打包忽略规则  
- templates/_helpers.tpl：Helm 模板函数  
- templates/deployment.yaml：Deployment 模板  
- templates/service.yaml：Service 模板  
- templates/ingress.yaml：Ingress 模板  
- templates/hpa.yaml：HPA 模板  
- templates/serviceaccount.yaml：ServiceAccount 模板  
- templates/NOTES.txt：安装后提示信息  
- templates/tests/test-connection.yaml：Helm 测试连接  

## 文档入口

- [API_V2.md](docs/API_V2.md)
- [MCP.md](docs/MCP.md)
