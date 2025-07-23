import axios from 'axios';

const api = axios.create({
  baseURL: 'https://todolist-backend-18bx.onrender.com', // หรือ URL ของ backend
  withCredentials: true, // ✅ สำคัญมากสำหรับ cookie/session
});

// เพิ่ม interceptor ก่อน export
api.interceptors.request.use(config => {
  const token = localStorage.getItem('jwt');
  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config;
});

export default api;
