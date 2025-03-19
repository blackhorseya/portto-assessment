package loggerx

import (
	"bytes"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"portto/pkg/contextx"
)

func TestGinTraceLoggingMiddleware(t *testing.T) {
	// 設定自訂的 slog.Handler，將 log 輸出寫入 buffer
	var buf bytes.Buffer
	handler := slog.NewTextHandler(&buf, nil)
	logger := slog.New(handler)

	// 設定 gin 測試模式並建立 gin 引擎
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(GinTraceLoggingMiddleware(logger))

	// 註冊一個測試路由，檢查 context 是否正確注入 logger
	r.GET("/test", func(c *gin.Context) {
		attachedLogger := contextx.GetLogger(c.Request.Context())
		if attachedLogger != logger {
			c.String(http.StatusInternalServerError, "logger not attached correctly")
			return
		}
		c.String(http.StatusOK, "ok")
	})

	// 建立測試請求並執行
	req, err := http.NewRequest("GET", "/test", nil)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// 檢查 HTTP 回應狀態碼
	if w.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, w.Code)
	}

	// 檢查 logger 輸出是否包含預期的 log 訊息
	logOutput := buf.String()
	t.Logf("Logger output:\n%s", logOutput)

	if !strings.Contains(logOutput, "request started") {
		t.Errorf("log output does not contain 'request started': %s", logOutput)
	}
	if !strings.Contains(logOutput, "request completed") {
		t.Errorf("log output does not contain 'request completed': %s", logOutput)
	}
}
