import api from '../api/api'

const state = {
    firstName: '',
    lastName: '',
    email: '',
    password: '',
    phone: '',
    birthday: '',
    gender: '',
    isSignedIn: false
}

const mutations = {
    setFirstName(state, value) {
        state.firstName = value
    },
    setLastName(state, value) {
        state.lastName = value
    },
    setEmail(state, value) {
        state.email = value
    },
    setPhone(state, value) {
        state.phone = value
    },
    setPassword(state, value) {
        state.password = value
    },
    setCode(state, value) {
        state.code = value
    },
    setBirthday(state, value) {
        state.birthday = value
    },
    setGender(state, value) {
        state.gender = value
    },
    setIsSignedIn(state, value) {
        state.isSignedIn = value
    }
}

const actions = {
    async getAccount({ commit }, token) {
        const res = await api.getAccount(token)
        if (res.ok) {
            const account = res.account
            commit("setFirstName", account.firstName)
            commit("setLastName", account.lastName)
            commit("setEmail", account.email)
            commit("setPhone", account.phone)
            commit("setBirthday", account.birthday)
            commit("setGender", account.gender)
            commit("setIsSignedIn", true)
        }
    }
}

export default {
    namespaced: true,
    state,
    actions,
    mutations
}
