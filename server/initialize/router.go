package initialize

import (
	"net/http"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/flipped-aurora/gin-vue-admin/server/router"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {

	if global.GVA_CONFIG.System.Env == "public" {
		gin.SetMode(gin.ReleaseMode)
	}

	Router := gin.New()

	if global.GVA_CONFIG.System.Env != "public" {
		Router.Use(gin.Logger(), gin.Recovery())
	}

	InstallPlugin(Router)
	systemRouter := router.RouterGroupApp.System
	exampleRouter := router.RouterGroupApp.Example

	if global.GVA_CONFIG.System.Env != "public" {
		Router.Static("/favicon.ico", "./dist/favicon.ico")
		Router.Static("/assets", "./dist/assets")
		Router.StaticFile("/", "./dist/index.html")
	}

	Router.StaticFS(global.GVA_CONFIG.Local.StorePath, http.Dir(global.GVA_CONFIG.Local.StorePath))

	PublicGroup := Router.Group(global.GVA_CONFIG.System.RouterPrefix)
	{
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})
	}
	{
		systemRouter.InitBaseRouter(PublicGroup)
		systemRouter.InitInitRouter(PublicGroup)
	}
	PrivateGroup := Router.Group(global.GVA_CONFIG.System.RouterPrefix)
	PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		systemRouter.InitApiRouter(PrivateGroup, PublicGroup)
		systemRouter.InitJwtRouter(PrivateGroup)
		systemRouter.InitUserRouter(PrivateGroup)
		systemRouter.InitMenuRouter(PrivateGroup)
		systemRouter.InitSystemRouter(PrivateGroup)
		systemRouter.InitCasbinRouter(PrivateGroup)
		systemRouter.InitAutoCodeRouter(PrivateGroup)
		systemRouter.InitAuthorityRouter(PrivateGroup)
		systemRouter.InitSysDictionaryRouter(PrivateGroup)
		systemRouter.InitAutoCodeHistoryRouter(PrivateGroup)
		systemRouter.InitSysOperationRecordRouter(PrivateGroup)
		systemRouter.InitSysDictionaryDetailRouter(PrivateGroup)
		systemRouter.InitAuthorityBtnRouterRouter(PrivateGroup)
		systemRouter.InitChatGptRouter(PrivateGroup)

		exampleRouter.InitCustomerRouter(PrivateGroup)
		exampleRouter.InitFileUploadAndDownloadRouter(PrivateGroup)
		router.RouterGroupApp.PlayerResources.InitPlayerResourcesRouter(PrivateGroup)
		router.RouterGroupApp.ServerAction.InitServerActionRouter(PrivateGroup)
	}

	return Router
}
