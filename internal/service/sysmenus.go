package service

import (
	"context"
	"errors"
	pb "github.com/Havens-blog/e-cas-service/api/http/cas/v1"
	"github.com/Havens-blog/e-cas-service/internal/biz"
	"github.com/Havens-blog/e-cas-service/internal/data/model"
	"github.com/Havens-blog/e-cas-service/internal/ecode"
	"github.com/Havens-blog/e-cas-service/pkg/utils"
	"github.com/Havens-blog/e-cas-service/pkg/utils/request"
	"github.com/Havens-blog/e-cas-service/pkg/utils/response"
	"github.com/ecodeclub/ginx/gctx"
	"github.com/ecodeclub/ginx/session"
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/log"
	"strconv"
	"strings"
)

type MenusService struct {
	menuUseCase     *biz.SysMenuUseCase
	roleMenuUseCase *biz.SysRoleMenuUseCase
	log             *log.Helper
}

func NewMenusService(menusUseCase *biz.SysMenuUseCase, roleMenuUseCase *biz.SysRoleMenuUseCase, logger log.Logger) *MenusService {
	return &MenusService{
		menuUseCase:     menusUseCase,
		roleMenuUseCase: roleMenuUseCase,
		log:             log.NewHelper(log.With(logger, "module", "service/menus")),
	}
}

func (s *MenusService) CreateMenus(ctx *gin.Context) {
	req := &pb.CreateMenusRequest{}
	err := request.ShouldBindJSON(ctx, req)
	if err != nil {
		response.Fail(ctx, ecode.ParamsFailed, err)
		return
	}
	err = s.menuUseCase.AddRoleMenu(ctx, &model.SysMenu{
		MenuName:   req.MenuName,
		Title:      req.Title,
		ParentID:   int64(req.ParentId),
		Sort:       req.Sort,
		Icon:       req.Icon,
		Path:       req.Path,
		Component:  req.Component,
		IsIframe:   req.IsIframe,
		Link:       req.IsLink,
		MenuType:   req.MenuType,
		Hidden:     int32(req.IsHide),
		KeepAlive:  req.IsKeepAlive,
		IsAffix:    req.IsAffix,
		Permission: req.Permission,
		Status:     req.Status,
		Remark:     req.Remark,
	})
	if err != nil {
		response.Fail(ctx, ecode.MenuListFailed, err)
		return
	}
	sess, err := session.Get(&gctx.Context{Context: ctx})
	if err != nil {
		response.Fail(ctx, ecode.Failed, err)
		return
	}
	// 获取用户ID
	roleIDStr := sess.Claims().Data["roleID"]
	roleID, err := strconv.ParseInt(roleIDStr, 10, 64)
	if err != nil {
		response.Fail(ctx, ecode.Failed, errors.New("invalid roleID"))
		return
	}
	permissions, err := s.roleMenuUseCase.GetPermission(ctx, roleID)
	if err != nil {
		response.Fail(ctx, ecode.Failed, err)
	}
	menus, err := s.roleMenuUseCase.FindMenuByRoleId(ctx, roleID)
	if err != nil {
		response.Fail(ctx, ecode.Failed, err)
	}
	response.Success(ctx, &pb.CreateMenusReply{
		Menus:       convertToMenuTree(menus),
		Permissions: permissions,
	})
}

func convertToMenuTree(sysMenus []*model.SysMenu) []*pb.MenuTree {
	menuMap := make(map[int64]*pb.MenuTree)

	// 创建 MenuTree 节点并将其存储在 menuMap 中
	for _, menu := range sysMenus {
		menuTree := &pb.MenuTree{
			MenuId:      menu.ID,
			MenuName:    menu.MenuName,
			Title:       menu.Title,
			ParentId:    menu.ParentID,
			Sort:        menu.Sort,
			Icon:        menu.Icon,
			Path:        menu.Path,
			Component:   menu.Component,
			IsIframe:    menu.IsIframe,
			IsLink:      menu.Link,
			MenuType:    menu.MenuType,
			IsHide:      menu.Hidden,
			IsKeepAlive: menu.KeepAlive,
			IsAffix:     menu.IsAffix,
			Permission:  menu.Permission,
			Status:      menu.Status,
			CreateBy:    menu.CreateBy,
			UpdateBy:    menu.UpdateBy,
			Remark:      menu.Remark,
			CreateTime:  utils.NewTimestamp(menu.CreatedAt),
			UpdateTime:  utils.NewTimestamp(menu.UpdatedAt),
			Children:    []*pb.MenuTree{},
		}
		menuMap[menu.ID] = menuTree
	}
	// 构建菜单树
	var rootMenuTrees []*pb.MenuTree
	for _, menu := range sysMenus {
		parentID := menu.ParentID
		if parentID != 0 {
			parentMenu, ok := menuMap[parentID]
			if ok {
				parentMenu.Children = append(parentMenu.Children, menuMap[menu.ID])
			}
		} else {
			rootMenuTrees = append(rootMenuTrees, menuMap[menu.ID])
		}
	}
	return rootMenuTrees
}
func (s *MenusService) ListMenus(ctx context.Context, req *pb.ListMenusRequest) (*pb.ListMenusReply, error) {
	return &pb.ListMenusReply{}, nil
}
func (s *MenusService) GetMenusTree(ctx context.Context, req *pb.GetMenusTreeRequest) (*pb.GetMenusTreeReply, error) {
	return &pb.GetMenusTreeReply{}, nil
}
func (s *MenusService) UpdateMenus(ctx context.Context, req *pb.UpdateMenusRequest) (*pb.UpdateMenusReply, error) {
	return &pb.UpdateMenusReply{}, nil
}
func (s *MenusService) DeleteMenus(ctx *gin.Context) {
	req := &pb.DeleteMenusRequest{}
	idStr := ctx.Query("id")
	if idStr == "" {
		response.Fail(ctx, ecode.ParamsFailed, errors.New("id is required"))
		return
	}

	// 将 id 字符串转换为 int64
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.Fail(ctx, ecode.ParamsFailed, err)
		return
	}
	req.Id = id
	err = s.menuUseCase.DeleteRoleMenuByID(ctx, req.Id)
	if err != nil {
		response.Fail(ctx, ecode.Failed, err)
		return
	}
	response.Success(ctx, nil)
}
func (s *MenusService) GetMenus(ctx *gin.Context) {
	req := &pb.GetMenusRequest{}
	idStr := ctx.Query("id")
	if idStr == "" {
		response.Fail(ctx, ecode.ParamsFailed, errors.New("id is required"))
		return
	}

	// 将 id 字符串转换为 int64
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.Fail(ctx, ecode.ParamsFailed, err)
		return
	}
	req.Id = id
	menu, err := s.menuUseCase.GetMenu(ctx, req.Id)
	if err != nil {
		response.Fail(ctx, ecode.Failed, err)
		return
	}
	response.Success(ctx, &pb.GetMenusReply{
		MenuId:      int32(menu.ID),
		MenuName:    menu.MenuName,
		Title:       menu.Title,
		ParentId:    int32(menu.ParentID),
		Sort:        menu.Sort,
		Icon:        menu.Icon,
		Path:        menu.Path,
		Component:   menu.Component,
		IsIframe:    menu.IsIframe,
		IsLink:      menu.Link,
		MenuType:    menu.MenuType,
		IsHide:      int64(menu.Hidden),
		IsKeepAlive: menu.KeepAlive,
		IsAffix:     menu.IsAffix,
		Permission:  menu.Permission,
		Status:      menu.Status,
		CreateBy:    menu.CreateBy,
		UpdateBy:    menu.UpdateBy,
		Remark:      menu.Remark,
		CreatedAt:   utils.NewTimestamp(menu.CreatedAt),
		UpdatedAt:   utils.NewTimestamp(menu.UpdatedAt),
	})
}
func (s *MenusService) RoleMenuTreeSelect(ctx context.Context, req *pb.RoleMenuTreeSelectRequest) (*pb.RoleMenuTreeSelectReply, error) {
	return &pb.RoleMenuTreeSelectReply{}, nil
}

// Build 构建前端路由
func Build(menus []*pb.MenuTree) []*pb.MenuTreeAuth {
	rvs := make([]*pb.MenuTreeAuth, 0)
	for _, ms := range menus {
		rv := &pb.MenuTreeAuth{
			Name:      ms.MenuName,
			Path:      ms.Path,
			Component: ms.Component,
			Meta: &pb.MenuTreeMeta{
				Title:       ms.MenuName,
				IsLink:      len(ms.IsLink) > 0,
				IsHide:      ms.IsHide == 2,
				IsKeepAlive: ms.IsKeepAlive == 1,
				IsAffix:     ms.IsAffix == 1,
				IsIframe:    ms.IsIframe == 1,
				Auth:        make([]string, 0),
				Icon:        ms.Icon,
			},
		}

		if ms.Permission != "" {
			rv.Meta.Auth = strings.Split(ms.Permission, ",")
		}
		rv.Children = Build(ms.Children)
		rvs = append(rvs, rv)
	}
	return rvs
}
