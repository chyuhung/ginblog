import Vue from 'vue'
import App from './App.vue'
import router from './router'
import axios from 'axios'

import './plugin/antui'
import './assets/css/style.css'
import 'ant-design-vue/dist/antd.css'

axios.defaults.baseURL = 'http://localhost:3000/api/v1'

// 请求拦截器，添加token
axios.interceptors.request.use((config) => {
  // 获取token
  const token = window.sessionStorage.getItem('token')
  config.headers.Authorization = `Bearer ${token}`
  return config
})
Vue.prototype.$http = axios
Vue.config.productionTip = false

new Vue({
  router,
  render: (h) => h(App)
}).$mount('#app')
