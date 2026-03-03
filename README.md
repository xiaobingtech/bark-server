# Bark Server

<img src="https://wx3.sinaimg.cn/mw690/0060lm7Tly1g0nfnjjxbbj30sg0sg757.jpg" width=200px height=200px />

[![Go Version](https://img.shields.io/badge/Go-1.23+-00ADD8?style=flat&logo=go)](https://golang.org)
[![Docker Automated build](https://img.shields.io/docker/automated/finab/bark-server.svg)](https://hub.docker.com/r/finab/bark-server)
[![Image Size](https://img.shields.io/docker/image-size/finab/bark-server?sort=date)](https://hub.docker.com/r/finab/bark-server)
[![License](https://img.shields.io/github/license/finb/bark-server)](LICENSE)

[Shark](https://github.com/xiaobingtech/Shark) 是一个 鸿蒙 App，用于把自定义通知推送到 鸿蒙手机。本仓库是 Shark 的推送服务端实现，基于 Go + Fiber 构建，具有以下特性：

- **多平台推送支持**：同时支持 iOS APNs 和 HarmonyOS PushKit
- **高性能**：基于 [Fiber](https://gofiber.io/) 框架，使用 fasthttp 引擎
- **灵活存储**：支持 Bbolt（默认）、MySQL、环境变量三种存储模式
- **API 兼容**：V1/V2 双版本 API 兼容，支持批量推送
- **MCP 支持**：提供 MCP (Model Context Protocol) HTTP 端点，可与 AI 工具集成
- **生产就绪**：内置优雅关停、健康检查、Basic Auth 认证
- **容器化**：完整的 Docker/Kubernetes 部署支持

## 目录

- [快速开始](#快速开始)
- [配置参数](#配置参数)
- [API 接口](#api-接口)
- [部署方式](#部署方式)
- [开发指南](#开发指南)
- [项目结构](#项目结构)

## 快速开始

### Docker 部署（推荐）

```sh
# 单容器运行
docker run -dt --name bark -p 8080:8080 -v `pwd`/bark-data:/data xiaobingtech/bark-server:latest
```

```sh
# Docker Compose 部署
mkdir bark-server && cd bark-server
curl -sL https://github.com/xiaobingtech/bark-server/raw/master/deploy/docker-compose.yaml > docker-compose.yaml
docker compose up -d
```

### 二进制运行

```sh
# 1. 从 releases 下载预编译版本
# https://github.com/xiaobingtech/bark-server/releases

# 2. 添加执行权限
chmod +x bark-server

# 3. 启动服务
./bark-server --addr 0.0.0.0:8080 --data ./bark-data

# 4. 验证服务
curl localhost:8080/ping
```

## 配置参数

| 参数 | 默认值 | 说明 |
|------|--------|------|
| `--addr` | `0.0.0.0:8080` | 监听地址 |
| `--data` | `/data` | 数据存储目录（Bbolt 模式） |
| `--dsn` | - | MySQL 连接字符串，格式：`user:pass@tcp(host)/dbname` |
| `--serverless` | `false` | Serverless 模式，使用环境变量存储 |
| `--user` | - | Basic Auth 用户名 |
| `--password` | - | Basic Auth 密码 |
| `--cert` | - | TLS 证书文件路径 |
| `--key` | - | TLS 私钥文件路径 |
| `--max-apns-client-count` | `10` | APNs 客户端连接池大小 |
| `--max-batch-push-count` | `100` | 批量推送数量限制 |

## API 接口

### 健康检查

| 端点 | 方法 | 说明 |
|------|------|------|
| `/` | GET | 健康检查 |
| `/ping` | GET | 返回 `pong` |
| `/healthz` | GET | 健康检查 |
| `/info` | GET | 服务器版本信息 |

### 设备注册

| 端点 | 方法 | 说明 |
|------|------|------|
| `/register` | POST | 注册设备，返回 `device_key` |
| `/register/:device_key` | GET | 检查设备是否已注册 |

**注册请求示例：**
```json
{
  "device_token": "xxxxxxxxxxxxxx",
  "platform": "ios"  // 可选值: ios, harmony
}
```

### 推送通知

#### V2 API（推荐）

```sh
POST /push
Content-Type: application/json

{
  "device_key": "your_device_key",
  "title": "通知标题",
  "body": "通知内容",
  "sound": "bell",
  "badge": 1,
  "icon": "https://example.com/icon.png",
  "group": "default",
  "url": "https://example.com"
}
```

#### V1 API（兼容）

```sh
# 基础推送
GET /:device_key/:body

# 带标题推送
GET /:device_key/:title/:body

# POST 表单推送
POST /:device_key
Content-Type: application/x-www-form-urlencoded

title=标题&body=内容&sound=bell
```

#### 批量推送

```sh
POST /push
Content-Type: application/json

[
  {"device_key": "key1", "title": "标题1", "body": "内容1"},
  {"device_key": "key2", "title": "标题2", "body": "内容2"}
]
```

### MCP 端点

MCP (Model Context Protocol) 支持 AI 工具调用推送：

| 端点 | 方法 | 说明 |
|------|------|------|
| `/mcp` | ALL | MCP 通用端点，需在参数中提供 `device_key` |
| `/mcp/:device_key` | ALL | 设备专用端点 |

详细文档请参考 [MCP.md](docs/MCP.md)。

## 部署方式

### MySQL 模式

适用于分布式部署场景：

```sh
./bark-server --dsn=user:password@tcp(mysql-host:3306)/bark
```

### Kubernetes 部署

提供完整的 Helm Chart：

```sh
cd deploy/helm-chart
helm install bark-server . -n bark --create-namespace
```

### systemd 服务

```sh
# 复制服务文件
sudo cp deploy/bark-server.service /etc/systemd/system/

# 启动服务
sudo systemctl daemon-reload
sudo systemctl enable bark-server
sudo systemctl start bark-server
```

## 开发指南

### 环境要求

- Go 1.23+
- [go-task](https://taskfile.dev/installation/) 构建工具

### 本地开发

```sh
# 克隆仓库
git clone https://github.com/xiaobingtech/bark-server.git
cd bark-server

# 设置代理（可选）
export GOPROXY=https://goproxy.cn,direct

# 运行服务
go run . --addr :8080 --data ./data

# 运行测试
go test -v
```

### 构建发布

```sh
# 安装 go-task
# macOS: brew install go-task
# Windows: scoop install task
# Linux: sh -c "$(curl --location https://taskfile.dev/install.sh)"

# 构建所有平台
task

# 构建特定平台
task linux_amd64
task linux_amd64_v3
task darwin_arm64
task windows_amd64.exe
```

### 支持的平台

| OS | Architectures |
|----|---------------|
| Linux | 386, amd64, amd64_v2/v3/v4, arm, armv5/v6/v7/v8, mips, mips64 |
| Windows | 386, amd64, amd64_v2/v3/v4 |
| macOS | amd64, arm64 |
| FreeBSD | amd64, arm64 |

## 项目结构

```
bark-server/
├── main.go                 # 程序入口
├── router.go               # 路由注册框架
├── route_auth.go           # Basic Auth 认证
├── route_misc.go           # 健康检查等辅助接口
├── route_register.go       # 设备注册接口
├── route_push.go           # 推送接口（V1/V2/批量）
├── route_mcp.go            # MCP HTTP 端点
├── util.go                 # 工具函数
│
├── apns/                   # Apple Push Notification Service
│   ├── apns.go             # APNs 客户端池与推送实现
│   └── apns_certs.go       # APNs 私钥与根证书
│
├── database/               # 数据库抽象层
│   ├── database.go         # 数据库接口定义
│   ├── bbolt.go            # Bbolt 嵌入式数据库
│   ├── mysql.go            # MySQL 实现
│   ├── membase.go          # 内存数据库（测试用）
│   └── envbase.go          # Serverless 环境变量存储
│
├── harmony/                # HarmonyOS 推送支持
│   └── harmony.go          # 华为 PushKit 推送实现
│
├── deploy/                 # 部署相关文件
│   ├── Dockerfile
│   ├── docker-compose.yaml
│   ├── entrypoint.sh
│   ├── bark-server.service
│   └── helm-chart/         # Kubernetes Helm Chart
│
├── docs/                   # 文档
│   ├── API_V2.md           # V2 API 详细文档
│   └── MCP.md              # MCP 说明文档
│
└── .github/workflows/      # GitHub Actions CI/CD
    └── ci.yaml             # 自动构建推送 Docker 镜像
```

## 技术栈

| 组件 | 技术 |
|------|------|
| Web 框架 | [Fiber v2](https://gofiber.io/) |
| APNs 客户端 | [apns2](https://github.com/sideshow/apns2) |
| MCP 协议 | [mcp-go](https://github.com/mark3labs/mcp-go) |
| 嵌入式数据库 | [Bbolt](https://github.com/etcd-io/bbolt) |
| MySQL 驱动 | [go-sql-driver/mysql](https://github.com/go-sql-driver/mysql) |
| CLI 框架 | [urfave/cli](https://github.com/urfave/cli) |
| JSON 编解码 | [json-iterator](https://github.com/json-iterator/go) |

## 工作流程

```
┌─────────────┐     ┌─────────────┐     ┌─────────────┐
│   Client    │────▶│  Bark Server│────▶│  APNs/PushKit│
│  (Bark App) │     │   (Go+Fiber)│     │   (Apple/HW) │
└─────────────┘     └─────────────┘     └─────────────┘
                           │
                           ▼
                    ┌─────────────┐
                    │  Database   │
                    │ (Bbolt/MySQL)│
                    └─────────────┘
```

1. **启动**：解析命令行参数，创建 Fiber 实例，初始化数据库
2. **认证**：可选启用 Basic Auth 中间件
3. **注册**：设备通过 `/register` 注册，获取 `device_key`
4. **推送**：通过 `/push` 或兼容路由发送通知，查询 `device_token` 并调用 APNs/PushKit
5. **MCP**：AI 工具通过 `/mcp` 端点调用推送功能
6. **关停**：优雅关闭 Fiber 与数据库连接

## 文档

- [API V2 文档](docs/API_V2.md)
- [MCP 集成文档](docs/MCP.md)

## 相关项目

- [Bark iOS App](https://github.com/Finb/Bark) - iOS 客户端
- [Bark](https://github.com/sunbinyuan/Bark) - 原版服务端

## License

[MIT License](LICENSE)
