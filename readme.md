# Quanfuxia 🔧

一套现代化、企业级的 Go 语言后端服务骨架，基于 [Gin](https://github.com/gin-gonic/gin) 框架构建，集成配置管理、多语言支持、JWT 鉴权、日志、数据库/缓存/RabbitMQ 初始化、代码生成器、分层解耦等最佳实践。

---

## 🧱 项目结构简介

quanfuxia/
├── cmd/ // 子命令入口：HTTP服务、worker、消息消费者等
├── configs/ // 多环境配置文件
├── internal/ // 内部业务模块（禁止外部引用）
│ ├── api/ // 请求入口层，按领域划分（如 user）
│ ├── service/ // 核心业务逻辑层
│ ├── domain/ // 领域模型预留
│ ├── repository/ // 数据访问层，支持多种数据源（MySQL/Redis）
│ ├── model/ // 结构体定义，含 gorm 生成的 model
│ ├── route/ // 路由注册
│ ├── job/ // 消息队列消费任务
│ ├── middleware/ // Gin 中间件
│ └── common/ // 响应封装、错误码、多语言、JWT 等通用工具
├── pkg/ // 第三方工具封装（viper/zap/mysql/redis/mq 等）
├── logs/ // 日志目录
├── scripts/ // 可执行脚本或 SQL 初始化
├── go.mod / go.sum
└── main.go // 应用入口，执行 RootCmd

---

## ✨ 技术栈与特性

| 类别       | 技术与说明                                       |
|------------|--------------------------------------------------|
| Web框架     | [Gin](https://github.com/gin-gonic/gin)         |
| 配置管理     | [Viper](https://github.com/spf13/viper)        |
| 路由控制     | 分组路由 + 中间件注册                          |
| ORM层       | [GORM](https://gorm.io/) + `gorm/gen` 代码生成 |
| 日志系统     | [Zap](https://github.com/uber-go/zap) + 分割日志 |
| 参数校验     | validator.v10 + 多语言错误提示支持              |
| 权限系统     | JWT 认证 + RefreshToken + Redis 状态管理       |
| 消息队列     | RabbitMQ（预留接入）                            |
| 多语言支持   | universal-translator + 自定义语言包加载         |
| 命令行工具   | Cobra（支持 `serve` / `gen` 命令）              |
| 容器支持     | Dockerfile（预留）                              |

---

## 🚀 快速开始

### 1. 克隆项目

```bash
git clone https://github.com/quanfuxia888/goskin.git
cd goskin
go mod tidy
```
### 2. 配置文件
默认配置路径为 configs/config.yaml，你可通过 --config 参数指定其他配置：


### 启动服务
```bash
go run main.go serve
```
### 生成 GORM 模型代码
```bash
# 生成所有表
go run main.go gen

# 指定表
go run main.go gen --tables=wa_user,wa_order
```
### 接口示例
1. 用户注册
POST /api/user/register
{
  "username": "test",
  "password": "123456"
}

2. 登录
POST /api/user/login
→ 返回 access_token、refresh_token、过期时间戳

3. 刷新 Token
POST /api/user/refresh
{
  "refresh_token": "xxx"
}

RefreshToken 绑定唯一 JTI（UUID）

Redis 状态管理：防止重复刷新、支持登出作废

多语言支持登录验证、Token 错误提示