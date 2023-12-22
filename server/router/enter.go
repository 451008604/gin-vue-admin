package router

import (
	playerresources "github.com/flipped-aurora/gin-vue-admin/server/router/PlayerResources"
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/serveraction"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
)

type RouterGroup struct {
	System          system.RouterGroup
	Example         example.RouterGroup
	PlayerResources playerresources.RouterGroup
	ServerAction    serveraction.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
