package serveraction

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/serveraction"
	"github.com/gin-gonic/gin"
)

func (a *ServerActionApi) SetTimedTasks(c *gin.Context) {
	var formData serveraction.PlanJson
	if err := c.ShouldBindJSON(&formData); err != nil {
		response.FailWithMessage("数据解析错误，请联系管理员", c)
	}
	strPlan := planJson2Str(formData)
	b, _ := json.Marshal(formData)
	SaveServerPlan(strPlan, string(b))
}

func planJson2Str(plan serveraction.PlanJson) string {
	// 停服时间（CLOSE_TIME）
	nCloseTime := Str2Time(plan.CLOSE_TIME)
	if nCloseTime < 1 {
		return ""
	}

	// 开放时间（OPEN_TIME）：最早为停服3分钟后
	nOpenTime := Str2Time(plan.OPEN_TIME)
	if nOpenTime < nCloseTime+180 {
		return ""
	}

	// 更新时间：固定为停服1分钟后
	// 启动时间：固定为更新1分钟后
	strPlan := strconv.Itoa(int(nCloseTime)) + ":0-"
	strPlan += strconv.Itoa(int(nCloseTime+55)) + ":2-"
	strPlan += strconv.Itoa(int(nCloseTime+115)) + ":1-"
	return strPlan
}

// 将str转为时间戳
func Str2Time(strTime string) uint32 {
	// 如果发生错误, 返回0
	t, err := time.ParseInLocation("2006-01-02 15:04:05", strTime, time.Local)
	if err != nil {
		return 0
	}

	return uint32(t.Unix())
}

// 写入定时开关服
func SaveServerPlan(plan, planJson string) bool {
	redisKey := "SERVER:PLAN"
	if err := global.GVA_REDIS.Set(context.Background(), redisKey, plan, 0).Err(); nil != err {
		return false
	}
	jsonKey := "SERVER:PLAN_JSON"
	if err := global.GVA_REDIS.Set(context.Background(), jsonKey, planJson, 0).Err(); nil != err {
		return false
	}
	return true
}

// 读取定时开关服
func LoadServerPlan() string {
	redisKey := "SERVER:PLAN_JSON"
	plan, err := global.GVA_REDIS.Get(context.Background(), redisKey).Result()
	if nil != err {
		return ""
	}
	return plan
}
