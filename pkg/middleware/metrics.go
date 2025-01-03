package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/Havens-blog/e-cas-service/configs/conf"
	"github.com/Havens-blog/e-cas-service/pkg/metric"
)

// Metrics returns a gin.HandlerFunc for exporting some Web metrics
func Metrics(g *conf.Global) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		method := c.Request.Method
		label := []string{g.Env, g.AppName, "http", path, method}

		metric.ReqCount.With(label...).Inc()

		c.Next()

		status := fmt.Sprintf("%d", c.Writer.Status())
		labels := []string{g.Env, g.AppName, "http", path, method, status} //http 是gin的埋点

		metric.RespCount.With(labels...).Inc()
		metric.RespDurationHistogram.With(labels...).Observe((time.Since(start).Seconds()))

	}
}
