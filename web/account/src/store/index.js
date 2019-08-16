import Vue from 'vue'
import Vuex from 'vuex'
import signin from './signin'
import signup from './signup'
import account from './account.js'

Vue.use(Vuex)

const debug = process.env.NODE_ENV !== 'production'

export default new Vuex.Store({
    modules: {
        signin,
        signup,
        account
    },
    strict: debug,
    // plugins: debug ? [createLogger()] : []
})
