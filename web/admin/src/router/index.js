import Vue from 'vue'
import VueRouter from 'vue-router'
import Login from '../views/Login'
import Admin from '../views/Admin'

Vue.use(VueRouter)

const routes = [
  {
    path: '/login',
    name: 'Login',
    component: Login
  },
  {
    path: '/admin',
    name: 'Admin',
    component: Admin
  }
]

const router = new VueRouter({
  routes
})

export default router
