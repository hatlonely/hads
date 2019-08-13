import Vue from 'vue';
import Vuex from 'vuex';
import VueRouter from 'vue-router';
import VueConfig from 'vue-config';
import VueCookies from 'vue-cookies'

import vuetify from './plugins/vuetify';
import config from "./assets/js/config"

import App from './App'
import Account from './pages/Account';
import SignIn from './pages/SignIn';
import SignUp from './pages/SignUp';
import Username from './components/signin/Username';
import Password from './components/signin/Password';
import Sorry from './components/signin/Sorry';
import SignUpSorry from './components/signup/Sorry';
import CreateAccount from './components/signup/CreateAccount';
import VerifyPhone from './components/signup/VerifyPhone';
import VerifyCode from './components/signup/VerifyCode';
import PersonalDetail from './components/signup/PersonalDetail';
import Privacy from "./components/signup/Privacy";
import HHomeBody from "./components/account/HHomeBody";
import HPersonInfoBody from "./components/account/HPersonInfoBody";
import HIntroduction from "./components/account/HIntroduction";

import 'material-design-icons-iconfont/dist/material-design-icons.css';
import 'font-awesome/css/font-awesome.min.css';

Vue.config.productionTip = false;
Vue.use(VueRouter);
Vue.use(Vuex);
Vue.use(VueCookies);
Vue.use(VueConfig, config);

const store = new Vuex.Store({
  state: {
    signup: {
      firstName: "",
      lastName: "",
      email: "",
      password: "",
      phone: "",
      code: "",
      birthday: "",
      gender: "",
    },
    signin: {
      username: "",
      password: ""
    },
    account: {
      firstName: "",
      lastName: "",
      email: "",
      password: "",
      phone: "",
      birthday: "",
      gender: "",
    },
  },
});

const router = new VueRouter({
  mode: 'history',
  base: __dirname,
  routes: [
    {
      path: '/account', component: Account,
      children: [
        { path: '', component: HHomeBody },
        { path: 'introduction', component: HIntroduction },
        { path: 'home', component: HHomeBody },
        { path: 'personinfo', component: HPersonInfoBody }
      ]
    },
    {
      path: '/signin', component: SignIn, children: [
        { path: '', component: Username },
        { path: 'password', component: Password },
        { path: 'sorry', component: Sorry },
      ]
    },
    {
      path: '/signup', component: SignUp, children: [
        { path: '', component: CreateAccount },
        { path: 'verifyphone', component: VerifyPhone },
        { path: 'verifycode', component: VerifyCode },
        { path: 'personaldetail', component: PersonalDetail },
        { path: 'privacy', component: Privacy },
        { path: 'sorry', component: SignUpSorry },
      ]
    },
  ]
});

new Vue({
  el: "#app",
  vuetify,
  router,
  store,
  render: h => h(App)
}).$mount('#app')
