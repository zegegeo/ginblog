<template>
  <div>
    <!-- 判断是否显示 Navbar -->
    <Navbar v-if="shouldShowNavbar" />
    <!-- 渲染路由 -->
    <router-view />
  </div>
</template>

<script>
import Navbar from './components/Navbar.vue';

export default {
  components: {
    Navbar,
  },
  computed: {
    shouldShowNavbar() {
      // 博客页面不显示导航栏
      const hideNavbarPages = ['/blog'];
      return !hideNavbarPages.includes(this.$route.path);
    },
  },
};
</script>

<style>
/* 重置默认样式 */
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

body {
  background: #1a1a1a;
  color: #ffffff;
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, 'Open Sans', 'Helvetica Neue', sans-serif;
}
</style>