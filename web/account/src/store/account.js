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
    verify({ commit }, { field, email, phone }) {
        if (field == "email") {
            const res = api.verify("email", email)
            if (res.ok) {
                commit("setEmail", email)
            }
            return res
        }
        if (field == "phone") {
            const res = api.verify("phone", phone)
            if (res.ok) {
                commit("setPhone", phone)
            }
            return res
        }
    },
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
    },
    async update({ commit }, { token, field, firstName, lastName, birthday, gender, password, oldPassword, email, phone }) {
        if (field == "name") {
            const res = await api.update(token, { field, firstName, lastName })
            if (res.ok) {
                commit("setFirstName", firstName)
                commit("setLastName", lastName)
            }
            return res
        }
        if (field == "birthday") {
            const res = await api.update(token, { field, birthday })
            if (res.ok) {
                commit("setBirthday", birthday)
            }
            return res
        }
        if (field == "gender") {
            const res = await api.update(token, { field, gender })
            if (res.ok) {
                commit("setGender", gender)
            }
            return res
        }
        if (field == "password") {
            return await api.update(token, { field, password, oldPassword })
        }
        if (field == "email") {
            const res = await api.update(token, { field, email })
            if (res.ok) {
                commit("setEmail", email)
            }
            return res
        }
        if (field == "phone") {
            const res = await api.update(token, { field, phone })
            if (res.ok) {
                commit("setPhone", phone)
            }
            return res
        }
    }
}

export default {
    namespaced: true,
    state,
    actions,
    mutations
}
