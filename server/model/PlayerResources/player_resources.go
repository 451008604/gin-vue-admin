// 自动生成模板PlayerResources
package PlayerResources

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 玩家资源 结构体  PlayerResources
type PlayerResources struct {
	global.GVA_MODEL
	UserId     string `json:"userId" form:"userId" gorm:"column:user_id;comment:用户ID;size:64;"`              // 用户ID
	ProperInfo string `json:"properInfo" form:"properInfo" gorm:"column:proper_info;comment:资源信息;size:256;"` // 资源信息
	CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
	UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
	DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}

// TableName 玩家资源 PlayerResources自定义表名 player_resources
func (PlayerResources) TableName() string {
	return "player_resources"
}
