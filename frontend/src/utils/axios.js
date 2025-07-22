import axios from 'axios';

const api = axios.create({
  baseURL: 'https://todo-backend-ly2q.onrender.com', // หรือ URL ของ backend
  withCredentials: true, // ✅ สำคัญมากสำหรับ cookie/session
});

export default api;
