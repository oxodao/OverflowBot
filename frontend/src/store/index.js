import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export const LS_ITEM_NAME = "overflowbot-token";

export default new Vuex.Store({
    state: {
        User: {
            id: null,
            username: null,
            token: null,
        },
        Sounds: [],
        Commands: []
    },
    mutations: {
        setUser: (state, payload) => {
            state.User = payload;
        },
        setSounds: (state, payload) => {
            state.Sounds = payload;
        },
        addSound: (state, sound) => {
            state.Sounds.push(sound);
        },
        editSound: (state, sound) => {
            let j = -1;
            for (let i = 0; i < state.Sounds.length; i++) {
                if (state.Sounds[i].id === sound.id) {
                    j = i;
                    break;
                }
            }

            if (j !== -1) {
                Vue.set(state.Sounds, j, sound);
            }
        },
        deleteSound: (state, id) => {
            state.Sounds = state.Sounds.filter(e => e.id !== id);
        },
        setCommands: (state, payload) => {
            state.Commands = payload;
        },
        addCommand: (state, sound) => {
            state.Commands.push(sound);
        },
        editCommand: (state, sound) => {
            let j = -1;
            for (let i = 0; i < state.Commands.length; i++) {
                if (state.Commands[i].id === sound.id) {
                    j = i;
                    break;
                }
            }

            if (j !== -1) {
                Vue.set(state.Commands, j, sound);
            }
        },
        deleteCommand: (state, id) => {
            state.Commands = state.Commands.filter(e => e.id !== id);
        }
    },
    actions: {
        setUser: ({commit}, user) => {
            localStorage.setItem(LS_ITEM_NAME, user.token);
            commit('setUser', user);
        }
    },
    getters: {
        isLoggedIn: state => (!!state.User && !!state.User.id && state.User.id.length > 0)
    },
    modules: {
    }
})
