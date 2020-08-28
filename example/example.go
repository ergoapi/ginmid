// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/ysicing/ginmid"
	"github.com/ysicing/logger"
	"net/http"
	"time"
)

func init() {
	cfg := logger.LogConfig{Simple: false}
	logger.InitLogger(&cfg)
}

func main() {

	r := gin.New()

	r.Use(mid.RequestID(), mid.PromMiddleware(nil), mid.Log())

	// Example ping request.
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong "+fmt.Sprint(time.Now().Unix()))
	})

	// Example / request.
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "id:"+mid.GetRequestID(c))
	})

	// Example /metrics
	r.GET("/metrics", mid.PromHandler(promhttp.Handler()))

	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
