<script setup>
import { useTodoStore } from '../stores/todo';
import { RouterLink, useRoute } from 'vue-router';
import { onMounted, ref } from 'vue';

const todoStore = useTodoStore();
const route = useRoute();
const isLoaded = ref(false);
const todoId = ref(-1);

// Toast visibility
const showToast = ref(false);

onMounted(async () => {
  await todoStore.loadTodo(route.params.id);
  todoId.value = parseInt(route.params.id);
  isLoaded.value = true;
});

const editTodo = async (selectedTodo) => {
  try {
    const bodyData = {
      title: selectedTodo.title,
      status: selectedTodo.status
    };
    await todoStore.editTodo(bodyData, todoId.value);

    // Show toast
    showToast.value = true;
    setTimeout(() => {
      showToast.value = false;
    }, 2000); // 2 seconds
  } catch (error) {
    console.log('error', error);
  }
};
</script>

<template>
  <div class="p-6 max-w-2xl mx-auto space-y-6 bg-Base 100 min-h-screen">
    <!-- Header -->
    <div class="flex justify-between items-center">
      <h1 class="text-3xl font-bold text-white">🛠️ Edit Todo</h1>
      <RouterLink :to="{ name: 'todo-list' }" class="btn btn-outline btn-accent">
        ← Back
      </RouterLink>
    </div>

    <!-- Toast -->
    <div v-if="showToast" class="toast toast-top toast-end">
      <div class="alert alert-success">
        <span>✅ บันทึกสำเร็จ</span>
      </div>
    </div>

    <div v-if="isLoaded" class="card bg-base-100 shadow-md p-6 space-y-4 border border-gray-600 rounded-lg">
      <!-- ID Display -->
      <!-- <div class="text-gray-400 text-sm">Editing ID: {{ todoId }}</div> -->

      <!-- Title -->
      <div>
        <label class="label text-white font-semibold">Task Name</label>
        <input
          type="text"
          class="input input-bordered w-full"
          v-model="todoStore.selectedTodo.title"
          placeholder="Edit task title"
        />
      </div>

      <!-- Status -->
      <div>
        <label class="label text-white font-semibold">Status</label>
        <select
          class="select select-bordered w-full"
          v-model="todoStore.selectedTodo.status"
        >
          <option disabled>Select status</option>
          <option
            v-for="status in todoStore.statuses"
            :key="status"
            :value="status"
            :style="{
              backgroundColor: 'oklch(25.33% 0.016 252.42)',
              color: 'white'
            }"
          >
            {{ status }}
          </option>
        </select>
      </div>

      <!-- Edit Button -->
      <div class="text-right">
        <button class="btn btn-primary" @click="editTodo(todoStore.selectedTodo)">
          Save Changes
        </button>
      </div>
    </div>

    <div v-else class="text-center text-white">
      <span class="loading loading-spinner loading-lg"></span>
    </div>
  </div>
</template>
