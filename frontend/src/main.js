import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store/index.js';
import axios from 'axios'
import VueAxios from 'vue-axios'
import VueTableLite from 'vue3-table-lite'

axios.defaults.baseURL = 'http://127.0.0.1:8000/api'
axios.defaults.withCredentials = true

const app = createApp(App);
app.use(router);
app.use(store);
app.use(VueAxios, axios);
app.component('VueTable', VueTableLite);
app.mount('#app')
