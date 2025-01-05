# DeepSpace

<div align="center">
<img src=assets/1_circle.png width="150"/>
</div>

DeepSpace（深度空间）是一款功能强大的 DeepSeek API 调试和监控工具，由社区制作，灵感来源于 [MoonPalace](https://github.com/MoonshotAI/moonpalace?tab=readme-ov-file)。

[Read this in English](README.md)

## 特性

- **全平台支持：**
  - [x] Mac（支持 Intel 和 Apple Silicon）
  - [x] Windows
  - [x] Linux
- **用户友好界面**：
  - 简单配置：启动后只需将 `base_url` 替换为 `http://localhost:9988`
  - 美观的控制台输出，支持彩色信息显示
  - 直观的命令行界面
- **全面的请求监控**：
  - 完整的请求和响应捕获
  - 详细的错误跟踪和调试信息
  - 网络故障分析和诊断
- **高级搜索和分析**：
  - 通过 `request_id` 或 `chatcmpl_id` 快速检索
  - 强大的过滤和排序功能
  - 请求历史记录追踪
- **专业调试工具**：
  - 一键导出结构化 BadCase 报告
  - 请求时间分析
  - 性能指标追踪
- **安全和隐私**：
  - 本地代理服务器
  - 不向外部服务器发送数据
  - 完全控制您的 API 交互

## 安装方式

### 方式一：使用 `go` 命令安装（推荐）

如果您已安装 Go 开发环境，可以使用以下命令安装：

```shell
$ go install github.com/2404589803/deepspace@latest
```

该命令会在您的 `$GOPATH/bin/` 目录中安装可执行文件。

### 方式二：直接下载

从 [Releases](https://github.com/2404589803/deepspace/releases) 页面下载预编译的二进制文件：

- `deepspace-linux` - 适用于 Linux 系统
- `deepspace-macos-amd64` - 适用于 Intel 芯片的 Mac
- `deepspace-macos-arm64` - 适用于 Apple Silicon 芯片的 Mac
- `deepspace-windows.exe` - 适用于 Windows 系统

## 使用指南

### 1. 启动服务

```shell
$ deepspace start [--port <端口>] [--host <主机地址>]
```

选项说明：
- `--port`：指定监听端口（默认：9988）
- `--host`：指定主机地址（默认：127.0.0.1）

### 2. 查询请求

```shell
$ deepspace list [--limit <数量>] [--offset <偏移>] [--id <请求ID>]
```

### 3. 检查请求详情

```shell
$ deepspace inspect <请求ID>
```

### 4. 导出请求数据

```shell
$ deepspace export <请求ID> [--output <文件名>]
```

### 5. 清理数据

```shell
$ deepspace cleanup [--before <日期>]
```

## 高级功能

### 请求过滤

您可以使用多种条件过滤请求：
- 按日期范围
- 按状态码
- 按响应时间
- 按请求类型

### 性能分析

DeepSpace 提供详细的时间信息：
- 总请求时长
- 首个 Token 生成时间（TTFT）
- Token 处理速度
- 网络延迟

### 错误处理

当发生错误时，DeepSpace 会捕获：
- 完整的错误堆栈跟踪
- 网络状况
- 请求/响应头信息
- 系统状态信息

## 参与贡献

我们欢迎各种形式的贡献：
- 提交 Bug 报告
- 提出新功能建议
- 创建 Pull Request
- 改进文档

## 支持

如需报告 Bug 或请求新功能，请使用 GitHub 的 Issues 系统。

对于 BadCase 报告和 API 相关问题，您可以将导出的文件发送至：
[api-service@deepseek.com](mailto:api-service@deepseek.com)

## 许可证

本项目采用 MIT 许可证 - 详见 LICENSE 文件