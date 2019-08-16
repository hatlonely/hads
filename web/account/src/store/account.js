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

const actions = {
    signin({ state }) {
        return api.siginin(state.username, state.password)
    }
}

export default {
    namespaced: true,
    state,
    actions
}
