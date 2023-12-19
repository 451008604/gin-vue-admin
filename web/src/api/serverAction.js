import service from '@/utils/request'

export const updateServerConfig = (data) => {
    return service({
        url: '/serverAction/updateServerConfig',
        method: 'put',
        data
    })
}

export const updateFixedConfig = (data) => {
    return service({
        url: '/serverAction/updateFixedConfig',
        method: 'put',
        data
    })
}

export const setTimedTasks = (data) => {
    return service({
        url: '/serverAction/setTimedTasks',
        method: 'post',
        data
    })
}
