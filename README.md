# Anduin — 企业综合管理平台后端

Anduin 是企业综合管理平台的后端服务，负责业务处理、数据访问及服务端基础能力。项目当前处于基础工程初始化阶段，尚未提供业务 API。

## 技术栈

- Go 1.16（当前开发环境基线；已停止安全维护，生产环境建议升级）
- Gin — HTTP Web 框架
- GORM + MySQL — 数据访问
- Viper — YAML 配置及环境变量
- Uber Zap — 结构化日志

## 快速开始

```bash
git clone https://github.com/3115826227/Anduin.git
cd Anduin
git switch feature_init_project_20261105
go mod tidy
go run .
```

默认监听 `8080` 端口；数据库连接默认关闭，无需先启动 MySQL。尚未注册业务路由，访问未实现路径会返回 404。

配置位于 `configs/config.yaml`，可使用 `ANDUIN_` 前缀的环境变量覆盖，例如 `ANDUIN_SERVER_PORT=9090`。环境变量示例参见 `.env.example`，程序不会自动加载 `.env`。

## 开发与验证

```bash
go fmt ./...
go vet ./...
go test ./...
go build ./...
```

当前依赖需要在 Go 1.16 环境下验证兼容性；请以 CI 检查结果为准。

## 仓库与协作

- 前端：[Garona](https://github.com/3115826227/Garona)
- 文档：[Illidan](https://github.com/3115826227/Illidan)
- `master`：生产稳定分支；`develop`：开发主线
- `feature_<description>_<YYYYMMDD>`：需求开发；`hotfix_<description>_<YYYYMMDD>`：紧急修复
- 功能从 `develop` 拉分支，通过 Pull Request 合并回 `develop`。
- Commit 和 PR 标题使用中英文双语，例如：`docs(readme): add project overview | 补充项目介绍`。

## 当前状态

基础工程建设中，尚未实现业务模块、身份认证及数据库迁移。
