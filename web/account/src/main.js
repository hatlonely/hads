import Vue from 'vue';
import Vuex from 'vuex';
import VueRouter from 'vue-router';
import VueConfig from 'vue-config';
import VueCookies from 'vue-cookies'

import vuetify from './plugins/vuetify';
import config from "./assets/js/config"

import App from './App'

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
      isSignedIn: false
    },
  },
});

const router = new VueRouter({
  mode: 'history',
  base: __dirname,
  routes: [
    {
      path: '/account', component: () => import("./pages/Account"),
      children: [
        { path: '', component: () => import("./components/account/HHomeBody") },
        { path: 'introduction', component: () => import("./components/account/HIntroduction") },
        { path: 'home', component: () => import("./components/account/HHomeBody") },
        { path: 'personinfo', component: () => import("./components/account/HPersonInfoBody") }
      ]
    },
    {
      path: '/signin', component: () => import("./pages/SignIn"),
      children: [
        { path: '', component: () => import("./components/signin/Username") },
        { path: 'password', component: () => import("./components/signin/Password") },
        { path: 'sorry', component: () => import("./components/signin/Sorry") },
      ]
    },
    {
      path: '/signup', component: () => import("./pages/SignUp"),
      children: [
        { path: '', component: () => import("./components/signup/CreateAccount") },
        { path: 'verifyphone', component: () => import("./components/signup/VerifyPhone") },
        { path: 'verifycode', component: () => import("./components/signup/VerifyCode") },
        { path: 'personaldetail', component: () => import("./components/signup/PersonalDetail") },
        { path: 'privacy', component: () => import("./components/signup/Privacy") },
        { path: 'sorry', component: () => import("./components/signup/Sorry") },
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
