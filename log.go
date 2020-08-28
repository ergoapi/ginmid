// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package mid

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ysicing/logger"
	"time"
)

const (
	errLogFormat = "requestid %v => %v | %v | %v | %v | %v <= err: %v"
	logFormat    = "requestid %v => %v | %v | %v | %v | %v "
)

// Log log
func Log() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		start := time.Now()
		c.Next()
		end := time.Now()
		latency := end.Sub(start)
		if len(c.Errors) > 0 || c.Writer.Status() >= 400 {
			logger.Logger.Error(fmt.Sprintf(errLogFormat, GetRequestID(c), c.Writer.Status(), c.ClientIP(), c.Request.Method, path, latency, c.Errors.String()))
		} else {
			logger.Logger.Debug(fmt.Sprintf(logFormat, GetRequestID(c), c.Writer.Status(), c.ClientIP(), c.Request.Method, path, latency))
		}
	}
}