// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v0.0.2

package v1

import (
	context "context"
	gin "github.com/gin-gonic/gin"
	http "net/http"
)

// This imports are custom by go-http.
import (
	restyv2 "github.com/go-resty/resty/v2"
	phttp "github.com/yusank/protoc-gen-go-http/http"
)

// This is a compile-time assertion to ensure that generated files are safe and compilable.
var _ context.Context

const _ = gin.Version

var _ http.Client
var _ phttp.CallOption

const _ = restyv2.Version

// 这里定义 handler interface
type HttpServiceHTTPHandler interface {
	DebugPerf(context.Context, *DebugPerfRequest) (*DebugPerfResponse, error)
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	Captcha(context.Context, *CaptchaRequest) (*CaptchaResponse, error)
	Setting(context.Context, *SettingRequest) (*SettingResponse, error)
}

// RegisterHttpServiceHTTPHandler define http router handle by gin.
// 注册路由 handler
func RegisterHttpServiceHTTPHandler(g *gin.RouterGroup, srv HttpServiceHTTPHandler) {
	g.POST("/v1/debug/perf", _HttpService_DebugPerf0_HTTP_Handler(srv))
	g.POST("/v1/login", _HttpService_Login0_HTTP_Handler(srv))
	g.GET("/v1/captcha", _HttpService_Captcha0_HTTP_Handler(srv))
	g.GET("/v1/get/setting", _HttpService_Setting0_HTTP_Handler(srv))
}

// 定义 handler
// 遍历之前解析到所有 rpc 方法信息

func _HttpService_DebugPerf0_HTTP_Handler(srv HttpServiceHTTPHandler) func(c *gin.Context) {
	return func(c *gin.Context) {
		var (
			err error
			in  = new(DebugPerfRequest)
			out = new(DebugPerfResponse)
			ctx = context.TODO()
		)

		if err = c.ShouldBind(in); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
			return
		}

		// execute
		out, err = srv.DebugPerf(ctx, in)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
			return
		}

		c.JSON(http.StatusOK, out)
	}
}

func _HttpService_Login0_HTTP_Handler(srv HttpServiceHTTPHandler) func(c *gin.Context) {
	return func(c *gin.Context) {
		var (
			err error
			in  = new(LoginRequest)
			out = new(LoginResponse)
			ctx = context.TODO()
		)

		if err = c.ShouldBind(in); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
			return
		}

		// execute
		out, err = srv.Login(ctx, in)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
			return
		}

		c.JSON(http.StatusOK, out)
	}
}

func _HttpService_Captcha0_HTTP_Handler(srv HttpServiceHTTPHandler) func(c *gin.Context) {
	return func(c *gin.Context) {
		var (
			err error
			in  = new(CaptchaRequest)
			out = new(CaptchaResponse)
			ctx = context.TODO()
		)

		if err = c.ShouldBind(in); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
			return
		}

		// execute
		out, err = srv.Captcha(ctx, in)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
			return
		}

		c.JSON(http.StatusOK, out)
	}
}

func _HttpService_Setting0_HTTP_Handler(srv HttpServiceHTTPHandler) func(c *gin.Context) {
	return func(c *gin.Context) {
		var (
			err error
			in  = new(SettingRequest)
			out = new(SettingResponse)
			ctx = context.TODO()
		)

		if err = c.ShouldBind(in); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
			return
		}

		// execute
		out, err = srv.Setting(ctx, in)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
			return
		}

		c.JSON(http.StatusOK, out)
	}
}

// HttpServiceHTTPClient defines call HttpServiceServer client
type HttpServiceHTTPClient interface {
	DebugPerf(ctx context.Context, req *DebugPerfRequest, opts ...phttp.CallOption) (*DebugPerfResponse, error)
	Login(ctx context.Context, req *LoginRequest, opts ...phttp.CallOption) (*LoginResponse, error)
	Captcha(ctx context.Context, req *CaptchaRequest, opts ...phttp.CallOption) (*CaptchaResponse, error)
	Setting(ctx context.Context, req *SettingRequest, opts ...phttp.CallOption) (*SettingResponse, error)
}

// HttpServiceHTTPClientImpl implement HttpServiceHTTPClient
type HttpServiceHTTPClientImpl struct {
	cli        *restyv2.Client
	clientOpts []phttp.ClientOption
}

func NewHttpServiceHTTPClient(cli *http.Client, opts ...phttp.ClientOption) (HttpServiceHTTPClient, error) {
	c := &HttpServiceHTTPClientImpl{
		clientOpts: opts,
	}

	hc := cli
	if hc == nil {
		hc = http.DefaultClient
	}

	c.cli = restyv2.NewWithClient(hc)
	for _, opt := range opts {
		if err := opt.Apply(c.cli); err != nil {
			return nil, err
		}
	}

	return c, nil
}

// DebugPerf is call [POST] /v1/debug/perf api.
func (c *HttpServiceHTTPClientImpl) DebugPerf(ctx context.Context, req *DebugPerfRequest, opts ...phttp.CallOption) (rsp *DebugPerfResponse, err error) {
	rsp = new(DebugPerfResponse)

	r := c.cli.R()
	for _, opt := range opts {
		if err = opt.Before(r); err != nil {
			return
		}
	}
	// set response data struct.
	r.SetResult(rsp)
	// do request
	restyResp, err := r.Execute("POST", "/v1/debug/perf")
	if err != nil {
		return nil, err
	}
	for _, opt := range opts {
		if err = opt.After(restyResp); err != nil {
			return
		}
	}

	return
}

// Login is call [POST] /v1/login api.
func (c *HttpServiceHTTPClientImpl) Login(ctx context.Context, req *LoginRequest, opts ...phttp.CallOption) (rsp *LoginResponse, err error) {
	rsp = new(LoginResponse)

	r := c.cli.R()
	for _, opt := range opts {
		if err = opt.Before(r); err != nil {
			return
		}
	}
	// set response data struct.
	r.SetResult(rsp)
	// do request
	restyResp, err := r.Execute("POST", "/v1/login")
	if err != nil {
		return nil, err
	}
	for _, opt := range opts {
		if err = opt.After(restyResp); err != nil {
			return
		}
	}

	return
}

// Captcha is call [GET] /v1/captcha api.
func (c *HttpServiceHTTPClientImpl) Captcha(ctx context.Context, req *CaptchaRequest, opts ...phttp.CallOption) (rsp *CaptchaResponse, err error) {
	rsp = new(CaptchaResponse)

	r := c.cli.R()
	for _, opt := range opts {
		if err = opt.Before(r); err != nil {
			return
		}
	}
	// set response data struct.
	r.SetResult(rsp)
	// do request
	restyResp, err := r.Execute("GET", "/v1/captcha")
	if err != nil {
		return nil, err
	}
	for _, opt := range opts {
		if err = opt.After(restyResp); err != nil {
			return
		}
	}

	return
}

// Setting is call [GET] /v1/get/setting api.
func (c *HttpServiceHTTPClientImpl) Setting(ctx context.Context, req *SettingRequest, opts ...phttp.CallOption) (rsp *SettingResponse, err error) {
	rsp = new(SettingResponse)

	r := c.cli.R()
	for _, opt := range opts {
		if err = opt.Before(r); err != nil {
			return
		}
	}
	// set response data struct.
	r.SetResult(rsp)
	// do request
	restyResp, err := r.Execute("GET", "/v1/get/setting")
	if err != nil {
		return nil, err
	}
	for _, opt := range opts {
		if err = opt.After(restyResp); err != nil {
			return
		}
	}

	return
}
