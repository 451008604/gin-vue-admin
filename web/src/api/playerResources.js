import service from '@/utils/request'

export const setPlayerResources = (data) => {
  return service({
    url: '/PResources/setPlayerResources',
    method: 'post',
    data: data
  })
}