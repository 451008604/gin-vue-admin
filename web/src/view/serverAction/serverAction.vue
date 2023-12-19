<template>
    <div>
        <el-form label-position="left" label-width="200px" size="large">
            <el-row class="text-algin:center">
                <el-col :span="12" class="grid-cell">
                    <div class="static-content-item">
                        <el-button type="primary" class="mt-5 mb-5" plain icon="Files" @click="submitServerConfig">服务端配置热更</el-button>
                    </div>
                </el-col>
                <el-col :span="12" class="grid-cell">
                    <div class="static-content-item">
                        <el-button type="primary" class="mt-5 mb-5" plain icon="Files" @click="submitFixedConfig">静态配置热更</el-button>
                    </div>
                </el-col>
            </el-row>
        </el-form>
        <el-collapse v-model="collapseActive">
            <el-collapse-item name="1">
                <template #title><span class="m-5 text-xl font-bold"> 定时计划 </span></template>
                <el-form class="p-5 pl-20 pr-20" :model="formData" ref="vForm" :rules="rules" label-position="right" label-width="80px" size="default" @submit.prevent>
                    <el-form-item label="白名单" prop="WHITE_IP" class="required label-right-align">
                        <el-input v-model="formData.WHITE_IP" type="text" placeholder="建议：公司的公网IP地址" clearable></el-input>
                    </el-form-item>
                    <el-form-item label="持续时间" prop="durationTime" class="required label-right-align">
                        <el-date-picker is-range v-model="formData.durationTime" type="datetimerange" class="full-width-input" format="YYYY-MM-DD HH:mm:ss" value-format="YYYY-MM-DD HH:mm:ss" start-placeholder="停服时间" end-placeholder="开放时间" clearable></el-date-picker>
                    </el-form-item>
                    <el-form-item label="公告标题" prop="TITLE" class="required label-right-align">
                        <el-input v-model="formData.TITLE" type="text" clearable :maxlength="20" :show-word-limit="true"></el-input>
                    </el-form-item>
                    <el-form-item label="公告内容" prop="BODY" class="required label-right-align">
                        <el-input type="textarea" v-model="formData.BODY" rows="4" :maxlength="200" :show-word-limit="true"></el-input>
                    </el-form-item>
                    <el-form-item label="按钮文字" prop="BUTTON" class="required label-right-align">
                        <el-input v-model="formData.BUTTON" type="text" clearable :maxlength="5"></el-input>
                    </el-form-item>
                    <div class="static-content-item">
                        <el-button style="width: 200px; height: 50px; font-size: x-large;" type="success" @click="submitForm" icon="success-filled">提交</el-button>
                    </div>
                </el-form>
            </el-collapse-item>
        </el-collapse>
    </div>
</template>

<script>
import { ElMessage } from "element-plus";
import { updateServerConfig, updateFixedConfig, setTimedTasks } from '@/api/serverAction.js';
import { defineComponent, toRefs, reactive, getCurrentInstance } from 'vue'
export default defineComponent({
    components: {},
    props: {},
    setup() {
        const state = reactive({
            collapseActive: ['1'],
            formData: {
                durationTime: "",
                WHITE_IP: "",
                CLOSE_TIME: "",
                OPEN_TIME: "",
                TITLE: "停服公告",
                BODY: "非常抱歉，当前正处于停服更新中，敬请期待",
                BUTTON: "确定",
            },
            rules: {
                WHITE_IP: [{
                    pattern: '^((25[0-5]|2[0-4][0-9]|[0-1]?[0-9][0-9]?)\\.){3}(25[0-5]|2[0-4][0-9]|[0-1]?[0-9][0-9]?)$',
                    trigger: ['blur', 'change'],
                    required: true,
                    message: '格式不正确，请检查',
                }],
                CLOSE_TIME: [{
                    required: true,
                    message: '字段值不可为空',
                }],
                TITLE: [{
                    required: true,
                    message: '字段值不可为空',
                }],
                BODY: [{
                    required: true,
                    message: '字段值不可为空',
                }],
                BUTTON: [{
                    required: true,
                    message: '字段值不可为空',
                }],
            },
        })

        const success = (res) => {
            if (res.code === 0) {
                ElMessage({ grouping: true, message: res.msg, type: 'success' })
            }
        }

        const submitServerConfig = async () => {
            const res = await updateServerConfig();
            success(res)
        }
        const submitFixedConfig = async () => {
            const res = await updateFixedConfig();
            success(res)
        }
        const instance = getCurrentInstance()
        const submitForm = async () => {
            instance.ctx.$refs['vForm'].validate(async valid => {
                if (!valid) return
                state.formData.WHITE_IP = state.formData.WHITE_IP.trim()
                state.formData.CLOSE_TIME = "" + state.formData.durationTime[0]
                state.formData.OPEN_TIME = "" + state.formData.durationTime[1]
                console.log(state.formData);
                const res = await setTimedTasks(state.formData);
                success(res)
            })
        }
        return {
            ...toRefs(state),
            submitServerConfig,
            submitFixedConfig,
            submitForm
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

        td.table-cell {
            display: table-cell;
            height: 36px;
            border: 1px solid #e1e2e3;
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

    :deep(.el-divider--horizontal) {
        margin: 0;
    }
}
</style>