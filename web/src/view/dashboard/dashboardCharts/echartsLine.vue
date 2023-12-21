<template>
  <div class="dashboard-line-box">
    <div class="dashboard-line-title"></div>
    <div ref="echart" class="dashboard-line" />
  </div>
</template>

<script setup>
import * as echarts from 'echarts'
import { nextTick, onMounted, onUnmounted, ref, shallowRef } from 'vue'
import 'echarts/theme/macarons'

const chart = shallowRef(null)
const echart = ref(null)
const oneMinute = 60 * 1000;
let data = [];
let now = new Date();
let value = Math.random() * 1000;

const randomData = () => {
  now = new Date(+now + oneMinute);
  value = value + Math.random() * 21 - 10;
  return {
    name: now.toString(),
    value: [now.getDate() + "日" + now.getHours() + ":" + now.getMinutes(), Math.round(value)]
  };
}
const setOption = () => {
  if (chart.value == null) {
    clearInterval(interval)
    return
  }
  chart.value.setOption({
    grid: {
      left: '40',
      right: '20',
      top: '20',
      bottom: '20',
    },
    tooltip: {
      trigger: 'axis',
      formatter: function (params) {
        params = params[0];
        var date = new Date(params.name);
        return (
          date.getHours() + ':' + (date.getMinutes()) + '  ' + params.value[1]
        );
      },
      axisPointer: {
        animation: false
      }
    },
    xAxis: {
      type: 'category',
      splitLine: {
        show: false
      }
    },
    yAxis: {
      type: 'value',
      boundaryGap: [0, '100%'],
      splitLine: {
        show: false
      }
    },
    dataZoom: [
      {
        type: "inside",
      }
    ],
    series: [
      {
        type: 'line',
        showSymbol: false,
        data: data
      }
    ]
  })
}

for (var i = 0; i < 24 * 60; i++) {
  data.push(randomData());
}

const initChart = () => {
  chart.value = echarts.init(echart.value, 'macarons')
  setOption();
}

var interval = setInterval(function () {
  data.shift();
  data.push(randomData());

  setOption();
}, 60 * 1000);

onMounted(async () => {
  await nextTick()
  initChart()
})

onUnmounted(() => {
  if (!chart.value) {
    return
  }
  chart.value.dispose()
  chart.value = null
})
</script>

<style lang="scss" scoped>
.dashboard-line-box {
  .dashboard-line {
    background-color: #fff;
    height: 360px;
    width: 100%;
  }

  .dashboard-line-title {
    font-weight: 600;
    margin-bottom: 12px;
  }
}
</style>
