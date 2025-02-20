package server

import (
	"github.com/ecodeclub/ginx/session"
	ginRedis "github.com/ecodeclub/ginx/session/redis"
	hhttp "net/http"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"

	"github.com/Havens-blog/e-cas-service/configs/conf"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/Havens-blog/e-cas-service/internal/consts"
	"github.com/Havens-blog/e-cas-service/internal/service"
	m "github.com/Havens-blog/e-cas-service/pkg/middleware"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, cs *conf.Session, s *service.HttpService,
	role *service.RolesService,
	menu *service.MenusService,
	user *service.SysuserService,
	lg log.Logger) *http.Server {
	var opts = []http.ServerOption{}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != "" {
		opts = append(opts, http.Timeout(viper.GetDuration(c.Http.Timeout)))
	}

	srv := http.NewServer(opts...)
	// session
	sp := ginRedis.NewSessionProvider(consts.Rdb, cs.SessionEncryptedKey)

	srv.HandlePrefix("/", routers(s, role, menu, user, sp, lg))

	return srv
}

func routers(s *service.HttpService, role *service.RolesService,
	menu *service.MenusService,
	user *service.SysuserService,
	sp session.Provider, lg log.Logger) *gin.Engine {
	session.SetDefaultProvider(sp)
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(m.Recovery(), m.Cors(consts.Conf.Cors), m.Tracing(consts.Conf.Global), m.Metrics(consts.Conf.Global))
	pprof.Register(router)
	router.StaticFS("/docs", hhttp.Dir("./docs/openapi"))
	r := router.Group("")
	{
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		r.GET("/metrics", gin.WrapH(promhttp.Handler()))
	}
	r1 := router.Group("/api/cas/")
	{
		r1.POST("/v1/login", s.Login)    //登录
		r1.GET("/v1/captcha", s.Captcha) //获取验证码
	}

	// 角色管理
	{
		r1.POST("/v1/system/role/", role.GetRoles)
		r1.POST("/v1/system/role/list", role.ListRoles)
	}
	// 权限管理
	{
		r1.POST("/v1/system/menu/", menu.CreateMenus)
		r1.GET("/v1/system/menu/", menu.GetMenus)
		r1.DELETE("/v1/system/menu/", menu.DeleteMenus)

		r1.GET("/v1/system/user/auth", user.Auth)
	}
	router.Use(session.CheckLoginMiddleware())
	r2 := router.Group("/api/cas/").Use(m.CasbinV1(sp, lg), m.Limiter(), m.OperationV1(sp))
	//r2 := router.Group("").Use(m.Jwt(), m.Casbin(), m.Limiter(), m.Operation())
	{
		r2.POST("/v1/debug/perf", s.DebugPerf)         //性能测试
		r2.GET("/v1/get/setting", s.GetSetting)        //获取模板设置
		r2.POST("/v1/update/setting", s.UpdateSetting) //设置模板配置
		r2.GET("/v1/get/dictList", s.GetDictList)      //获取字典序

		r2.GET("/v1/get/allMenuList", s.GetAllMenuList)      //获取全部路由菜单
		r2.GET("/v1/get/roleMenuList", s.GetRoleMenuList)    //获取角色路由菜单
		r2.POST("/v1/add/roleMenu", s.AddRoleMenu)           //添加菜单
		r2.POST("/v1/delete/roleMenu", s.DeleteRoleMenuByID) //通过ID删除菜单
		r2.POST("/v1/update/roleMenu", s.UpdateRoleMenu)     //更新菜单

		r2.POST("/v1/get/roleCasbinList", s.GetRoleCasbinList) //获取权限列表
		r2.POST("/v1/add/roleCasbin", s.AddRoleCasbin)         //添加权限
		r2.POST("/v1/delete/roleCasbin", s.DeleteRoleCasbin)   //更新权限
		r2.POST("/v1/update/roleCasbin", s.UpdateRoleCasbin)   //删除权限

		r2.GET("/v1/get/userInfo", s.GetUserInfo)        //获取用户信息
		r2.POST("/v1/update/userInfo", s.UpdateUserInfo) //更新用户信息
		r2.POST("/v1/update/password", s.UpdatePassword) //更新用户密码
	}

	return router
}
