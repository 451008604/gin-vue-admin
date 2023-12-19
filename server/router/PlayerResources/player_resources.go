package PlayerResources

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type PlayerResourcesRouter struct {
}

// InitPlayerResourcesRouter 初始化 玩家资源 路由信息
func (s *PlayerResourcesRouter) InitPlayerResourcesRouter(Router *gin.RouterGroup) {
	PResourcesRouter := Router.Group("playerResources").Use(middleware.OperationRecord())
	var PResourcesApi = v1.ApiGroupApp.PlayerResourcesApiGroup.PlayerResourcesApi
	{
		PResourcesRouter.POST("setPlayerResources", PResourcesApi.SetPlayerResources) // 设置玩家资源
	}
}
