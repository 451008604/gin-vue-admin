<template>
  <el-form :model="formData" ref="vForm" :rules="rules" label-position="left" label-width="80px" @submit.prevent>
    <div class="table-container">
      <table class="table-layout">
        <tbody>
          <tr>
            <td class="table-cell" colspan="4" style="height: 100% !important">
              <div class="static-content-item"> 货币资源 </div>
            </td>
          </tr>
          <tr>
            <td class="table-cell">
              <el-form-item label="" label-width="0" prop="input100456">
                <el-input v-model="formData.input100456" type="text" placeholder="金币数量" clearable></el-input>
              </el-form-item>
            </td>
            <td class="table-cell">
              <el-form-item label="" label-width="0" prop="input108790">
                <el-input v-model="formData.input108790" type="text" placeholder="钻石数量" clearable></el-input>
              </el-form-item>
            </td>
            <td class="table-cell">
              <el-form-item label="" label-width="0" prop="input81625">
                <el-input v-model="formData.input81625" type="text" placeholder="体力数量" clearable></el-input>
              </el-form-item>
            </td>
            <td class="table-cell">
              <el-form-item label="" label-width="0" prop="input46774">
                <el-input v-model="formData.input46774" type="text" placeholder="经验数量" clearable></el-input>
              </el-form-item>
            </td>
          </tr>
          <tr>
            <td class="table-cell">
              <el-form-item label="" label-width="0" prop="input16999">
                <el-input v-model="formData.input16999" type="text" placeholder="工具数量" clearable></el-input>
              </el-form-item>
            </td>
            <td class="table-cell">
              <el-form-item label="" label-width="0" prop="input17994">
                <el-input v-model="formData.input17994" type="text" placeholder="护具数量" clearable></el-input>
              </el-form-item>
            </td>
            <td class="table-cell">
              <el-form-item label="" label-width="0" prop="input107097">
                <el-input v-model="formData.input107097" type="text" placeholder="油漆数量" clearable></el-input>
              </el-form-item>
            </td>
            <td class="table-cell">
              <el-form-item label="" label-width="0" prop="input89875">
                <el-input v-model="formData.input89875" type="text" placeholder="线材数量" clearable></el-input>
              </el-form-item>
            </td>
          </tr>
          <tr>
            <td class="table-cell" colspan="4" style="height: 100% !important">
              <div class="static-content-item"> 道具资源 </div>
            </td>
          </tr>
          <tr>
            <td class="table-cell" colspan="4">
              <el-form-item label="" label-width="0" prop="textarea47524">
                <el-input type="textarea" v-model="formData.textarea47524" placeholder="逗号分隔道具，冒号分隔ID和数量。（例：20001:2,20002:3 = 2个20001和3个20002）" rows="8" :show-word-limit="true"></el-input>
              </el-form-item>
            </td>
          </tr>
          <tr>
            <td class="table-cell" colspan="4">
              <el-divider />
              <div class="static-content-item"> 玩家列表 </div>
            </td>
          </tr>
          <tr>
            <td class="table-cell" colspan="4">
              <el-form-item label="" label-width="0" prop="textarea94373">
                <el-input type="textarea" v-model="formData.textarea94373" placeholder="玩家ID列表，通过逗号分隔" rows="3" :show-word-limit="true"></el-input>
              </el-form-item>
            </td>
          </tr>
          <tr>
            <td class="table-cell" colspan="3">
            </td>
            <td class="table-cell">
              <div class="static-content-item">
                <el-button style="width: 200px; height: 50px; font-size: x-large;" type="success" @click="submitForm" icon="success-filled">提交</el-button>
              </div>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </el-form>
</template>
  
<script>
import { ElMessage } from "element-plus";
import { setPlayerResources } from '@/api/playerResources.js'
import { defineComponent, toRefs, reactive, getCurrentInstance } from 'vue'

export default defineComponent({
  components: {},
  props: {},
  setup() {
    const state = reactive({
      formData: {
        input100456: "",
        input108790: "",
        input81625: "",
        input46774: "",
        input16999: "",
        input17994: "",
        input107097: "",
        input89875: "",
        textarea47524: "",
        textarea94373: "",
      },
      rules: {
        textarea47524: [{
          pattern: '^\\d+:\\d+(,\\d+:\\d+)*$',
          trigger: ['blur', 'change'],
          message: '格式不正确，请检查'
        }],
        textarea94373: [{
          pattern: '^\\d+(,\\d+)*$',
          trigger: ['blur', 'change'],
          message: '格式不正确，请检查'
        }],
      },
    })
    const instance = getCurrentInstance()
    const submitForm = () => {
      instance.ctx.$refs['vForm'].validate(async valid => {
        if (!valid) return
        const res = await setPlayerResources(state.formData);
        if (res.code === 0) {
          ElMessage({
            grouping: true,
            message: res.msg,
            type: 'success'
          })
        }
      })
    }
    const resetForm = () => {
      instance.ctx.$refs['vForm'].resetFields()
    }
    return {
      ...toRefs(state),
      submitForm,
      resetForm
    }
  }
})
</script>
  
<style lang="scss">
.el-input-number.full-width-input,
.el-cascader.full-width-input {
  width: 100% !important;
}

.el-form-item--medium {
  .el-radio {
    line-height: 36px !important;
  }

  .el-rate {
    margin-top: 8px;
  }
}

.el-form-item--small {
  .el-radio {
    line-height: 32px !important;
  }

  .el-rate {
    margin-top: 6px;
  }
}

.el-form-item--mini {
  .el-radio {
    line-height: 28px !important;
  }

  .el-rate {
    margin-top: 4px;
  }
}

.clear-fix:before,
.clear-fix:after {
  display: table;
  content: "";
}

.clear-fix:after {
  clear: both;
}

.float-right {
  float: right;
}
</style>

<style lang="scss" scoped>
div.table-container {
  table.table-layout {
    width: 100%;
    table-layout: fixed;
    border-collapse: collapse;
    border: 0;

    td.table-cell {
      display: table-cell;
      height: 36px;
      padding: 0 5px;
      border: 0;
    }
  }
}

.label-left-align :deep(.el-form-item__label) {
  text-align: left;
}

.label-center-align :deep(.el-form-item__label) {
  text-align: center;
}

.label-right-align :deep(.el-form-item__label) {
  text-align: right;
}

.static-content-item {
  min-height: 20px;
  display: flex;
  align-items: center;
  justify-content: center;
  margin: 10px 0;
  font-size: x-large;

  :deep(.el-divider--horizontal) {
    margin: 0;
  }
}
</style>
