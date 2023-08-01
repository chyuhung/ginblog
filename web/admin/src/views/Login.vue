<template>
  <div class="container">
    <div class="loginBox">
      <a-form-model :rules="rules" :model="formdata" class="loginForm">
        <a-form-model-item prop="username">
          <a-input v-model="formdata.username" placeholder="用户名">
            <a-icon slot="prefix" type="user" style="color: rgba(0, 0, 0, 0.25)"
          /></a-input>
        </a-form-model-item>
        <a-form-model-item prop="password">
          <a-input v-model="formdata.password" placeholder="密码" type="password">
            <a-icon slot="prefix" type="lock" style="color: rgba(0, 0, 0, 0.25)"
          /></a-input>
        </a-form-model-item>
        <a-form-model-item class="loginBtn">
          <a-button type="primary" style="margin: 10px">登录</a-button>
          <a-button type="info" style="margin: 10px">取消</a-button>
        </a-form-model-item>
      </a-form-model>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      formdata: {
        username: '',
        password: ''
      },
      rules: {
        username: [
          { required: true, message: '请输入用户名', trigger: 'blur' },
          { min: 4, max: 12, message: '用户名必须在4到12个字符之间', trigger: 'blur' }
        ],
        password: [
          { required: true, message: '请输入密码', trigger: 'blur' },
          { min: 6, max: 20, message: '密码必须在6到20个字符之间', trigger: 'blur' }
        ]
      }
    }
  },
  methods: {
    onSubmit() {
      this.$refs.ruleForm.validate((valid) => {
        if (valid) {
          alert('submit!')
        } else {
          console.log('error submit!!')
          return false
        }
      })
    },
    resetForm() {
      this.$refs.ruleForm.resetFields()
    }
  }
}
</script>

<style scoped>
.container {
  height: 100%;
  background-color: rgb(75, 78, 106);
}
.loginBox {
  width: 450px;
  height: 300px;
  background-color: rgb(255, 255, 255);
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  border-radius: 9px;
}

.loginForm {
  width: 100%;
  position: absolute;
  bottom: 10%;
  padding: 0% 20%;
  box-sizing: border-box;
}
</style>
