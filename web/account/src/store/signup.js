import api from '../api/api'

const state = {
    firstName: '',
    lastName: '',
    email: '',
    password: '',
    phone: '',
    code: '',
    birthday: '',
    gender: '',
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
}

const actions = {
    verify({ state }, type) {
        if (type == "email") {
            return api.verify("email", state.email)
        }
        if (type == "phone") {
            return api.verify("phone", state.phone)
        }
    },
    genAuthCode({ state }, type) {
        if (type == "email") {
            return api.genAuthCode(type, state.email, state.firstName, state.lastName)
        }
        if (type == "phone") {
            return api.genAuthCode(type, state.phone, state.firstName, state.lastName)
        }
    },
    verifyAuthCode({ state }, type) {
        if (type == "email") {
            return api.verifyAuthCode(type, state.email, state.code)
        }
        if (type == "phone") {
            return api.verifyAuthCode(type, state.phone, state.code)
        }
    },
    signUp({ state }) {
        return api.signUp(state.email, state.phone, state.password, state.firstName, state.lastName, state.birthday, state.gender, state.code)
    }
}

export default {
    namespaced: true,
    state,
    actions,
    mutations,
}
