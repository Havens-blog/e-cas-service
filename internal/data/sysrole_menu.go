package data

import (
	pb "github.com/Havens-blog/e-cas-service/api/http/cas/v1"
	"github.com/Havens-blog/e-cas-service/internal/biz"
	"github.com/Havens-blog/e-cas-service/internal/data/model"
	"github.com/Havens-blog/e-cas-service/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/log"
)

type sysRoleMenuRepo struct {
	data *Data
	log  *log.Helper
}

func NewSysRoleMenuRepo(data *Data, logger log.Logger) biz.SysRoleMenuRepo {
	return &sysRoleMenuRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (s *sysRoleMenuRepo) FindMenuByRoleId(ctx *gin.Context, id int64) ([]*model.SysMenu, error) {
	query := s.data.Query(ctx)
	roleMenu := query.SysRoleMenu
	menu := query.SysMenu

	return menu.WithContext(ctx).
		LeftJoin(roleMenu, menu.ID.EqCol(roleMenu.MenuID)).
		Where(roleMenu.RoleID.Eq(id)).
		Where(menu.MenuType.In("C", "M")).
		Order(menu.Sort).Find()
}

func (s *sysRoleMenuRepo) GetPermission(ctx *gin.Context, id int64) ([]string, error) {
	query := s.data.Query(ctx)
	roleMenu := query.SysRoleMenu
	menu := query.SysMenu
	var result []string
	err := menu.WithContext(ctx).
		Select(menu.Permission).
		LeftJoin(roleMenu, menu.ID.EqCol(roleMenu.MenuID)).
		Where(roleMenu.RoleID.Eq(id)).
		Where(menu.MenuType.In("C", "F")).Scan(&result)
	return result, err
}

func (s *sysRoleMenuRepo) SelectMenuRole(c *gin.Context, roleKey string) ([]*pb.MenuTree, error) {
	redData := make([]*pb.MenuTree, 0)
	menuList := s.GetMenuByRoleKey(c, roleKey)
	for i := 0; i < len(menuList); i++ {
		if menuList[i].ParentID != 0 {
			continue
		}
		menuTree := &pb.MenuTree{}
		_ = utils.CopyStructFields(menuTree, menuList[i])
		menuTree.MenuId = menuList[i].ID
		menusInfo := DiguiMenu(menuList, menuTree)
		redData = append(redData, menusInfo)
	}
	return redData, nil
}

func (s *sysRoleMenuRepo) GetMenuByRoleKey(c *gin.Context, roleKey string) []*model.SysMenu {
	query := s.data.Query(c)
	menu := query.SysMenu
	rolemenu := query.SysRoleMenu
	db := menu.WithContext(c).Select().
		LeftJoin(rolemenu, rolemenu.MenuID.EqCol(menu.ID)).
		Where(rolemenu.RoleName.Eq(roleKey)).
		Where(menu.MenuType.In("C", "M")).
		Where(menu.Status.Eq(1)).Or(menu.Status.Eq(0)).Order(menu.Sort)
	menus, err := db.Debug().Find()
	if err != nil {
		return nil
	}
	return menus
}
