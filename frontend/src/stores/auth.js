import { defineStore } from 'pinia';
import api from '../utils/axios';
import router from '../router';

export const useAuthStore = defineStore('auth', {
  state: () => ({
    user: null,
  }),

  actions: {
    async login(email, password) {
      try {
        const response = await api.post('/login', { email, password });
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
        // สมมติ backend ไม่ส่ง user มา แต่ส่งแค่ message
        // หรือถ้า backend ส่ง user มา ก็เก็บได้เหมือน login
        this.user = response.data.user || null; 
        // ถ้าอยากให้ login อัตโนมัติหลังสมัครสำเร็จ อาจ redirect ไป /home ได้เลย
        // router.push('/home');
        return response;
      } catch (err) {
        console.error('Register failed', err.response?.data);
        throw err;
      }
    },

    // async fetchUser() {
    //   try {
    //     const response = await api.get('/me');
    //     this.user = response.data.user;
    //   } catch {
    //     this.user = null;
    //   }
    // },

    async logout() {
      await api.post('/logout');
      this.user = null;
      router.push('/login');
    },
  },
});
