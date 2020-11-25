import Vue from 'vue'
import VueRouter from 'vue-router'

import store from '@/store';
import Home from '@/views/Home.vue'
import Login from "@/views/Login";
import UploadCours from "@/views/UploadCours";
import Soundboard from "@/views/Sounds";
import CustomCommands from "@/views/CustomCommands";

Vue.use(VueRouter)

const isAuthenticated = (to, from, next) => {
  if (store.getters.isLoggedIn) {
    next();
    return;
  }

  next('/login')
}

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home,
    beforeEnter: isAuthenticated
  },
  {
    path: '/cours',
    name: 'Cours',
    component: UploadCours,
    beforeEnter: isAuthenticated
  },
  {
    path: '/sound',
    name: 'Soundboard',
    component: Soundboard,
    beforeEnter: isAuthenticated
  },
  {
    path: '/commands',
    name: 'Commandes',
    component: CustomCommands,
    beforeEnter: isAuthenticated
  },
  {
    path: '/login/:id?/:username?/:token?',
    component: Login
  }
]

const router = new VueRouter({
  routes
})

export default router
