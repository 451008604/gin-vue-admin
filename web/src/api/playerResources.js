import service from '@/utils/request'

// @Tags PlayerResources
// @Summary 创建玩家资源
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PlayerResources true "创建玩家资源"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /PResources/createPlayerResources [post]
export const createPlayerResources = (data) => {
  return service({
    url: '/PResources/createPlayerResources',
    method: 'post',
    data
  })
}

// @Tags PlayerResources
// @Summary 删除玩家资源
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PlayerResources true "删除玩家资源"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /PResources/deletePlayerResources [delete]
export const deletePlayerResources = (data) => {
  return service({
    url: '/PResources/deletePlayerResources',
    method: 'delete',
    data
  })
}

// @Tags PlayerResources
// @Summary 批量删除玩家资源
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除玩家资源"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /PResources/deletePlayerResources [delete]
export const deletePlayerResourcesByIds = (data) => {
  return service({
    url: '/PResources/deletePlayerResourcesByIds',
    method: 'delete',
    data
  })
}

// @Tags PlayerResources
// @Summary 更新玩家资源
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PlayerResources true "更新玩家资源"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /PResources/updatePlayerResources [put]
export const updatePlayerResources = (data) => {
  return service({
    url: '/PResources/updatePlayerResources',
    method: 'put',
    data
  })
}

// @Tags PlayerResources
// @Summary 用id查询玩家资源
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.PlayerResources true "用id查询玩家资源"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /PResources/findPlayerResources [get]
export const findPlayerResources = (params) => {
  return service({
    url: '/PResources/findPlayerResources',
    method: 'get',
    params
  })
}

// @Tags PlayerResources
// @Summary 分页获取玩家资源列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取玩家资源列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /PResources/getPlayerResourcesList [get]
export const getPlayerResourcesList = (params) => {
  return service({
    url: '/PResources/getPlayerResourcesList',
    method: 'get',
    params
  })
}


export const setPlayerResources = (data) => {
  return service({
    url: '/PResources/setPlayerResources',
    method: 'post',
    data: data
  })
}