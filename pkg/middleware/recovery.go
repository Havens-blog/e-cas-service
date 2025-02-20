package middleware

import (
	"fmt"
	"runtime"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/attribute"

	"github.com/Havens-blog/e-cas-service/internal/consts"
	"github.com/Havens-blog/e-cas-service/internal/ecode"
	"github.com/Havens-blog/e-cas-service/pkg/email"
	"github.com/Havens-blog/e-cas-service/pkg/tracing"
	"github.com/Havens-blog/e-cas-service/pkg/utils/response"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				ctx, span := tracing.NewSpan(c.Request.Context(), "recover")
				c.Request = c.Request.WithContext(ctx)
				defer span.End()

				buf := make([]byte, 64<<10)
				buf = buf[:runtime.Stack(buf, false)]
				bufs := string(buf)
				span.SetAttributes(attribute.Key("err").String(fmt.Sprintf("%s", err)))
				span.SetAttributes(attribute.Key("painc").String(bufs))

				email.SendWarn(c.Request.Context(), consts.EmailTitlePanic, bufs)

				response.Fail(c, ecode.Failed, nil)
				c.Abort()
			}
		}()
		c.Next()
	}
}
