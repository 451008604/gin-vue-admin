<template>
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
</template>
  
<script>
import { ElMessage } from "element-plus";
import { updateServerConfig, updateFixedConfig } from '@/api/serverAction.js';
import { defineComponent, toRefs, reactive, getCurrentInstance } from 'vue'
export default defineComponent({
    components: {},
    props: {},
    setup() {
        const submitServerConfig = async () => {
            const res = await updateServerConfig();
            if (res.code === 0) {
                ElMessage({
                    grouping: true,
                    message: res.msg,
                    type: 'success'
                })
            }
        }
        const submitFixedConfig = async () => {
            const res = await updateFixedConfig();
            if (res.code === 0) {
                ElMessage({
                    grouping: true,
                    message: res.msg,
                    type: 'success'
                })
            }
        }
        return {
            submitServerConfig,
            submitFixedConfig
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