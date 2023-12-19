package system

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type OperationRecordRouter struct{}

func (s *OperationRecordRouter) InitSysOperationRecordRouter(Router *gin.RouterGroup) {
	operationRecordRouter := Router.Group("sysOperationRecord").Use(middleware.OperationRecord())
	operationRecordRouterWithoutRecord := Router.Group("sysOperationRecord")
	authorityMenuApi := v1.ApiGroupApp.SystemApiGroup.OperationRecordApi
	{
		operationRecordRouter.POST("createSysOperationRecord", authorityMenuApi.CreateSysOperationRecord)             // 新建SysOperationRecord
		operationRecordRouter.DELETE("deleteSysOperationRecord", authorityMenuApi.DeleteSysOperationRecord)           // 删除SysOperationRecord
		operationRecordRouter.DELETE("deleteSysOperationRecordByIds", authorityMenuApi.DeleteSysOperationRecordByIds) // 批量删除SysOperationRecord
	}
	operationRecordRouterWithoutRecord.GET("findSysOperationRecord", authorityMenuApi.FindSysOperationRecord)         // 根据ID获取SysOperationRecord
	operationRecordRouterWithoutRecord.GET("getSysOperationRecordList", authorityMenuApi.GetSysOperationRecordList)   // 获取SysOperationRecord列表
	operationRecordRouterWithoutRecord.GET("getUserOperationRecordList", authorityMenuApi.GetUserOperationRecordList) // 获取UserOperationRecord列表
}
