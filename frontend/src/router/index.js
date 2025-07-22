import { createRouter, createWebHistory } from 'vue-router';

import EditView from '../views/EditView.vue';
import HomeView from '../views/HomeView.vue';
import LoginView from '../views/LoginView.vue';
import RegisterView from '../views/RegisterView.vue';


const routes = [
  { path: '/', redirect: '/login' },
  { path: '/login', name: 'todo-login',component: LoginView },
  { path: '/register', name: 'todo-register',component: RegisterView },
  { path: '/home', name: 'todo-list',component: HomeView },
  { path: '/edit/:id', name: 'todo-edit',component: EditView },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
