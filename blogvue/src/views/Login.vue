<template>
  <div class="login-container">
    <div class="login-section">
      <div class="content">
        <h1>欢迎回来</h1>
        <h2>登录您的账号</h2>
        <form @submit.prevent="handleLogin" class="login-form">
          <div class="input-group">
            <input
              type="text"
              v-model="username"
              placeholder="用户名"
              required
            />
          </div>
          <div class="input-group">
            <input
              type="password"
              v-model="password"
              placeholder="密码"
              required
            />
          </div>
          <button type="submit" class="btn login-btn">
            <span>立即登录</span>
            <div class="btn-bg"></div>
          </button>
        </form>
        <p class="register-link">
          还没有账号？
          <router-link to="/register" class="link">立即注册</router-link>
        </p>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: "Login",
  data() {
    return {
      username: "",
      password: "",
    };
  },
  methods: {
    async handleLogin() {
      try {
        const response = await axios.post('http://localhost:8888/api/users/login', {
          username: this.username,
          password: this.password
        });

        if (response.status === 200) {
          // 登录成功
          const token = response.data.token; // 假设后端返回了token
          // 可以将token存储到localStorage或vuex中
          localStorage.setItem('token', token);
          
          // 提示登录成功
          alert('登录成功！');
          
          // 跳转到博客页面
          this.$router.push('/blog');
        }
      } catch (error) {
        console.error('登录失败：', error);
        if (error.response && error.response.data) {
          alert(`登录失败：${error.response.data.message}`);
        } else {
          alert('登录失败，请稍后重试！');
        }
      }
    },
  },
};
</script>

<style scoped>
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

.login-container {
  min-height: 100vh;
  width: 100vw;
  display: flex;
  align-items: center;
  justify-content: center;
  background: #1a1a1a;
  margin: 0;
  padding: 0;
  overflow-x: hidden;
  position: fixed;
  top: 0;
  left: 0;
}

.login-section {
  background: #222222;
  width: 100vw;
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  position: relative;
  padding: 0;
  margin: 0;
}

.login-section::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 4px;
  background: linear-gradient(90deg, #00a8ff, #00ff7f);
}

.content {
  text-align: center;
  width: 100%;
  max-width: 400px;
  padding: 0 20px;
}

h1 {
  color: #ffffff;
  margin-bottom: 1rem;
  font-size: 3rem;
  font-weight: 700;
  letter-spacing: -1px;
  text-shadow: 0 0 10px rgba(0, 168, 255, 0.3);
}

h2 {
  color: #888888;
  font-size: 1.5rem;
  font-weight: 400;
  margin-bottom: 3rem;
}

.login-form {
  display: flex;
  flex-direction: column;
  gap: 1.5rem;
}

.input-group {
  position: relative;
  width: 100%;
}

input {
  width: 100%;
  padding: 1rem 1.5rem;
  background: #2a2a2a;
  border: 2px solid #333333;
  border-radius: 8px;
  color: #ffffff;
  font-size: 1rem;
  transition: all 0.3s ease;
}

input:focus {
  outline: none;
  border-color: #00a8ff;
  box-shadow: 0 0 10px rgba(0, 168, 255, 0.2);
}

input::placeholder {
  color: #666666;
}

.btn {
  padding: 1rem;
  border: none;
  border-radius: 8px;
  font-size: 1.1rem;
  cursor: pointer;
  position: relative;
  overflow: hidden;
  transition: all 0.3s ease;
  background: transparent;
  margin-top: 1rem;
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

.login-btn {
  border: 2px solid #00a8ff;
  color: #00a8ff;
  width: 100%;
}

.login-btn:hover {
  color: #222222;
}

.login-btn .btn-bg {
  background: #00a8ff;
  transform: translateX(-100%);
}

.login-btn:hover .btn-bg {
  transform: translateX(0);
}

.register-link {
  margin-top: 2rem;
  color: #666666;
  font-size: 0.9rem;
}

.link {
  color: #00ff7f;
  text-decoration: none;
  font-weight: 500;
  transition: color 0.3s ease;
}

.link:hover {
  color: #00a8ff;
}

@media (max-width: 768px) {
  .content {
    padding: 0 1.5rem;
  }

  h1 {
    font-size: 2.5rem;
  }

  h2 {
    font-size: 1.2rem;
    margin-bottom: 2rem;
  }

  .login-form {
    gap: 1rem;
  }

  input {
    padding: 0.8rem 1.2rem;
  }
}
</style>