<template>
    <nav class="navbar">
      <div class="navbar-container">
        <!-- 网站Logo -->
        <div class="navbar-logo">
          <router-link to="/" class="logo">GoBlog</router-link>
        </div>
  
        <!-- 导航菜单 -->
        <ul class="navbar-menu">
          <li>
            <router-link to="/" class="navbar-item">
              <span>首页</span>
              <div class="nav-item-line"></div>
            </router-link>
          </li>
          <li>
            <router-link to="/about" class="navbar-item">
              <span>关于</span>
              <div class="nav-item-line"></div>
            </router-link>
          </li>
          <!-- 显示登录注册或用户信息 -->
          <li v-if="!isAuthenticated">
            <router-link to="/login" class="navbar-item">
              <span>登录</span>
              <div class="nav-item-line"></div>
            </router-link>
          </li>
          <li v-if="!isAuthenticated">
            <router-link to="/register" class="navbar-item">
              <span>注册</span>
              <div class="nav-item-line"></div>
            </router-link>
          </li>
          <li v-if="isAuthenticated">
            <router-link to="/profile" class="navbar-item">
              <span>我的博客</span>
              <div class="nav-item-line"></div>
            </router-link>
          </li>
          <li v-if="isAuthenticated">
            <button @click="handleLogout" class="navbar-item logout-btn">
              <span>退出</span>
              <div class="nav-item-line"></div>
            </button>
          </li>
        </ul>
      </div>
    </nav>
  </template>
  
  <script setup>
  import { ref } from 'vue';
  import { useRouter } from 'vue-router';
  
  // 模拟用户登录状态
  const isAuthenticated = ref(false);
  
  // 处理退出功能
  const router = useRouter();
  const handleLogout = () => {
    // 清除Token或相关的认证信息
    localStorage.removeItem('token');
    isAuthenticated.value = false;
    router.push('/login');
  };
  </script>
  
  <style scoped>
  .navbar {
    background-color: #222222;
    padding: 1rem 0;
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    z-index: 1000;
    border-bottom: 1px solid #333333;
  }
  
  .navbar::before {
    content: '';
    position: absolute;
    top: 0;
    left: 0;
    right: 0;
    height: 4px;
    background: linear-gradient(90deg, #00a8ff, #00ff7f);
  }
  
  .navbar-container {
    display: flex;
    justify-content: space-between;
    align-items: center;
    width: 90%;
    max-width: 1200px;
    margin: 0 auto;
    padding: 0 2rem;
  }
  
  .navbar-logo .logo {
    font-size: 1.5rem;
    color: #ffffff;
    font-weight: 700;
    text-decoration: none;
    text-shadow: 0 0 10px rgba(0, 168, 255, 0.3);
    transition: color 0.3s ease;
  }
  
  .navbar-logo .logo:hover {
    color: #00a8ff;
  }
  
  .navbar-menu {
    list-style-type: none;
    display: flex;
    gap: 2rem;
    margin: 0;
    padding: 0;
  }
  
  .navbar-item {
    color: #888888;
    text-decoration: none;
    font-size: 1rem;
    position: relative;
    padding: 0.5rem 0;
    transition: color 0.3s ease;
    background: none;
    border: none;
    cursor: pointer;
    display: inline-block;
  }
  
  .nav-item-line {
    position: absolute;
    bottom: 0;
    left: 0;
    width: 100%;
    height: 2px;
    background: linear-gradient(90deg, #00a8ff, #00ff7f);
    transform: scaleX(0);
    transition: transform 0.3s ease;
    transform-origin: right;
  }
  
  .navbar-item:hover {
    color: #ffffff;
  }
  
  .navbar-item:hover .nav-item-line {
    transform: scaleX(1);
    transform-origin: left;
  }
  
  .router-link-active {
    color: #ffffff;
  }
  
  .router-link-active .nav-item-line {
    transform: scaleX(1);
  }
  
  .logout-btn {
    font-size: 1rem;
    color: #888888;
  }
  
  .logout-btn:hover {
    color: #ffffff;
  }
  
  @media (max-width: 768px) {
    .navbar-container {
      padding: 0 1rem;
    }
    
    .navbar-menu {
      gap: 1rem;
    }
    
    .navbar-item {
      font-size: 0.9rem;
    }
  }
  </style>