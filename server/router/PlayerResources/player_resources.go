package PlayerResources

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type PlayerResourcesRouter struct {
}

// InitPlayerResourcesRouter 初始化 玩家资源 路由信息
func (s *PlayerResourcesRouter) InitPlayerResourcesRouter(Router *gin.RouterGroup) {
	PResourcesRouter := Router.Group("PResources").Use(middleware.OperationRecord())
	PResourcesRouterWithoutRecord := Router.Group("PResources")
	var PResourcesApi = v1.ApiGroupApp.PlayerResourcesApiGroup.PlayerResourcesApi
	{
		PResourcesRouter.POST("createPlayerResources", PResourcesApi.CreatePlayerResources)             // 新建玩家资源
		PResourcesRouter.DELETE("deletePlayerResources", PResourcesApi.DeletePlayerResources)           // 删除玩家资源
		PResourcesRouter.DELETE("deletePlayerResourcesByIds", PResourcesApi.DeletePlayerResourcesByIds) // 批量删除玩家资源
		PResourcesRouter.PUT("updatePlayerResources", PResourcesApi.UpdatePlayerResources)              // 更新玩家资源

		PResourcesRouter.POST("setPlayerResources", PResourcesApi.SetPlayerResources) // 设置玩家资源
	}
	{
		PResourcesRouterWithoutRecord.GET("findPlayerResources", PResourcesApi.FindPlayerResources)       // 根据ID获取玩家资源
		PResourcesRouterWithoutRecord.GET("getPlayerResourcesList", PResourcesApi.GetPlayerResourcesList) // 获取玩家资源列表
	}
}
