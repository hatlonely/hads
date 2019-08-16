import Vue from 'vue';
import Vuex from 'vuex';
import VueRouter from 'vue-router';
import VueConfig from 'vue-config';
import VueCookies from 'vue-cookies';

import vuetify from './plugins/vuetify';
import config from './assets/js/config';
import store from './store';

import App from './App';

import 'material-design-icons-iconfont/dist/material-design-icons.css';
import 'font-awesome/css/font-awesome.min.css';

Vue.config.productionTip = false;
Vue.use(VueRouter);
Vue.use(Vuex);
Vue.use(VueCookies);
Vue.use(VueConfig, config);

const router = new VueRouter({
  mode: 'history',
  base: __dirname,
  routes: [
    {
      path: '/account', component: () => import('./pages/Account'),
      children: [
        { path: '', component: () => import('./components/account/HHomeBody') },
        { path: 'introduction', component: () => import('./components/account/HIntroduction') },
        { path: 'home', component: () => import('./components/account/HHomeBody') },
        { path: 'personinfo', component: () => import('./components/account/HPersonInfoBody') },
        { path: 'update/name', component: () => import('./components/account/update/Name') },
        { path: 'update/birthday', component: () => import('./components/account/update/Birthday') },
        { path: 'update/password', component: () => import('./components/account/update/Password') },
        { path: 'update/sorry', component: () => import('./components/account/update/Sorry') },
        { path: 'update/gender', component: () => import('./components/account/update/Gender') },
        { path: 'update/phone', component: () => import('./components/account/update/Phone') },
        { path: 'update/email', component: () => import('./components/account/update/Email') },
      ]
    },
    {
      path: '/signin', component: () => import('./pages/SignIn'),
      children: [
        { path: '', component: () => import('./components/signin/Username') },
        { path: 'password', component: () => import('./components/signin/Password') },
        { path: 'sorry', component: () => import('./components/signin/Sorry') },
      ]
    },
    {
      path: '/signup', component: () => import('./pages/SignUp'),
      children: [
        { path: '', component: () => import('./components/signup/CreateAccount') },
        { path: 'phone', component: () => import('./components/signup/Phone') },
        { path: 'verifyphone', component: () => import('./components/signup/VerifyPhone') },
        { path: 'verifyemail', component: () => import('./components/signup/VerifyEmail') },
        { path: 'personaldetail', component: () => import('./components/signup/PersonalDetail') },
        { path: 'privacy', component: () => import('./components/signup/Privacy') },
        { path: 'sorry', component: () => import('./components/signup/Sorry') },
      ]
    },
  ]
});

new Vue({
  vuetify,
  router,
  store,
  render: h => h(App)
}).$mount('#app')
