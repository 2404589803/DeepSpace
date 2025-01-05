# DeepSpace

<div align="center">
<img src=assets/1.jpg width="200px" style="border-radius: 50%"/>
</div>

[简体中文](README_cn.md)

DeepSpace is a powerful debugging and monitoring tool for the DeepSeek API, developed by the community and inspired by [MoonPalace](https://github.com/MoonshotAI/moonpalace?tab=readme-ov-file).

## Features

- **Cross-platform Support:**
  - [x] Mac (Intel & Apple Silicon)
  - [x] Windows
  - [x] Linux
- **User-friendly Interface**: 
  - Simple setup: Just replace `base_url` with `http://localhost:9988` after starting the service
  - Beautiful console output with color-coded information
  - Intuitive command-line interface
- **Comprehensive Request Monitoring**:
  - Complete request and response capture
  - Detailed error tracking and debugging information
  - Network failure analysis and diagnostics
- **Advanced Search & Analysis**:
  - Quick retrieval using `request_id` or `chatcmpl_id`
  - Filtering and sorting capabilities
  - Request history tracking
- **Professional Debugging Tools**:
  - One-click export of structured BadCase reports
  - Request timing analysis
  - Performance metrics tracking
- **Security & Privacy**:
  - Local proxy server
  - No data sent to external servers
  - Full control over your API interactions

## Installation

### Method 1: Using the `go` Command (Recommended)

If you have the Go toolchain installed, install DeepSpace with:

```shell
$ go install github.com/2404589803/deepspace@latest
```

The binary will be installed in your `$GOPATH/bin/` directory.

### Method 2: Direct Download

Download pre-compiled binaries from [Releases](https://github.com/2404589803/deepspace/releases):

- `deepspace-linux` - For Linux systems
- `deepspace-macos-amd64` - For Intel-based Macs
- `deepspace-macos-arm64` - For Apple Silicon Macs
- `deepspace-windows.exe` - For Windows systems

## Usage Guide

### 1. Starting the Service

```shell
$ deepspace start [--port <PORT>] [--host <HOST>]
```

Options:
- `--port`: Specify listening port (default: 9988)
- `--host`: Specify host address (default: 127.0.0.1)

### 2. Querying Requests

```shell
$ deepspace list [--limit <N>] [--offset <N>] [--id <REQUEST_ID>]
```

### 3. Inspecting Request Details

```shell
$ deepspace inspect <REQUEST_ID>
```

### 4. Exporting Request Data

```shell
$ deepspace export <REQUEST_ID> [--output <FILE>]
```

### 5. Cleaning Up Data

```shell
$ deepspace cleanup [--before <DATE>]
```

## Advanced Features

### Request Filtering

You can filter requests using various criteria:
- By date range
- By status code
- By response time
- By request type

### Performance Analysis

DeepSpace provides detailed timing information:
- Total request duration
- Time to first token (TTFT)
- Token processing speed
- Network latency

### Error Handling

When errors occur, DeepSpace captures:
- Full error stack traces
- Network conditions
- Request/response headers
- System state information

## Contributing

We welcome contributions! Please feel free to:
- Submit bug reports
- Propose new features
- Create pull requests
- Improve documentation

## Support

For bug reports and feature requests, please use the GitHub issues system.

For BadCase reports and API-related issues, you can submit exported files to:
[api-service@deepseek.com](mailto:api-service@deepseek.com)

## License

This project is licensed under the MIT License - see the LICENSE file for details.