<script setup>

import { onMounted, ref, computed } from 'vue';
import { useTodoStore } from '../stores/todo';
import { useAuthStore } from '../stores/auth'
import {RouterLink} from 'vue-router'


const todoStore = useTodoStore();
const auth = useAuthStore();
const todoText = ref('')
const isLoading = ref(false)

onMounted(async() => {
    isLoading.value = true
    // เรียก initializeAuth ก่อนเลย
    await auth.initializeAuth();
    await todoStore.loadTodos()
    isLoading.value = false
})

const addTodo = async(todoText) => {
    try{
        isLoading.value = true
        await todoStore.addTodo(todoText)
        // await todoStore.loadTodos()  //แบบนี้จะเรียก api อีก
        isLoading.value = false
    } catch(error){
        console.log('error',error)
    }
}

const editStatus = async(todoData, todoId) => {
    try{
        isLoading.value = true
        const updateStatus = {
            status: todoData.status
        }
        await todoStore.editTodo(updateStatus, todoId)
        isLoading.value = false
    } catch(error){
        console.log('error',error)
    }
}

const deleteTodo = async(todoId) => {
    try{
        isLoading.value = true
        await todoStore.removeTodo(todoId)
        await todoStore.loadTodos()
        isLoading.value = false
    } catch(error){
        console.log('error',error)
    }
}

const currentTab = ref('Pending')
const statuses = ['All', 'Pending', 'Doing', 'Done']

const filteredTodos = computed(() => {
  if (currentTab.value === 'All') return todoStore.list
  return todoStore.list.filter((todo) => todo.status === currentTab.value)
})




</script>

<template>
    <div class="p-6 max-w-2xl mx-auto space-y-6 min-h-screen" style="background-color: #1D232A;" >

        <!-- Title -->
        <h1 class="text-white text-4xl font-bold drop-shadow-md text-center">
        🚀 My Todo List
        </h1>
        
        <!-- logout -->
        <div class="flex justify-end">
            <button class="btn btn-outline btn-error" @click="auth.logout">Logout</button>
        </div>

        <!-- Add Todo -->
        <div class="flex">
            <input class="input input-bordered w-full" v-model="todoText" placeholder="Add New Task . . .">
            <button class="btn btn-primary ml-2" :disabled="!todoText.trim()" @click="addTodo(todoText)">Add</button>
        </div>

        <!-- Tabs -->
        <div class="flex gap-2 p-2  rounded-xl shadow-inner w-max">
        <a
            v-for="(status, index) in statuses"
            :key="status"
            @click="currentTab = status"
            :class="[
            'transition-all duration-200 ease-in-out font-medium text-sm px-4 py-2 rounded-lg cursor-pointer',
            currentTab === status ? 'text-white shadow-md' : '',
            status === 'Pending' ? (currentTab === status ? 'bg-yellow-500' : 'text-yellow-600 border border-yellow-400') : '',
            status === 'Doing' ? (currentTab === status ? 'bg-blue-500' : 'text-blue-600 border border-blue-400') : '',
            status === 'Done' ? (currentTab === status ? 'bg-green-500' : 'text-green-600 border border-green-400') : '',
            status === 'All' ? (currentTab === status ? 'bg-purple-600' : 'text-purple-800 border border-purple-500') : ''
            ]"
        >
            {{ status }}
        </a>
        </div>

        <!-- Loading State -->
        <div v-if="isLoading" class="text-center text-gray-500">
            <span class="loading loading-spinner loading-lg"></span>
        </div>

        <!-- Todo List -->
        <div v-else class="space-y-4">
            <div
                class="card shadow-md bg-base-100 p-4 border border-gray-500 rounded-lg"   
                v-for="todo in filteredTodos"
                :key="todo.id"
            >
                <div class="flex flex-col md:flex-row md:items-center justify-between space-y-2 md:space-y-0 md:space-x-4">
                    <!-- Title Task-->
                    <div class="flex-1 font-medium text-lg">
                        {{ todo.title }}
                    </div>

                    <!-- Status -->
                    <div class="w-40">
                        <select class="select select-bordered w-full"
                        :class="{
                            'bg-yellow-500 text-white': todo.status === 'Pending',
                            'bg-blue-500 text-white': todo.status === 'Doing',
                            'bg-green-500 text-white': todo.status === 'Done',
                            'bg-white text-black': !['Pending', 'Doing', 'Done'].includes(todo.status)
                            }"
                        v-model="todo.status"
                        @change="editStatus(todo, todo.id)"
                        >
                        <option disabled :style="{
                                backgroundColor: 'oklch(25.33% 0.016 252.42)',
                                color: 'white'
                            }">Select status</option>
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

                    <!-- Action Buttons -->
                    <div class="flex space-x-2">
                        <router-link :to="{ name: 'todo-edit', params: { id: todo.id } }">
                        <button class="btn btn-sm btn-info">Edit</button>
                        </router-link>
                        <button class="btn btn-sm btn-error" @click="deleteTodo(todo.id)">
                        Delete
                        </button>
                    </div>
                </div>    
            </div>
        </div>
    </div>
</template>


