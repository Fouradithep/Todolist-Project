import { defineStore } from 'pinia';
import api from '../utils/axios';
import router from '../router';

export const useAuthStore = defineStore('auth', {
  state: () => ({
    user: null,
    token: null, // เพิ่ม token state
  }),

  actions: {
    async login(email, password) {
      try {
        const response = await api.post('/login', { email, password });
        
        // เก็บ token สำหรับ iOS (fallback)
        if (response.data.token) {
          this.token = response.data.token;
          localStorage.setItem('jwt', response.data.token);
        }
        
        this.user = response.data.user;
        router.push('/home');
      } catch (err) {
        console.error('Login failed', err.response?.data);
        throw err;
      }
    },

    async register(email, password) {
      try {
        const response = await api.post('/register', { email, password });
        
        // ถ้า backend ส่ง token มาหลัง register (auto login)
        if (response.data.token) {
          this.token = response.data.token;
          localStorage.setItem('jwt', response.data.token);
        }
        
        this.user = response.data.user || null;
        return response;
      } catch (err) {
        console.error('Register failed', err.response?.data);
        throw err;
      }
    },

    // ถ้าต้องการ fetch user profile ในอนาคต
    // async fetchUser() {
    //   try {
    //     const response = await api.get('/me');
    //     this.user = response.data.user;
    //   } catch (err) {
    //     console.error('Fetch user failed', err);
    //     this.user = null;
    //     this.token = null;
    //     localStorage.removeItem('jwt');
    //   }
    // },

    async logout() {
      try {
        await api.post('/logout');
      } catch (err) {
        console.error('Logout API failed', err);
      } finally {
        // ล้างข้อมูลทั้งหมด
        this.user = null;
        this.token = null;
        localStorage.removeItem('jwt');
        router.push('/login');
      }
    },

    // เพิ่ม method สำหรับ initialize app (check existing token)
    async initializeAuth() {
      const storedToken = localStorage.getItem('jwt');
      if (storedToken) {
        this.token = storedToken;
        // ถ้ามี backend endpoint /me ในอนาคต ก็เอา comment ออกได้
        // await this.fetchUser();
      }
    },

    // // Helper method สำหรับตรวจสอบว่า login อยู่หรือไม่
    // isAuthenticated() {
    //   return !!(this.user || this.token || localStorage.getItem('jwt'));
    // }
  },
});