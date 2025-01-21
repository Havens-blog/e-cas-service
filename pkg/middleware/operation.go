package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ecodeclub/ginx"
	"io"
	"net/http"
	"time"

	"github.com/Havens-blog/e-cas-service/internal/consts"
	"github.com/Havens-blog/e-cas-service/pkg/email"
	"github.com/Havens-blog/e-cas-service/pkg/jwt"
	"github.com/Havens-blog/e-cas-service/pkg/kafka"
	"github.com/Havens-blog/e-cas-service/pkg/metric"
	"github.com/Havens-blog/e-cas-service/pkg/utils"
	"github.com/ecodeclub/ginx/session"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

// Operation 记录请求流水
func Operation() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		claims, _ := c.Get("claims")
		userInfo := claims.(*jwt.Claims)
		uid := userInfo.UID
		reqBody, _ := io.ReadAll(c.Request.Body)
		c.Request.Body = io.NopCloser(bytes.NewBuffer(reqBody))
		writer := responseBodyWriter{
			ResponseWriter: c.Writer,
			body:           &bytes.Buffer{},
		}
		c.Writer = writer

		c.Next()

		var g errgroup.Group
		g.Go(func() error {
			var respbody map[string]interface{}
			var reqsBody map[string]interface{}
			_ = json.Unmarshal(writer.body.Bytes(), &respbody)
			_ = json.Unmarshal(reqBody, &reqsBody)
			record := map[string]interface{}{
				"uid":       uid,
				"date_time": time.Now().Local().Format(time.DateTime),
				"ip":        c.ClientIP(),
				"method":    c.Request.Method,
				"path":      c.Request.RequestURI,
				"agent":     c.Request.UserAgent(),
				"status":    int32(c.Writer.Status()),
				"latency":   time.Since(start).String(),
				"req_body":  reqsBody,
				"resp_body": respbody,
			}
			producer, err := kafka.NewProducer(consts.KafkaTopicOperationRecord)
			if err != nil {
				return err
			}
			data := utils.JsonToMarshal(record)
			return producer.Send(data)
		})
		if err := g.Wait(); err != nil {
			email.SendWarn(c.Request.Context(), consts.EmailTitleKafkaProducer, err.Error())
			metric.Count.With(fmt.Sprintf("producer_%s_error", consts.KafkaTopicOperationRecord)).Inc()
		}
	}
}

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

// Operation 记录请求流水
func OperationV1(sp session.Provider) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		gCtx := &ginx.Context{Context: c}
		sess, err := sp.Get(gCtx)
		if err != nil {
			gCtx.AbortWithStatus(http.StatusForbidden)
			return
		}
		uid := sess.Claims().Uid

		reqBody, _ := io.ReadAll(c.Request.Body)
		c.Request.Body = io.NopCloser(bytes.NewBuffer(reqBody))
		writer := responseBodyWriter{
			ResponseWriter: c.Writer,
			body:           &bytes.Buffer{},
		}
		c.Writer = writer

		c.Next()

		var g errgroup.Group
		g.Go(func() error {
			var respbody map[string]interface{}
			var reqsBody map[string]interface{}
			_ = json.Unmarshal(writer.body.Bytes(), &respbody)
			_ = json.Unmarshal(reqBody, &reqsBody)
			record := map[string]interface{}{
				"uid":       uid,
				"date_time": time.Now().Local().Format(time.DateTime),
				"ip":        c.ClientIP(),
				"method":    c.Request.Method,
				"path":      c.Request.RequestURI,
				"agent":     c.Request.UserAgent(),
				"status":    int32(c.Writer.Status()),
				"latency":   time.Since(start).String(),
				"req_body":  reqsBody,
				"resp_body": respbody,
			}
			producer, err := kafka.NewProducer(consts.KafkaTopicOperationRecord)
			if err != nil {
				return err
			}
			data := utils.JsonToMarshal(record)
			return producer.Send(data)
		})
		if err := g.Wait(); err != nil {
			email.SendWarn(c.Request.Context(), consts.EmailTitleKafkaProducer, err.Error())
			metric.Count.With(fmt.Sprintf("producer_%s_error", consts.KafkaTopicOperationRecord)).Inc()
		}
	}
}
