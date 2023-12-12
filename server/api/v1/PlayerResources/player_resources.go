package PlayerResources

import (
	"context"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/PlayerResources"
	PlayerResourcesReq "github.com/flipped-aurora/gin-vue-admin/server/model/PlayerResources/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	"strings"
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

type SetFormData struct {
	Input100456   string `json:"input100456"`   // 金币数量
	Input108790   string `json:"input108790"`   // 钻石数量
	Input81625    string `json:"input81625"`    // 体力数量
	Input46774    string `json:"input46774"`    // 经验数量
	Input16999    string `json:"input16999"`    // 工具数量
	Input17994    string `json:"input17994"`    // 护具数量
	Input107097   string `json:"input107097"`   // 油漆数量
	Input89875    string `json:"input89875"`    // 线材数量
	Textarea47524 string `json:"textarea47524"` // 道具资源
	Textarea94373 string `json:"textarea94373"` // 玩家ID列表
}

type MailJson struct {
	Roleid        []uint32    // 角色ID
	Prize         []ItemPrize // 奖励
	DescribeID    int32       // 描述ID
	ExpireTime    uint32      // 过期时间, 时间戳
	StrExpireTime string      // 可读性较强的时间
	AliveTime     uint32      // 存活时间, 单位秒
}

// 道具奖励
type ItemPrize struct {
	ID  int
	Num int
}

func (PResourcesApi *PlayerResourcesApi) SetPlayerResources(c *gin.Context) {
	var formData SetFormData
	if err := c.ShouldBindJSON(&formData); err != nil {
		response.FailWithMessage("数据解析错误，请联系管理员", c)
	}

	fmt.Printf("%+v\n", formData)
	// 检查用户ID是否有效
	roleList := strings.Split(formData.Textarea94373, ",")
	// 用于RoleID去重校验
	redisRoleKey := map[string]string{}
	for i := range roleList {
		redisRoleKey[roleList[i]] = "roledata:" + roleList[i]
		roleList[i] = redisRoleKey[roleList[i]]
	}
	count, err := global.GVA_REDIS.Exists(context.Background(), roleList...).Result()
	if len(redisRoleKey) != int(count) || err != nil {
		response.FailWithMessage("用户ID有误，请检查", c)
		return
	}

	// 格式化资源列表
	var prizeData []ItemPrize
	for i := 1; i < 9; i++ {
		data := ItemPrize{ID: i}
		switch data.ID {
		case 1: // 金币
			data.Num, _ = strconv.Atoi(formData.Input100456)
		case 2: // 钻石
			data.Num, _ = strconv.Atoi(formData.Input108790)
		case 3: // 体力
			data.Num, _ = strconv.Atoi(formData.Input81625)
		case 4: // 经验
			data.Num, _ = strconv.Atoi(formData.Input46774)
		case 5: // 工具
			data.Num, _ = strconv.Atoi(formData.Input16999)
		case 6: // 护具
			data.Num, _ = strconv.Atoi(formData.Input17994)
		case 7: // 油漆
			data.Num, _ = strconv.Atoi(formData.Input107097)
		case 8: // 线材
			data.Num, _ = strconv.Atoi(formData.Input89875)
		default:
			response.FailWithMessage("无效的资源类型，请联系管理员添加", c)
		}
		if data.Num < 1 {
			continue
		}
		prizeData = append(prizeData, data)
	}
	for _, itemData := range strings.Split(formData.Textarea47524, ",") {
		split := strings.Split(itemData, ":")
		if len(split) < 2 {
			continue
		}
		prize := ItemPrize{}
		prize.ID, _ = strconv.Atoi(split[0])
		prize.Num, _ = strconv.Atoi(split[1])
		if prize.Num < 1 {
			continue
		}
		prizeData = append(prizeData, prize)
	}

	// 缓存玩家邮件数据
	for k := range redisRoleKey {
		mailData := MailJson{DescribeID: 51, ExpireTime: 2556115199, StrExpireTime: "2050-12-31 23:59:59", Prize: prizeData}
		strMail := Mail2Str(mailData)
		global.GVA_REDIS.Append(context.Background(), "mail:"+k, strMail)
	}

	// 广播邮件数据处理通知（将所有玩家ID广播出去，各服务自行读取缓存的当前在线玩家邮件数据）
	err = global.GVA_REDIS.Publish(context.Background(), "subscribe_mail", formData.Textarea94373).Err()
	if err != nil {
		response.FailWithMessage("广播失败邮件无法实时到账，等待玩家重新登录时会收到", c)
	}
	response.Ok(c)
}

// 序列化邮件
func Mail2Str(mailJson MailJson) string {
	strMail := strconv.Itoa(int(mailJson.DescribeID)) + "-"
	strMail += strconv.Itoa(int(mailJson.ExpireTime)) + "-"
	strMail += strconv.Itoa(int(mailJson.AliveTime)) + "-"
	for _, v := range mailJson.Prize {
		strMail += strconv.Itoa(v.ID) + ":" + strconv.Itoa(v.Num) + ","
	}
	return strMail + ";"
}
