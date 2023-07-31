<template>
  <div class="container">
    <div class="loginBox">
      <a-form-model class="loginForm">
        <a-form-model-item>
          <a-input placeholder="请输入用户名">
            <a-icon
              slot="prefix"
              type="user"
              style="color: rgba(0, 0, 0, 0.25)"
          /></a-input>
        </a-form-model-item>
        <a-form-model-item>
          <a-input placeholder="请输入密码"
            ><a-icon
              slot="prefix"
              type="lock"
              style="color: rgba(0, 0, 0, 0.25)"
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
