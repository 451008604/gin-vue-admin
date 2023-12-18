<template>
  <div id="userLayout" onselectstart="return false;" class="w-full h-full relative bg">
    <div class="md:w-5/5 w-12/12 h-full flex items-center justify-evenly">
      <!-- 分割斜块 -->
      <div class="z-[999] p-12 md:w-96 w-full flex flex-col justify-between box-border" style="background-color: rgba(255,255,255,1);border-radius: 20px;">
        <div class="flex items-center justify-center">
          <img class="w-36" src="@/assets/icon256.png" style="border-radius: 20px;" draggable="false" alt>
        </div>
        <div class="m-10 mt-5">
          <p class="text-center text-xl font-bold">{{ $GIN_VUE_ADMIN.appName }}</p>
        </div>
        <el-form ref="loginForm" :model="loginFormData" :rules="rules" :validate-on-rule-change="false" @keyup.enter="submitForm">
          <el-form-item prop="username" class="mb-6">
            <el-input v-model="loginFormData.username" size="large" placeholder="请输入用户名" suffix-icon="user" />
          </el-form-item>
          <el-form-item prop="password" class="mb-6">
            <el-input v-model="loginFormData.password" show-password size="large" type="password" placeholder="请输入密码" />
          </el-form-item>
          <el-form-item v-if="loginFormData.openCaptcha" prop="captcha" class="mb-6">
            <div class="flex w-full justify-between">
              <el-input v-model="loginFormData.captcha" placeholder="请输入验证码" size="large" class="flex-1 mr-5" />
              <div class="w-1/3 h-11 bg-[#c3d4f2] rounded">
                <img v-if="picPath" class="w-full h-full" :src="picPath" alt="请输入验证码" @click="loginVerify()">
              </div>
            </div>
          </el-form-item>
          <el-form-item class="mb-0">
            <el-button class="shadow shadow-blue-600 h-11 w-full" type="primary" size="large" @click="submitForm">登 录</el-button>
          </el-form-item>
        </el-form>
      </div>
    </div>
  </div>
</template>

<script setup>
import { captcha } from '@/api/user'
import { checkDB } from '@/api/initdb'
import BottomInfo from '@/view/layout/bottomInfo/bottomInfo.vue'
import { reactive, ref } from 'vue'
import { ElMessage } from 'element-plus'
import { useRouter } from 'vue-router'
import { useUserStore } from '@/pinia/modules/user'

defineOptions({
  name: 'Login',
})

const router = useRouter()
// 验证函数
const checkUsername = (rule, value, callback) => {
  if (value.length < 5) {
    return callback(new Error('请输入正确的用户名'))
  } else {
    callback()
  }
}
const checkPassword = (rule, value, callback) => {
  if (value.length < 6) {
    return callback(new Error('请输入正确的密码'))
  } else {
    callback()
  }
}

// 获取验证码
const loginVerify = () => {
  captcha({}).then(async (ele) => {
    rules.captcha.push({
      max: ele.data.captchaLength,
      min: ele.data.captchaLength,
      message: `请输入${ele.data.captchaLength}位验证码`,
      trigger: 'blur',
    })
    picPath.value = ele.data.picPath
    loginFormData.captchaId = ele.data.captchaId
    loginFormData.openCaptcha = ele.data.openCaptcha
  })
}
loginVerify()

// 登录相关操作
const loginForm = ref(null)
const picPath = ref('')
const loginFormData = reactive({
  username: 'guohaoqin',
  password: 'guohaoqin',
  captcha: '',
  captchaId: '',
  openCaptcha: false,
})
const rules = reactive({
  username: [{ validator: checkUsername, trigger: 'blur' }],
  password: [{ validator: checkPassword, trigger: 'blur' }],
  captcha: [
    {
      message: '验证码格式不正确',
      trigger: 'blur',
    },
  ],
})

const userStore = useUserStore()
const login = async () => {
  return await userStore.LoginIn(loginFormData)
}
const submitForm = () => {
  loginForm.value.validate(async (v) => {
    if (v) {
      const flag = await login()
      if (!flag) {
        loginVerify()
      }
    } else {
      ElMessage({
        type: 'error',
        message: '请正确填写登录信息',
        showClose: true,
      })
      loginVerify()
      return false
    }
  })
}

// 跳转初始化
const checkInit = async () => {
  const res = await checkDB()
  if (res.code === 0) {
    if (res.data?.needInit) {
      userStore.NeedInit()
      router.push({ name: 'Init' })
    } else {
      ElMessage({
        type: 'info',
        message: '已配置数据库信息，无法初始化',
      })
    }
  }
}

</script>

<style>
.bg {
  widows: 100%;
  height: 100%;
  background-image: url('/src/assets/backgroud.png');
  background-repeat: no-repeat;
  background-size: cover;
  background-position: center center;
}
</style>
