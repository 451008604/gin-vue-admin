package PlayerResources

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/PlayerResources"
	PlayerResourcesReq "github.com/flipped-aurora/gin-vue-admin/server/model/PlayerResources/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type PlayerResourcesApi struct {
}

var PResourcesService = service.ServiceGroupApp.PlayerResourcesServiceGroup.PlayerResourcesService

// CreatePlayerResources 创建玩家资源
// @Tags PlayerResources
// @Summary 创建玩家资源
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body PlayerResources.PlayerResources true "创建玩家资源"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /PResources/createPlayerResources [post]
func (PResourcesApi *PlayerResourcesApi) CreatePlayerResources(c *gin.Context) {
	var PResources PlayerResources.PlayerResources
	err := c.ShouldBindJSON(&PResources)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	PResources.CreatedBy = utils.GetUserID(c)
	verify := utils.Rules{
		"UserId":     {utils.NotEmpty()},
		"ProperInfo": {utils.NotEmpty()},
	}
	if err := utils.Verify(PResources, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := PResourcesService.CreatePlayerResources(&PResources); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeletePlayerResources 删除玩家资源
// @Tags PlayerResources
// @Summary 删除玩家资源
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body PlayerResources.PlayerResources true "删除玩家资源"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /PResources/deletePlayerResources [delete]
func (PResourcesApi *PlayerResourcesApi) DeletePlayerResources(c *gin.Context) {
	var PResources PlayerResources.PlayerResources
	err := c.ShouldBindJSON(&PResources)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	PResources.DeletedBy = utils.GetUserID(c)
	if err := PResourcesService.DeletePlayerResources(PResources); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeletePlayerResourcesByIds 批量删除玩家资源
// @Tags PlayerResources
// @Summary 批量删除玩家资源
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除玩家资源"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /PResources/deletePlayerResourcesByIds [delete]
func (PResourcesApi *PlayerResourcesApi) DeletePlayerResourcesByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	deletedBy := utils.GetUserID(c)
	if err := PResourcesService.DeletePlayerResourcesByIds(IDS, deletedBy); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdatePlayerResources 更新玩家资源
// @Tags PlayerResources
// @Summary 更新玩家资源
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body PlayerResources.PlayerResources true "更新玩家资源"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /PResources/updatePlayerResources [put]
func (PResourcesApi *PlayerResourcesApi) UpdatePlayerResources(c *gin.Context) {
	var PResources PlayerResources.PlayerResources
	err := c.ShouldBindJSON(&PResources)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	PResources.UpdatedBy = utils.GetUserID(c)
	verify := utils.Rules{
		"UserId":     {utils.NotEmpty()},
		"ProperInfo": {utils.NotEmpty()},
	}
	if err := utils.Verify(PResources, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := PResourcesService.UpdatePlayerResources(PResources); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindPlayerResources 用id查询玩家资源
// @Tags PlayerResources
// @Summary 用id查询玩家资源
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query PlayerResources.PlayerResources true "用id查询玩家资源"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /PResources/findPlayerResources [get]
func (PResourcesApi *PlayerResourcesApi) FindPlayerResources(c *gin.Context) {
	var PResources PlayerResources.PlayerResources
	err := c.ShouldBindQuery(&PResources)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rePResources, err := PResourcesService.GetPlayerResources(PResources.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rePResources": rePResources}, c)
	}
}

// GetPlayerResourcesList 分页获取玩家资源列表
// @Tags PlayerResources
// @Summary 分页获取玩家资源列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query PlayerResourcesReq.PlayerResourcesSearch true "分页获取玩家资源列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /PResources/getPlayerResourcesList [get]
func (PResourcesApi *PlayerResourcesApi) GetPlayerResourcesList(c *gin.Context) {
	var pageInfo PlayerResourcesReq.PlayerResourcesSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := PResourcesService.GetPlayerResourcesInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
