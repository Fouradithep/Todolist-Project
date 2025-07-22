<script setup>
import { ref } from 'vue';
import { useAuthStore } from '../stores/auth';
import { RouterLink } from 'vue-router';

const email = ref('');
const password = ref('');
const auth = useAuthStore();

const showError = ref(false);

const handleLogin = async () => {
  try {
    await auth.login(email.value, password.value);
    // ถ้าสำเร็จ อาจ redirect หรือทำอย่างอื่นเพิ่มได้
  } catch (err) {
    showError.value = true;
    setTimeout(() => {
      showError.value = false;
    }, 2500);
  }
};
</script>

<template>
  <div class="p-6 max-w-md mx-auto space-y-6">
    <h1 class="text-3xl font-bold text-center text-white mb-6">🔑 Login</h1>

    <form @submit.prevent="handleLogin" class="space-y-4">
      <input
        type="email"
        v-model="email"
        placeholder="Email"
        class="input input-bordered w-full"
        required
      />
      <input
        type="password"
        v-model="password"
        placeholder="Password"
        class="input input-bordered w-full"
        required
      />
      <button type="submit" class="btn btn-primary w-full">Login</button>
    </form>

    <div class="text-center mt-4">
      <RouterLink :to="{ name: 'todo-register' }" class="text-primary hover:underline">
        Register
      </RouterLink>
    </div>

    <!-- Toast error message -->
    <div
      v-if="showError"
      class="toast toast-top toast-center fixed z-50"
      style="top: 20px;"
    >
      <div class="alert alert-error shadow-lg">
        <span>❌ Login failed. Please check your credentials.</span>
      </div>
    </div>
  </div>
</template>
