<script setup>
import { ref } from 'vue';
import { useAuthStore } from '../stores/auth';
import { useRouter } from 'vue-router';

const email = ref('');
const password = ref('');
const confirmPassword = ref('');
const auth = useAuthStore();
const router = useRouter();

const showError = ref('');
const showSuccess = ref(false);

const handleRegister = async () => {
  if (password.value !== confirmPassword.value) {
    showError.value = 'Passwords do not match';
    setTimeout(() => (showError.value = ''), 2500);
    return;
  }

  try {
    await auth.register(email.value, password.value);
    showSuccess.value = true;
    setTimeout(() => {
      showSuccess.value = false;
      router.push({ name: 'todo-login' });
    }, 2000);
  } catch (err) {
    showError.value = 'Registration failed';
    setTimeout(() => (showError.value = ''), 2500);
  }
};
</script>

<template>
  <div class="p-6 max-w-2xl mx-auto space-y-6 bg-Base 100 min-h-screen" style="background-color: #1D232A;">
    <h1 class="text-3xl font-bold text-center text-white mb-6">📝 Register</h1>

    <form @submit.prevent="handleRegister" class="space-y-4">
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
      <input
        type="password"
        v-model="confirmPassword"
        placeholder="Confirm Password"
        class="input input-bordered w-full"
        required
      />
      <button type="submit" class="btn btn-primary w-full">Register</button>
    </form>

    <div class="text-center mt-4">
      <RouterLink :to="{ name: 'todo-login' }" class="text-primary hover:underline">
        Already have an account? Login
      </RouterLink>
    </div>

    <!-- Toast error -->
    <div v-if="showError" class="toast toast-top toast-center fixed z-50" style="top: 20px;">
      <div class="alert alert-error shadow-lg">
        <span>❌ {{ showError }}</span>
      </div>
    </div>

    <!-- Toast success -->
    <div v-if="showSuccess" class="toast toast-top toast-center fixed z-50" style="top: 20px;">
      <div class="alert alert-success shadow-lg">
        <span>✅ Registration successful! Redirecting...</span>
      </div>
    </div>
  </div>
</template>
