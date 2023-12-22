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
      right: '5',
      top: '20',
      bottom: '20',
    },
    tooltip: {
      trigger: 'axis',
      axisPointer: {
        animation: false
      }
    },
    toolbox: {
      show: true,
      itemSize: 20,
      feature: {
        dataZoom: {
          yAxisIndex: 'none'
        },
        dataView: { readOnly: true },
        magicType: { type: ['line', 'bar'] },
        restore: {},
        saveAsImage: {},
      }
    },
    animation: false,
    legend: {
      itemGap: 20,
      itemWidth: 20,
      itemHeight: 10,
    },
    xAxis: {
      type: 'category',
    },
    yAxis: {
      type: 'value',
      boundaryGap: [0, '30%'],
    },
    dataZoom: [
      {
        type: "inside",
        filterMode: "none",
        minSpan: "3",
        zoomOnMouseWheel: false,
      }
    ],
    series: [
      {
        name: "服务1",
        type: 'line',
        showSymbol: false,
        data: data,
      },
      {
        name: "服务2",
        type: 'bar',
        showSymbol: false,
        data: data
      }
    ]
  })
}

// 24 * 60 = 一天的分钟数
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
