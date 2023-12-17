package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

type SysOperationRecordSearch struct {
	system.SysOperationRecord
	request.PageInfo
	UserName string `json:"userName" form:"userName"`
}

type UserOperationRecordSearch struct {
	system.SysOperationRecord
	request.PageInfo
}
