package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/PlayerResources"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup          system.ApiGroup
	ExampleApiGroup         example.ApiGroup
	PlayerResourcesApiGroup PlayerResources.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
