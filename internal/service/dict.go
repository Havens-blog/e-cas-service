package service

import (
	"github.com/gin-gonic/gin"

	"github.com/Havens-blog/e-cas-service/internal/consts"
	"github.com/Havens-blog/e-cas-service/internal/ecode"
	"github.com/Havens-blog/e-cas-service/pkg/consul"
	"github.com/Havens-blog/e-cas-service/pkg/utils/response"
)

func (s *HttpService) GetDictList(c *gin.Context) {
	cli, err := consul.NewConsulClient(consts.Conf.Consul.Host, consts.Conf.Consul.Token)
	if err != nil {
		response.Fail(c, ecode.GetDictListFailed, err)
		return
	}

	dict := make(map[string]interface{}, 0)
	_, err = consul.GetConsulKV(cli, consts.Conf.Consul.Kv.DictPath, &dict)
	if err != nil {
		response.Fail(c, ecode.GetDictListFailed, err)
		return
	}
	res := map[string]interface{}{
		"filename": "dict.json",
		"dictInfo": dict,
	}
	response.Success(c, res)
}
