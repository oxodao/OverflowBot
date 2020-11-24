import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'

import {LS_ITEM_NAME} from './store';
import axios from "axios";

Vue.config.productionTip = false

// Meh, ugly AF https://stackoverflow.com/questions/53703581/how-to-get-something-to-run-before-vue-router
const token = localStorage.getItem(LS_ITEM_NAME);
if (token !== null) {
  axios.get("/api/auth/validate", { headers: { "Authorization": token } })
      .then((resp) => {
          new Vue({store}).$store.commit('setUser', resp.data);
          InitializeVue();
      })
      .catch(() => {
        localStorage.removeItem(LS_ITEM_NAME);
        InitializeVue();
      })
} else {
    InitializeVue();
}

function InitializeVue() {
  window.app = new Vue({
    router,
    store,
    render: h => h(App)
  }).$mount('#app')
}

