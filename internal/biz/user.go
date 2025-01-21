package biz

import (
	"github.com/go-kratos/kratos/v2/log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"go.opentelemetry.io/otel/attribute"

	v1 "github.com/Havens-blog/e-cas-service/api/http/v1"
	"github.com/Havens-blog/e-cas-service/internal/data/model"
	pJwt "github.com/Havens-blog/e-cas-service/pkg/jwt"
	"github.com/Havens-blog/e-cas-service/pkg/tracing"
	"github.com/Havens-blog/e-cas-service/pkg/utils"
)

var repoUsecase *HttpUsecase //暂时没有想到好的办法,解决consumer 依赖 datarepo实例

type HttpUsecase struct {
	repo Repo
	log  *log.Helper
}

func NewHttpUsecase(repo Repo, lg log.Logger) *HttpUsecase {
	repoUsecase = &HttpUsecase{
		repo: repo,
		log:  log.NewHelper(lg),
	}
	return repoUsecase
}

// Login 用户登录
func (uc *HttpUsecase) Login(c *gin.Context, req *v1.LoginRequest) (*v1.LoginResponse, error) {
	ctx, span := tracing.NewSpan(c.Request.Context(), "biz-Login")
	defer span.End()

	userInfo, err := uc.repo.FirstUser(ctx, &model.SysUser{Username: req.Username})
	if err != nil {
		return nil, err
	}

	if utils.BcryptCheck(req.Password, userInfo.Password) == false {
		return nil, errors.New("密码错误")
	}
	/*token, err := pJwt.Create(pJwt.Claims{
		UID:        userInfo.UID,
		UserName:   userInfo.Username,
		Phone:      userInfo.Phone,
		RoleID:     userInfo.RoleID,
		RoleName:   userInfo.RoleName,
		State:      userInfo.State,
		BufferTime: consts.Conf.Jwt.BufferTime,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,                               // 签名生效时间
			ExpiresAt: time.Now().Unix() + int64(consts.Conf.Jwt.ExpiresTime), // 过期时间
			Issuer:    consts.Conf.Jwt.Issuer,                                 // 签名的发行者
		},
	})
	if err != nil {
		return nil, errors.New("生成token失败")
	}*/

	span.SetAttributes(attribute.Key("userList").String(cast.ToString(userInfo)))
	res := &v1.LoginResponse{
		ID:       userInfo.ID,
		UID:      userInfo.UUID,
		UserName: userInfo.Username,
		NickName: userInfo.NickName,
		Birth:    "",
		Avatar:   userInfo.Avatar,
		RoleName: userInfo.RoleIds,
		Phone:    userInfo.Phone,
		Email:    userInfo.Email,

		Token:        "",
		RefreshToken: "",
	}
	return res, nil
}

// Login 用户登录
func (uc *HttpUsecase) GetUserInfo(c *gin.Context) (*v1.GetUserInfoResponse, error) {
	claims, ok := c.Get("claims")
	if !ok {
		return nil, errors.New("token解析失败")
	}
	userInfo := claims.(*pJwt.Claims)

	u, err := uc.repo.FirstUser(c.Request.Context(), &model.SysUser{UUID: userInfo.UID})
	if err != nil {
		return nil, err
	}
	res := &v1.GetUserInfoResponse{
		ID:       u.ID,
		UserName: u.Username,
		Avatar:   u.Avatar,
		Phone:    u.Phone,
		Email:    u.Email,
	}
	return res, nil
}

// Login 用户登录
func (uc *HttpUsecase) UpdateUserInfo(c *gin.Context, req *v1.UpdateUserInfoRequest) error {
	t, _ := time.Parse(time.DateOnly, req.Birth)
	return uc.repo.UpdateUserInfo(c.Request.Context(), &model.User{
		UID:      req.UID,
		Nickname: req.NickName,
		Birth:    t,
		Phone:    req.Phone,
		Wechat:   req.Wechat,
		Email:    req.Email,
		Motto:    req.Motto,
	})
}

// Login 用户登录
func (uc *HttpUsecase) UpdatePassword(c *gin.Context, req *v1.UpdatePasswordRequest) error {
	u, err := uc.repo.FirstUser(c.Request.Context(), &model.SysUser{UUID: req.UID})
	if err != nil {
		return err
	}
	if u.Password != utils.Md5([]byte(req.OldPassword+u.Salt)) {
		return errors.New("原密码错误")
	}
	salt := utils.RandString(8)
	return uc.repo.UpdatePassword(c.Request.Context(), &model.User{
		UID:      req.UID,
		Password: utils.Md5([]byte(req.NewPassword + salt)),
		Salt:     salt,
	})
}
