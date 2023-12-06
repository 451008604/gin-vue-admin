
import axios from 'axios'
import { ref } from 'vue'

const weatherInfo = ref('天气预报努力加载中...')
const amapKey = '03ef28a6a2f106607c6a8991452a23a3'

export const useWeatherInfo = () => {
  ip()
  return weatherInfo
}

export const ip = async () => {
  // key换成你自己的 https://console.amap.com/dev/index
  if (amapKey === '') {
    return false
  }
  const res = await axios.get('https://restapi.amap.com/v3/ip?key=' + amapKey)
  if (res.data.adcode) {
    getWeather(res.data.adcode)
  }
}

const getWeather = async (code) => {
  const response = await axios.get('https://restapi.amap.com/v3/weather/weatherInfo?key=' + amapKey + '&extensions=base&city=' + code)
  if (response.data.status === '1') {
    const s = response.data.lives[0]
    weatherInfo.value = s.city + ' / ' + s.weather + ' / ' + s.temperature + 'C˚ / ' + s.winddirection + '风' + s.windpower + '级 / 空气湿度' + s.humidity + '%'
  }
}
