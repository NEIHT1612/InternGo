package logger

import (
	"bytes"
	"io"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GinZapLogger(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// Read the request body
		var reqBody []byte
		if c.Request.Body != nil {
			reqBody, _ = io.ReadAll(c.Request.Body)
			c.Request.Body = io.NopCloser(bytes.NewBuffer(reqBody))
		}

		// Record the response body
		bodyWriter := &responseBodyWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = bodyWriter

		// Execute handler
		c.Next()

		// Calculate latency
		latency := time.Since(start)

		// Log the details
		logger.Info("API Request",
			zap.String("time", time.Now().Format("02/01/2006 15:04:05")),
			zap.String("client_ip", c.ClientIP()),
			zap.String("method", c.Request.Method),
			zap.String("path", "http://" + c.Request.Host + c.Request.RequestURI),
			zap.String("query", c.Request.URL.RawQuery),
			zap.Int("status", c.Writer.Status()),
			zap.String("request_body", string(reqBody)),
			zap.String("response_body", bodyWriter.body.String()),
			zap.String("user_agent", c.Request.UserAgent()),
			zap.Duration("latency", latency),
		)
	}
}

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w *responseBodyWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}
