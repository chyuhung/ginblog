<template>
  <a-layout class="container">
    <a-layout-sider>
      <Nav></Nav>
    </a-layout-sider>
    <a-layout>
      <a-layout-header class="headerBtn"><Header></Header></a-layout-header>
      <a-layout-content>
        <RouterView></RouterView>
      </a-layout-content>
      <a-layout-footer>
        <Footer></Footer>
      </a-layout-footer>
    </a-layout>
  </a-layout>
</template>
<script>
import Nav from '../components/admin/Nav.vue'
import Footer from '../components/admin/Footer.vue'
import Header from '../components/admin/Header.vue'
import { RouterView } from 'vue-router'
export default {
  components: { Nav, Footer, Header, RouterView }
}
</script>

<style scoped>
.container {
  height: 100%;
}
.headerBtn {
  display: flex;
  justify-content: flex-end;
  align-items: center;
}
</style>
