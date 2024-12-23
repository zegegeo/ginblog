import { createRouter, createWebHistory } from "vue-router";
import Home from "../views/Home.vue";
import Login from "../views/Login.vue";
import Register from "../views/Register.vue";
import About from '../views/About.vue';
import Blog from '../views/Blog.vue';

const routes = [
{ 
    path: "/", 
    name: "Home", 
    component: Home 
},
{ 
    path: "/login", 
    name: "Login", 
    component: Login 
},
{   
    path: "/register", 
    name: "Register", 
    component: Register 
},
{
    path: '/about',
    name: 'About',
    component: About
},
{
    path: '/blog',
    name: 'Blog',
    component: Blog,
    meta: {
        requiresAuth: true
    }
}
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

router.beforeEach((to, from, next) => {
  if (to.meta.requiresAuth) {
    const token = localStorage.getItem('token');
    if (!token) {
      next('/login');
    } else {
      next();
    }
  } else {
    next();
  }
});

export default router;