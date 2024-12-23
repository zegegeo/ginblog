<template>
  <div class="blog-container">
    <nav class="blog-nav">
      <div class="nav-content">
        <h2 class="nav-logo">GoBlog</h2>
        <div class="nav-actions">
          <button class="btn create-btn">
            <span>写文章</span>
            <div class="btn-bg"></div>
          </button>
          <div class="user-menu">
            <div class="user-avatar">
              <img src="https://via.placeholder.com/40" alt="用户头像" />
            </div>
            <div class="dropdown-menu">
              <button class="logout-btn" @click="handleLogout">
                <span>退出登录</span>
                <div class="btn-bg"></div>
              </button>
            </div>
          </div>
        </div>
      </div>
    </nav>

    <main class="main-content">
      <div class="content-wrapper">
        <!-- 博客列表 -->
        <div class="blog-list">
          <article class="blog-card" v-for="i in 5" :key="i">
            <div class="blog-info">
              <div class="blog-meta">
                <span class="author">作者名称</span>
                <span class="date">2024-03-21</span>
              </div>
              <h3 class="blog-title">这是一个示例博客标题，展示文章的主要内容</h3>
              <p class="blog-excerpt">
                这里是博客的摘要内容，简短介绍文章的主要内容���吸引读者点击阅读更多...
              </p>
              <div class="blog-tags">
                <span class="tag">技术</span>
                <span class="tag">编程</span>
                <span class="tag">Vue.js</span>
              </div>
            </div>
          </article>
        </div>
      </div>
    </main>
  </div>
</template>

<script setup>
import { useRouter } from 'vue-router';

const router = useRouter();

const handleLogout = () => {
  // 清除登录状态
  localStorage.removeItem('token');
  // 跳转到首页
  router.push('/');
};
</script>

<style scoped>
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

.blog-container {
  min-height: 100vh;
  background: #1a1a1a;
  color: #ffffff;
}

/* 导航栏样式 */
.blog-nav {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  background: #222222;
  padding: 1rem 0;
  border-bottom: 1px solid #333333;
  z-index: 100;
}

.blog-nav::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 4px;
  background: linear-gradient(90deg, #00a8ff, #00ff7f);
}

.nav-content {
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 2rem;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.nav-logo {
  color: #ffffff;
  font-size: 1.5rem;
  font-weight: 700;
  text-shadow: 0 0 10px rgba(0, 168, 255, 0.3);
}

.nav-actions {
  display: flex;
  align-items: center;
  gap: 1.5rem;
}

/* 主要内容区域 */
.main-content {
  padding-top: 5rem;
  min-height: calc(100vh - 4rem);
}

.content-wrapper {
  max-width: 1200px;
  margin: 0 auto;
  padding: 2rem;
}

/* 博客卡片样式 */
.blog-card {
  background: #222222;
  border-radius: 12px;
  padding: 2rem;
  margin-bottom: 1.5rem;
  transition: transform 0.3s ease, box-shadow 0.3s ease;
  cursor: pointer;
}

.blog-card:hover {
  transform: translateY(-3px);
  box-shadow: 0 4px 20px rgba(0, 168, 255, 0.1);
}

.blog-meta {
  display: flex;
  gap: 1rem;
  margin-bottom: 1rem;
  color: #888888;
  font-size: 0.9rem;
}

.blog-title {
  color: #ffffff;
  font-size: 1.5rem;
  margin-bottom: 1rem;
  line-height: 1.4;
}

.blog-excerpt {
  color: #888888;
  margin-bottom: 1.5rem;
  line-height: 1.6;
}

.blog-tags {
  display: flex;
  gap: 0.5rem;
}

.tag {
  background: #2a2a2a;
  color: #00a8ff;
  padding: 0.3rem 0.8rem;
  border-radius: 4px;
  font-size: 0.8rem;
}

/* 按钮样式 */
.btn {
  padding: 0.8rem 1.5rem;
  border: none;
  border-radius: 8px;
  font-size: 1rem;
  cursor: pointer;
  position: relative;
  overflow: hidden;
  transition: all 0.3s ease;
  background: transparent;
}

.btn span {
  position: relative;
  z-index: 1;
  font-weight: 500;
}

.btn-bg {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  transition: transform 0.5s ease;
}

.create-btn {
  border: 2px solid #00ff7f;
  color: #00ff7f;
}

.create-btn:hover {
  color: #222222;
}

.create-btn .btn-bg {
  background: #00ff7f;
  transform: translateX(-100%);
}

.create-btn:hover .btn-bg {
  transform: translateX(0);
}

/* 用户头像 */
.user-avatar {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  overflow: hidden;
  border: 2px solid #333333;
  transition: border-color 0.3s ease;
}

.user-avatar:hover {
  border-color: #00a8ff;
}

.user-avatar img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

@media (max-width: 768px) {
  .nav-content {
    padding: 0 1rem;
  }

  .content-wrapper {
    padding: 1rem;
  }

  .blog-card {
    padding: 1.5rem;
  }

  .blog-title {
    font-size: 1.2rem;
  }

  .btn {
    padding: 0.6rem 1.2rem;
    font-size: 0.9rem;
  }
}

.user-menu {
  position: relative;
  cursor: pointer;
}

.user-menu:hover .dropdown-menu {
  display: block;
}

.dropdown-menu {
  display: none;
  position: absolute;
  top: 100%;
  right: 0;
  margin-top: 0.5rem;
  background: #2a2a2a;
  border-radius: 8px;
  padding: 0.5rem;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.2);
}

.logout-btn {
  padding: 0.8rem 1.5rem;
  border: none;
  border-radius: 6px;
  font-size: 1rem;
  cursor: pointer;
  position: relative;
  overflow: hidden;
  transition: all 0.3s ease;
  background: transparent;
  color: #ff4757;
  border: 2px solid #ff4757;
  width: 100%;
  white-space: nowrap;
}

.logout-btn:hover {
  color: #222222;
}

.logout-btn .btn-bg {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: #ff4757;
  transform: translateX(-100%);
  transition: transform 0.5s ease;
}

.logout-btn:hover .btn-bg {
  transform: translateX(0);
}
</style>
