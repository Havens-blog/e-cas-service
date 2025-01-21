package biz

import (
	"github.com/ecodeclub/ginx/gctx"
	"github.com/ecodeclub/ginx/session"
	"github.com/gin-gonic/gin"

	v1 "github.com/Havens-blog/e-cas-service/api/http/v1"
	"github.com/Havens-blog/e-cas-service/internal/data/model"
)

// GetCasbinList 获取权限列表
func (uc *HttpUsecase) GetSetting(c *gin.Context) (*v1.SettingResponse, error) {
	sess, err := session.Get(&gctx.Context{Context: c})
	uid := sess.Claims().Data["uID"]
	setting, err := uc.repo.GetSetting(c.Request.Context(), &model.Setting{UID: uid})
	if err != nil {
		return nil, err
	}
	res := &v1.SettingResponse{
		ID:                    setting.ID,
		UID:                   setting.UID,
		Lang:                  setting.Lang,
		SideModeColor:         setting.SideModeColor,
		Collapse:              setting.Collapse,
		Breadcrumb:            setting.Breadcrumb,
		DefaultRouter:         setting.DefaultRouter,
		ActiveTextColor:       setting.ActiveTextColor,
		ActiveBackgroundColor: setting.ActiveBackgroundColor,
	}
	return res, nil
}

// SetSetting 设置layout配置
func (uc *HttpUsecase) UpdateSetting(c *gin.Context, s *v1.SettingRequest) error {
	sess, err := session.Get(&gctx.Context{Context: c})
	if err != nil {
		return err
	}
	username := sess.Claims().Data["username"]
	uid := sess.Claims().Data["uID"]
	return uc.repo.UpdateSetting(c.Request.Context(), &model.Setting{
		UID:                   uid,
		Lang:                  s.Lang,
		SideModeColor:         s.SideModeColor,
		Collapse:              s.Collapse,
		Breadcrumb:            s.Breadcrumb,
		DefaultRouter:         s.DefaultRouter,
		ActiveTextColor:       s.ActiveTextColor,
		ActiveBackgroundColor: s.ActiveBackgroundColor,
		UpdateUser:            username,
	})
}
