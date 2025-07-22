import { defineStore } from 'pinia';
import api from '../utils/axios';
// import router from '../router';

export const useTodoStore = defineStore('todo', {
  state: () => ({
    list: [],
    selectedTodo: {},
    statuses: ['Pending', 'Doing', 'Done']
  }),

  actions: {
    async loadTodos() {
      try {
        const response = await api.get('/tasks');
        this.list = response.data;
        console.log('load todo list complete')
      } catch (error) {
        console.log('error',error)
      }
    },

    async loadTodo(id) {
      try {
        const response = await api.get(`/task/${id}`);
        this.selectedTodo = response.data;
        console.log('load todo by id complete')
      } catch (error) {
        console.log('error',error)
      }
    },

    async addTodo(todoText) {
        const bodyData ={
            title: todoText,
        }
      try {
        const response = await api.post(`/task`,bodyData);
        console.log(response.data)
        this.list.push(response.data)
        console.log('Add Todo Successful');
      } catch (error) {
        console.log('error',error)
      }
    },

    async editTodo(todoData, id){
        /*
        {
        title:'Do something'
        status:'Doing'
        }
        */
       try{
        const response = await api.put(`/task/${id}`,todoData);
        console.log('Edit Todo Successful');
       }catch (error) {
        console.log('error',error)
      }
    },

    async removeTodo(id){
       try{
        const response = await api.delete(`/task/${id}`);
        console.log('Delete Todo Successful');
       }catch (error) {
        console.log('error',error)
      }
    },

  }
});
