package serveraction

import (
	"context"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
)

type ServerActionApi struct {
}

func (a *ServerActionApi) UpdateServerConfig(c *gin.Context) {
	err := global.GVA_REDIS.Publish(context.Background(), "subscribe_server_config", "").Err()
	if err != nil {
		response.FailWithMessage("指令发送失败，请联系管理员", c)
	}
	response.Ok(c)
}

func (a *ServerActionApi) UpdateFixedConfig(c *gin.Context) {
	err := global.GVA_REDIS.Publish(context.Background(), "subscribe_fixed_config", "").Err()
	if err != nil {
		response.FailWithMessage("指令发送失败，请联系管理员", c)
	}
	response.Ok(c)
}
