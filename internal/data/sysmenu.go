package data

import (
	pb "github.com/Havens-blog/e-cas-service/api/http/cas/v1"
	"github.com/Havens-blog/e-cas-service/internal/biz"
	"github.com/Havens-blog/e-cas-service/internal/data/model"
	"github.com/Havens-blog/e-cas-service/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/log"
)

type sysMenuRepo struct {
	data *Data
	log  *log.Helper
}

func NewSysMenuRepo(data *Data, logger log.Logger) biz.SysMenuRepo {
	return &sysMenuRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/sysmenu")),
	}
}

func (s *sysMenuRepo) FindById(ctx *gin.Context, id int64) (*model.SysMenu, error) {
	q := s.data.Query(ctx).SysMenu
	return q.WithContext(ctx).Where(q.ID.Eq(id)).First()
}

func (s *sysMenuRepo) DeleteRoleMenuByID(ctx *gin.Context, id int64) error {
	q := s.data.Query(ctx).SysMenu
	_, err := q.WithContext(ctx).Where(q.ID.Eq(id)).Delete()
	return err
}

func (s *sysMenuRepo) AddRoleMenu(ctx *gin.Context, req *model.SysMenu) error {
	q := s.data.Query(ctx).SysMenu
	return q.WithContext(ctx).Create(req)
}

func DiguiMenu(menulist []*model.SysMenu, menu *pb.MenuTree) *pb.MenuTree {
	list := menulist

	min := make([]*pb.MenuTree, 0)
	for j := 0; j < len(list); j++ {

		if menu.MenuId != list[j].ParentID {
			continue
		}
		mi := &pb.MenuTree{}
		mi.MenuId = list[j].ID
		mi.MenuName = list[j].MenuName
		mi.Title = list[j].Title
		mi.Icon = list[j].Icon
		mi.Path = list[j].Path
		mi.MenuType = list[j].MenuType
		mi.IsKeepAlive = list[j].KeepAlive
		mi.Permission = list[j].Permission
		mi.ParentId = list[j].ParentID
		mi.IsAffix = list[j].IsAffix
		mi.IsIframe = list[j].IsIframe
		mi.IsLink = list[j].Link
		mi.Component = list[j].Component
		mi.Sort = list[j].Sort
		mi.Status = list[j].Status
		mi.IsHide = list[j].Hidden
		mi.CreateTime = utils.NewTimestamp(list[j].CreatedAt)
		mi.UpdateTime = utils.NewTimestamp(list[j].UpdatedAt)
		mi.Children = []*pb.MenuTree{}

		if mi.MenuType != "F" {
			ms := DiguiMenu(menulist, mi)
			min = append(min, ms)
		} else {
			min = append(min, mi)
		}
	}
	menu.Children = min
	return menu
}
