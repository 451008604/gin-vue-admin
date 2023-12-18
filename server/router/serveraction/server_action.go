package serveraction

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ServerActionRouter struct {
}

// InitServerActionRouter 初始化 游戏服务指令 路由信息
func (s *ServerActionRouter) InitServerActionRouter(Router *gin.RouterGroup) {
	serverActionRouter := Router.Group("serverAction").Use(middleware.OperationRecord())
	var serverActionApi = v1.ApiGroupApp.ServerActionApiGroup.ServerActionApi
	{
		serverActionRouter.PUT("updateServerConfig", serverActionApi.UpdateServerConfig) // 服务配置更新
		serverActionRouter.PUT("updateFixedConfig", serverActionApi.UpdateFixedConfig)   // 固定配置更新
	}
}
