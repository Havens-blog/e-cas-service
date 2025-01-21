package service

import (
	"context"
	"errors"
	"github.com/Havens-blog/e-cas-service/internal/biz"
	"github.com/Havens-blog/e-cas-service/internal/ecode"
	"github.com/Havens-blog/e-cas-service/pkg/common/constant"
	"github.com/Havens-blog/e-cas-service/pkg/utils"
	"github.com/Havens-blog/e-cas-service/pkg/utils/response"
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/log"

	pb "github.com/Havens-blog/e-cas-service/api/http/cas/v1"
)

type SysuserService struct {
	pb.UnimplementedSysuserServer
	log          *log.Helper
	userCase     *biz.SysUserUseCase
	roleCase     *biz.SysRoleUseCase
	roleMenuCase *biz.SysRoleMenuUseCase
}

func NewSysuserService(userCase *biz.SysUserUseCase, roleCase *biz.SysRoleUseCase, roleMenuCase *biz.SysRoleMenuUseCase, logger log.Logger) *SysuserService {
	return &SysuserService{
		userCase:     userCase,
		roleCase:     roleCase,
		roleMenuCase: roleMenuCase,
		log:          log.NewHelper(log.With(logger, "module", "service/sysuser")),
	}
}

func (s *SysuserService) CreateSysuser(ctx context.Context, req *pb.CreateSysuserRequest) (*pb.CreateSysuserReply, error) {
	return &pb.CreateSysuserReply{}, nil
}
func (s *SysuserService) UpdateSysuser(ctx context.Context, req *pb.UpdateSysuserRequest) (*pb.UpdateSysuserReply, error) {
	return &pb.UpdateSysuserReply{}, nil
}
func (s *SysuserService) DeleteSysuser(ctx context.Context, req *pb.DeleteSysuserRequest) (*pb.DeleteSysuserReply, error) {
	return &pb.DeleteSysuserReply{}, nil
}
func (s *SysuserService) GetSysuser(ctx context.Context, req *pb.GetSysuserRequest) (*pb.GetSysuserReply, error) {
	return &pb.GetSysuserReply{}, nil
}
func (s *SysuserService) ListSysuser(ctx context.Context, req *pb.ListSysuserRequest) (*pb.ListSysuserReply, error) {
	return &pb.ListSysuserReply{}, nil
}
func (s *SysuserService) GetCaptcha(ctx context.Context, req *pb.GetCaptchaRequest) (*pb.GetCaptchaReply, error) {
	return &pb.GetCaptchaReply{}, nil
}
func (s *SysuserService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginReply, error) {
	return &pb.LoginReply{}, nil
}
func (s *SysuserService) Logout(ctx context.Context, req *pb.LogoutRequest) (*pb.LogoutReply, error) {
	return &pb.LogoutReply{}, nil
}
func (s *SysuserService) Auth(c *gin.Context) {
	req := &pb.AuthRequest{}
	username := c.Query("username")
	if username == "" {
		response.Fail(c, ecode.ParamsFailed, errors.New("username is required"))
		return
	}
	req.Username = username
	user, err := s.userCase.FindUserByUsername(c, req.Username)
	if err != nil {
		response.Fail(c, ecode.GetUserInfoFailed, err)
		return
	}
	role, err := s.roleCase.GetRole(c, user.RoleID)
	if err != nil {
		response.Fail(c, ecode.GetUserInfoFailed, err)
		return
	}
	permits, err := s.roleMenuCase.GetPermission(c, role.ID)
	if err != nil {
		response.Fail(c, ecode.GetUserInfoFailed, err)
		return
	}
	var menus []*pb.MenuTree
	if role.Status == constant.StatusMenusForbidden {
		menus = make([]*pb.MenuTree, 0)
	} else {
		menus, err = s.roleMenuCase.SelectMenuRole(c, role.RoleName)
		if err != nil {
			response.Fail(c, ecode.GetUserInfoFailed, err)
			return
		}
	}

	pbUser := &pb.AuthReply_User{
		UserId:    user.ID,
		NickName:  user.NickName,
		Phone:     user.Phone,
		RoleId:    user.RoleID,
		Avatar:    user.Avatar,
		Sex:       user.Sex,
		Email:     user.Email,
		DeptId:    user.DeptID,
		PostId:    user.PostID,
		RoleIds:   user.RoleIds,
		PostIds:   user.PostIds,
		CreateBy:  user.CreateBy,
		UpdateBy:  user.UpdateBy,
		Remark:    user.Remark,
		Status:    user.Status,
		CreatedAt: utils.NewTimestamp(user.CreatedAt),
		UpdatedAt: utils.NewTimestamp(user.UpdatedAt),
		Username:  user.Username,
		RoleName:  role.RoleName,
	}

	pbRole := &pb.AuthReply_Role{
		RoleId:    role.ID,
		RoleName:  role.RoleName,
		Status:    role.Status,
		RoleKey:   role.RoleKey,
		RoleSort:  role.RoleSort,
		DataScope: role.DataScope,
		CreateBy:  role.CreateBy,
		UpdateBy:  role.UpdateBy,
		Remark:    role.Remark,
		ApiIds:    nil,
		MenuIds:   nil,
		DeptIds:   nil,
		CreatedAt: utils.NewTimestamp(user.CreatedAt),
		UpdatedAt: utils.NewTimestamp(user.UpdatedAt),
	}

	response.Success(c, &pb.AuthReply{
		User:        pbUser,
		Role:        pbRole,
		Permissions: permits,
		Menus:       Build(menus),
	})
}
func (s *SysuserService) ChangeStatus(ctx context.Context, req *pb.ChangeStatusRequest) (*pb.ChangeStatusReply, error) {
	return &pb.ChangeStatusReply{}, nil
}
func (s *SysuserService) UpdatePassword(ctx context.Context, req *pb.UpdatePasswordRequest) (*pb.UpdatePasswordReply, error) {
	return &pb.UpdatePasswordReply{}, nil
}
func (s *SysuserService) GetPostInit(ctx context.Context, req *pb.GetPostInitRequest) (*pb.GetPostInitReply, error) {
	return &pb.GetPostInitReply{}, nil
}
func (s *SysuserService) GetUserRolePost(ctx context.Context, req *pb.GetUserRolePostRequest) (*pb.GetUserRolePostReply, error) {
	return &pb.GetUserRolePostReply{}, nil
}
func (s *SysuserService) GetUserGoogleSecret(ctx context.Context, req *pb.GetUserGoogleSecretRequest) (*pb.GetUserGoogleSecretReply, error) {
	return &pb.GetUserGoogleSecretReply{}, nil
}
