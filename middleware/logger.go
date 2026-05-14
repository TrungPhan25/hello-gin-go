package middleware

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func LoggerMiddleware() gin.HandlerFunc {
	logPath := "logs/http.log"

	// create folder if not exist
	if error := os.MkdirAll(filepath.Dir(logPath), os.ModePerm); error != nil {
		panic(error)
	}

	// create log file if not exist, and open it for appending
	logFile, error := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if error != nil {
		panic(error)
	}

	logger := zerolog.New(logFile).With().Timestamp().Logger()

	return func(c *gin.Context) {
		start := time.Now()

		// pull
		bodyByte, err := io.ReadAll(c.Request.Body)

		if err != nil {
			logger.Error().Err(err).Msg("Failed to read request body")
		}

		// restore
		c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyByte))
		// fmt.Println(string(bodyByte))
		c.Next()
		duration := time.Since(start)

		statusCode := c.Writer.Status()

		logEvent := logger.Info()

		if statusCode >= 500 {
			logEvent = logger.Error()
		} else if statusCode >= 400 {
			logEvent = logger.Warn()
		}

		logEvent.
			Str("method", c.Request.Method).
			Str("path", c.Request.URL.Path).
			Str("query", c.Request.URL.RawQuery).
			Str("client_ip", c.ClientIP()).
			Str("user_agent", c.Request.UserAgent()).
			Str("referer", c.Request.Referer()).
			Str("protocol", c.Request.Proto).
			Int64("duration_ms", duration.Milliseconds()).
			Msg("HTTP request")
	}
}
