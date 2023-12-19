import service from '@/utils/request'

export const setPlayerResources = (data) => {
    return service({
        url: '/playerResources/setPlayerResources',
        method: 'post',
        data
    })
}
