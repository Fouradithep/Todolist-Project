import axios from 'axios';

const api = axios.create({
  baseURL: 'https://todolist-backend-18bx.onrender.com', // หรือ URL ของ backend
  withCredentials: true, // ✅ สำคัญมากสำหรับ cookie/session
});

export default api;
