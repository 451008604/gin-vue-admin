package system

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type SysRouter struct{}

func (s *SysRouter) InitSystemRouter(Router *gin.RouterGroup) {
	sysRouter := Router.Group("system").Use(middleware.OperationRecord())
	sysRouterWithoutRecord := Router.Group("system")
	systemApi := v1.ApiGroupApp.SystemApiGroup.SystemApi
	{
		sysRouter.POST("setSystemConfig", systemApi.SetSystemConfig) // 设置配置文件内容
		sysRouter.POST("reloadSystem", systemApi.ReloadSystem)       // 重启服务
	}
	sysRouterWithoutRecord.POST("getSystemConfig", systemApi.GetSystemConfig) // 获取配置文件内容
	sysRouterWithoutRecord.POST("getServerInfo", systemApi.GetServerInfo)     // 获取服务器信息
}
