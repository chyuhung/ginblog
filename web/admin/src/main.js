import Vue from 'vue'
import App from './App.vue'
import router from './router'
import axios from 'axios'

import './plugin/antui'
import './assets/css/style.css'
import 'ant-design-vue/dist/antd.css'

Vue.prototype.$http = axios
axios.defaults.baseURL = 'http://localhost:3000/api/v1'
Vue.config.productionTip = false

new Vue({
  router,
  render: (h) => h(App)
}).$mount('#app')
