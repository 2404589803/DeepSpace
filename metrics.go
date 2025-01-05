package main

import (
	"net/http"
	"time"
)

// 简单的计数器实现
type Counter struct {
	value int64
}

func (c *Counter) Inc() {
	c.value++
}

func (c *Counter) Add(delta int64) {
	c.value += delta
}

// 简单的监控指标
var (
	requestsTotal   = &Counter{}
	activeRequests  = &Counter{}
	errorCount      = &Counter{}
	totalTokenCount = &Counter{}
)

// 记录请求
func recordRequest(method, path string, status int, duration time.Duration) {
	requestsTotal.Inc()
}

// 记录令牌数量
func recordTokens(tokenType string, count int) {
	totalTokenCount.Add(int64(count))
}

// 记录错误
func recordError(errorType string) {
	errorCount.Inc()
}

// 更新活跃请求数
func updateActiveRequests(delta int) {
	activeRequests.Add(int64(delta))
}

// 启动指标服务器
func startMetricsServer(config *Config) {
	if !config.EnableMetrics {
		return
	}

	go func() {
		mux := http.NewServeMux()
		mux.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "text/plain")
			w.Write([]byte(`
# HELP deepspace_requests_total Total number of requests processed
# TYPE deepspace_requests_total counter
deepspace_requests_total ` + string(requestsTotal.value) + `

# HELP deepspace_active_requests Current number of active requests
# TYPE deepspace_active_requests gauge
deepspace_active_requests ` + string(activeRequests.value) + `

# HELP deepspace_errors_total Total number of errors
# TYPE deepspace_errors_total counter
deepspace_errors_total ` + string(errorCount.value) + `

# HELP deepspace_tokens_total Total number of tokens processed
# TYPE deepspace_tokens_total counter
deepspace_tokens_total ` + string(totalTokenCount.value) + `
`))
		})
		http.ListenAndServe(":"+string(config.MetricsPort), mux)
	}()
}
