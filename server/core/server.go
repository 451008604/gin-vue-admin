package core

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/initialize"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	// 从db加载jwt数据
	if global.GVA_CONFIG.System.UseMultipoint || global.GVA_CONFIG.System.UseRedis {
		initialize.Redis()
	}
	// 初始化redis服务
	if global.GVA_DB != nil {
		system.LoadAll()
	}

	Router := initialize.Routers()
	Router.Static("/form-generator", "./resource/page")
	s := initServer(fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr), Router)
	global.GVA_LOG.Error(s.ListenAndServe().Error())
}
