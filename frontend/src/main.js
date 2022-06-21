import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store/index.js';
import axios from 'axios'
import VueAxios from 'vue-axios'

axios.defaults.baseURL = 'http://127.0.0.1:8000/api'
axios.defaults.withCredentials = true

// router.beforeEach((to, from, next) => {
//     // redirect to login page if not logged in and trying to access a restricted page
//
//     let auth = false;
//     console.log("Token", store.state.token);
//     if (store.state.token !== undefined && store.state.token !== "") {
//         auth = true;
//     }
//     console.log("Authentificated", auth);
//
//     if (!auth) {
//         return next('/login');
//     }
//
//     next();
// })

createApp(App).use(router).use(store).use(VueAxios, axios).mount('#app')
