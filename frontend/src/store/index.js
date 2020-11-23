import Vue from 'vue'
import Vuex from 'vuex'

import axios from 'axios';

Vue.use(Vuex)

export const LS_ITEM_NAME = "overflowbot-token";

export const setHeaders = (token) => {
    axios.interceptors.request.use(config => {
        config.headers.Authorization = token;
    });
}

export default new Vuex.Store({
    state: {
        User: {
            id: null,
            username: null,
            token: null,
        }
    },
    mutations: {
        setUser: (state, payload) => {
            state.User = payload;
        }
    },
    actions: {
        setUser: ({commit}, user) => {
            localStorage.setItem(LS_ITEM_NAME, user.token);
            setHeaders(user.token);
            commit('setUser', user);
        }
    },
    getters: {
      isLoggedIn: state => (!!state.User && !!state.User.id && state.User.id.length > 0)
    },
    modules: {
    }
})
