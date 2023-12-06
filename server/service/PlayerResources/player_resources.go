package PlayerResources

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/PlayerResources"
	PlayerResourcesReq "github.com/flipped-aurora/gin-vue-admin/server/model/PlayerResources/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"gorm.io/gorm"
)

type PlayerResourcesService struct {
}

// CreatePlayerResources 创建玩家资源记录
// Author [piexlmax](https://github.com/piexlmax)
func (PResourcesService *PlayerResourcesService) CreatePlayerResources(PResources *PlayerResources.PlayerResources) (err error) {
	err = global.GVA_DB.Create(PResources).Error
	return err
}

// DeletePlayerResources 删除玩家资源记录
// Author [piexlmax](https://github.com/piexlmax)
func (PResourcesService *PlayerResourcesService) DeletePlayerResources(PResources PlayerResources.PlayerResources) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&PlayerResources.PlayerResources{}).Where("id = ?", PResources.ID).Update("deleted_by", PResources.DeletedBy).Error; err != nil {
			return err
		}
		if err = tx.Delete(&PResources).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeletePlayerResourcesByIds 批量删除玩家资源记录
// Author [piexlmax](https://github.com/piexlmax)
func (PResourcesService *PlayerResourcesService) DeletePlayerResourcesByIds(ids request.IdsReq, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&PlayerResources.PlayerResources{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", ids.Ids).Delete(&PlayerResources.PlayerResources{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdatePlayerResources 更新玩家资源记录
// Author [piexlmax](https://github.com/piexlmax)
func (PResourcesService *PlayerResourcesService) UpdatePlayerResources(PResources PlayerResources.PlayerResources) (err error) {
	err = global.GVA_DB.Save(&PResources).Error
	return err
}

// GetPlayerResources 根据id获取玩家资源记录
// Author [piexlmax](https://github.com/piexlmax)
func (PResourcesService *PlayerResourcesService) GetPlayerResources(id uint) (PResources PlayerResources.PlayerResources, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&PResources).Error
	return
}

// GetPlayerResourcesInfoList 分页获取玩家资源记录
// Author [piexlmax](https://github.com/piexlmax)
func (PResourcesService *PlayerResourcesService) GetPlayerResourcesInfoList(info PlayerResourcesReq.PlayerResourcesSearch) (list []PlayerResources.PlayerResources, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&PlayerResources.PlayerResources{})
	var PResourcess []PlayerResources.PlayerResources
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["user_id"] = true
	if orderMap[info.Sort] {
		OrderStr = info.Sort
		if info.Order == "descending" {
			OrderStr = OrderStr + " desc"
		}
		db = db.Order(OrderStr)
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&PResourcess).Error
	return PResourcess, total, err
}
