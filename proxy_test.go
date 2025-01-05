package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestProxyServer(t *testing.T) {
	// 创建测试服务器
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := map[string]interface{}{
			"id":      "chatcmpl-123",
			"object":  "chat.completion",
			"created": 1677652288,
			"choices": []map[string]interface{}{
				{
					"index": 0,
					"message": map[string]interface{}{
						"role":    "assistant",
						"content": "Hello, how can I help you today?",
					},
				},
			},
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer ts.Close()

	// 创建测试请求
	requestBody := map[string]interface{}{
		"model":    "deepseek-chat",
		"messages": []map[string]string{{"role": "user", "content": "Hello"}},
	}
	body, _ := json.Marshal(requestBody)

	req := httptest.NewRequest("POST", "/v1/chat/completions", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	// 创建响应记录器
	w := httptest.NewRecorder()

	// 检查响应状态码
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}
}

func TestErrorHandling(t *testing.T) {
	tests := []struct {
		name           string
		serverResponse int
		serverBody     string
		expectedError  bool
	}{
		{
			name:           "成功请求",
			serverResponse: http.StatusOK,
			serverBody:     `{"id": "test-123"}`,
			expectedError:  false,
		},
		{
			name:           "服务器错误",
			serverResponse: http.StatusInternalServerError,
			serverBody:     `{"error": "internal server error"}`,
			expectedError:  true,
		},
		{
			name:           "无效响应",
			serverResponse: http.StatusOK,
			serverBody:     `invalid json`,
			expectedError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 创建测试服务器
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tt.serverResponse)
				io.WriteString(w, tt.serverBody)
			}))
			defer ts.Close()

			// 创建测试请求
			req := httptest.NewRequest("POST", "/v1/chat/completions", strings.NewReader("test"))
			w := httptest.NewRecorder()

			// 验证错误处理
			if tt.expectedError && w.Code == http.StatusOK {
				t.Error("Expected error response, got success")
			}
			if !tt.expectedError && w.Code != http.StatusOK {
				t.Errorf("Expected success response, got error: %d", w.Code)
			}
		})
	}
}
