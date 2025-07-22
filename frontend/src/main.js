import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import { createPinia } from 'pinia'   // เพิ่มตรงนี้

import './style.css'

const app = createApp(App)
const pinia = createPinia()

app.use(pinia)   // ลงทะเบียน Pinia ก่อน router
app.use(router)


app.mount('#app')
