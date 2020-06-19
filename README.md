## Ginmid

> Request ID middleware for Gin Framework

```go
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ysicing/ginmid"
	"net/http"
	"time"
)

func main() {

	r := gin.New()

	r.Use(mid.RequestID())

	// Example ping request.
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong "+fmt.Sprint(time.Now().Unix()))
	})

	// Example / request.
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "id:"+mid.GetRequestID(c))
	})

	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
```