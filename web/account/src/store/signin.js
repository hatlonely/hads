import api from '../api/api'

const state = {
    username: "",
    password: ""
}

const mutations = {
    setUsername(state, value) {
        state.username = value
    },
    setPassword(state, value) {
        state.password = value
    }
}

const actions = {
    signIn({ state }) {
        return api.signIn(state.username, state.password)
    },
    verify({ state }, type) {
        if (type == "username") {
            return api.verify("username", state.username)
        }
    },
}

export default {
    namespaced: true,
    state,
    actions,
    mutations,
}
