package service

import (
	"context"
	pb "github.com/Havens-blog/e-cas-service/api/http/cas/v1"
	"github.com/Havens-blog/e-cas-service/internal/biz"
	"github.com/Havens-blog/e-cas-service/internal/ecode"
	"github.com/Havens-blog/e-cas-service/pkg/utils"
	"github.com/Havens-blog/e-cas-service/pkg/utils/request"
	"github.com/Havens-blog/e-cas-service/pkg/utils/response"
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/log"
)

type RolesService struct {
	log *log.Helper
	rc  *biz.SysRoleUseCase
}

func NewRolesService(logger log.Logger, rc *biz.SysRoleUseCase) *RolesService {
	return &RolesService{
		log: log.NewHelper(logger),
		rc:  rc,
	}
}

func (r *RolesService) CreateRoles(ctx context.Context, req *pb.CreateRolesRequest) (*pb.CreateRolesReply, error) {
	return nil, nil
}
func (r *RolesService) UpdateRoles(ctx context.Context, req *pb.UpdateRolesRequest) (*pb.UpdateRolesReply, error) {
	return nil, nil
}
func (r *RolesService) ListRoles(ctx *gin.Context) {
	req := &pb.ListRolesRequest{}
	err := request.ShouldBindJSON(ctx, req)
	if err != nil {
		response.Fail(ctx, ecode.ParamsFailed, err)
		return
	}
	rolelist, total, err := r.rc.ListPage(ctx, req.RoleName, req.RoleKey, req.Status, req.PageNum, req.PageSize)
	if err != nil {
		response.Fail(ctx, ecode.Failed, err)
		return
	}
	data := make([]*pb.RoleData, len(rolelist))
	for i, d := range rolelist {
		data[i] = &pb.RoleData{
			RoleId:    d.ID,
			RoleName:  d.RoleName,
			Status:    d.Status,
			RoleKey:   d.RoleKey,
			RoleSort:  d.RoleSort,
			DataScope: int64(d.DataScope),
			CreateBy:  d.CreateBy,
			UpdateBy:  d.UpdateBy,
			Remark:    d.Remark,
			MenuIds:   nil,
		}
	}
	response.Success(ctx, &pb.ListRolesReply{
		Total:    total,
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
		Data:     data,
	})
}
func (r *RolesService) ChangeRoleStatus(ctx context.Context, req *pb.ChangeRoleStatusRequest) (*pb.ChangeRoleStatusReply, error) {
	return nil, nil
}

func (r *RolesService) GetRoles(ctx *gin.Context) {
	req := &pb.GetRolesRequest{}
	err := request.ShouldBindJSON(ctx, req)
	if err != nil {
		response.Fail(ctx, ecode.ParamsFailed, err)
		return
	}
	role, err := r.rc.GetRole(ctx, req.Id)
	if err != nil {
		response.Fail(ctx, ecode.Failed, err)
		return
	}
	response.Success(ctx, &pb.GetRolesReply{
		Role: &pb.RoleData{
			RoleId:     role.ID,
			RoleName:   role.RoleName,
			Status:     role.Status,
			RoleKey:    role.RoleKey,
			RoleSort:   role.RoleSort,
			DataScope:  int64(role.DataScope),
			MenuIds:    nil,
			CreateBy:   role.CreateBy,
			UpdateBy:   role.UpdateBy,
			Remark:     role.Remark,
			CreateTime: utils.NewTimestamp(role.CreatedAt),
			UpdateTime: utils.NewTimestamp(role.UpdatedAt),
		},
	})
}
func (r *RolesService) DataScope(ctx context.Context, req *pb.DataScopeRequest) (*pb.DataScopeReply, error) {
	return nil, nil
}
func (r *RolesService) DeleteRoles(ctx context.Context, req *pb.DeleteRolesRequest) (*pb.RoleData, error) {
	return nil, nil
}
