import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store/index.js';
import axios from 'axios'
import VueAxios from 'vue-axios'

axios.defaults.baseURL = 'http://localhost:8000/api'
axios.defaults.withCredentials

createApp(App).use(router).use(store).use(VueAxios, axios).mount('#app')
