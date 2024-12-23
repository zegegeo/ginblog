<template>
  <div class="register-container">
    <div class="register-section">
      <div class="content">
        <h1>注册账号</h1>
        <h2>开启你的创作之旅</h2>
        <form @submit.prevent="handleRegister" class="register-form">
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
          <div class="input-group">
            <input
              type="password"
              v-model="confirmPassword"
              placeholder="确认密码"
              required
            />
          </div>
          <button type="submit" class="btn register-btn">
            <span>立即注册</span>
            <div class="btn-bg"></div>
          </button>
        </form>
        <p class="login-link">
          已有账号？
          <router-link to="/login" class="link">立即登录</router-link>
        </p>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: "Register",
  data() {
    return {
      username: "",
      password: "",
      confirmPassword: "",
    };
  },
  methods: {
    async handleRegister() {
      if (this.password !== this.confirmPassword) {
        alert("两次输入的密码不一致！");
        return;
      }

      // 组织注册信息
      const registerData = {
        username: this.username,
        password: this.password,
      };

      try {
        // 调用后端注册 API
        const response = await axios.post(
          "http://localhost:8888/api/users/register", // 后端注册接口地址
          registerData
        );

        // 如果注册成功
        if (response.status === 200) {
          alert("注册成功！");
          console.log("后端返回数据：", response.data);
          this.$router.push("/login"); // 跳转到登录页面
        }
      } catch (error) {
        // 处理注册失败
        console.error("注册失败：", error);
        if (error.response && error.response.data) {
          alert(`注册失败：${error.response.data.message}`);
        } else {
          alert("注册失败，请稍后重试！");
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

.register-container {
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

.register-section {
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

.register-section::before {
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

.register-form {
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

.register-btn {
  border: 2px solid #00ff7f;
  color: #00ff7f;
  width: 100%;
}

.register-btn:hover {
  color: #222222;
}

.register-btn .btn-bg {
  background: #00ff7f;
  transform: translateX(-100%);
}

.register-btn:hover .btn-bg {
  transform: translateX(0);
}

.login-link {
  margin-top: 2rem;
  color: #666666;
  font-size: 0.9rem;
}

.link {
  color: #00a8ff;
  text-decoration: none;
  font-weight: 500;
  transition: color 0.3s ease;
}

.link:hover {
  color: #00ff7f;
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

  .register-form {
    gap: 1rem;
  }

  input {
    padding: 0.8rem 1.2rem;
  }
}
</style>