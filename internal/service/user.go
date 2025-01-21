package service

import (
	v1 "github.com/Havens-blog/e-cas-service/api/http/v1"
	"github.com/Havens-blog/e-cas-service/internal/biz"
	"github.com/Havens-blog/e-cas-service/internal/ecode"
	"github.com/Havens-blog/e-cas-service/pkg/utils/request"
	"github.com/Havens-blog/e-cas-service/pkg/utils/response"
	"github.com/ecodeclub/ginx/gctx"
	"github.com/ecodeclub/ginx/session"
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/log"
)

type HttpService struct {
	log *log.Helper
	uc  *biz.HttpUsecase
}

func NewHttpService(uc *biz.HttpUsecase, lg log.Logger) *HttpService {
	return &HttpService{
		uc:  uc,
		log: log.NewHelper(lg),
	}
}

func (s *HttpService) Login(c *gin.Context) {
	req := &v1.LoginRequest{}
	err := request.ShouldBindJSON(c, req)
	if err != nil {
		response.Fail(c, ecode.ParamsFailed, err)
		return
	}

	if !store.Verify(req.CaptchaID, req.Captcha, true) {
		response.Fail(c, ecode.CaptchaVerifyFailed, nil)
		return
	}

	data, err := s.uc.Login(c, req)
	if err != nil {
		response.Fail(c, ecode.LoginFailed, err)
		return
	}
	jwtDataAny := map[string]string{
		"roleID":   data.RoleID,
		"username": data.UserName,
		"uID":      data.GetUID(),
	}

	if err != nil {
		response.Fail(c, ecode.LoginFailed, err)
		return
	}
	_, err = session.NewSessionBuilder(&gctx.Context{Context: c}, data.GetID()).SetJwtData(jwtDataAny).Build()

	if err != nil {
		response.Fail(c, ecode.LoginFailed, err)
		return
	}

	response.Success(c, data)
}

func (s *HttpService) GetUserInfo(c *gin.Context) {
	data, err := s.uc.GetUserInfo(c)
	if err != nil {
		response.Fail(c, ecode.GetUserInfoFailed, err)
		return
	}
	response.Success(c, data)
}

func (s *HttpService) UpdateUserInfo(c *gin.Context) {
	req := &v1.UpdateUserInfoRequest{}
	err := request.ShouldBindJSON(c, req)
	if err != nil {
		response.Fail(c, ecode.ParamsFailed, err)
		return
	}
	err = s.uc.UpdateUserInfo(c, req)
	if err != nil {
		response.Fail(c, ecode.UpdateUserInfoFailed, err)
		return
	}
	response.Success(c, nil)
}

func (s *HttpService) UpdatePassword(c *gin.Context) {
	req := &v1.UpdatePasswordRequest{}
	err := request.ShouldBindJSON(c, req)
	if err != nil {
		response.Fail(c, ecode.ParamsFailed, err)
		return
	}
	err = s.uc.UpdatePassword(c, req)
	if err != nil {
		response.Fail(c, ecode.UpdatePasswordFailed, err)
		return
	}
	response.Success(c, nil)
}
