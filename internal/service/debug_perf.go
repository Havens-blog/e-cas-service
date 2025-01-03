package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	v1 "github.com/Havens-blog/e-cas-service/api/http/v1"
	"github.com/Havens-blog/e-cas-service/internal/ecode"
	"github.com/Havens-blog/e-cas-service/pkg/utils/request"
	"github.com/Havens-blog/e-cas-service/pkg/utils/response"
)

func (s *HttpService) DebugPerf(c *gin.Context) {
	req := &v1.DebugPerfRequest{}
	err := request.ShouldBindJSON(c, req)
	if err != nil {
		response.Fail(c, ecode.ParamsFailed, err)
		return
	}
	res, err := s.uc.DebugPerf(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	response.Success(c, res)
}
