<template><a-button ghost @click="logout">退出 </a-button></template>

<script scoped>
export default {
  methods: {
    logout() {
      window.sessionStorage.clear('token')
      this.$router.push('login')
    }
  }
}
</script>
