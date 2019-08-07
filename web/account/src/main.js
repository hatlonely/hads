import Vue from 'vue';
import Vuex from 'vuex';
import VueRouter from 'vue-router';

import vuetify from './plugins/vuetify';

import App from './App'
import Account from './pages/Account';
import Login from './pages/Login';
import HHomeBody from "./components/HHomeBody";
import HPersonInfoBody from "./components/HPersonInfoBody";
import HIntroduction from "./components/HIntroduction";

import 'material-design-icons-iconfont/dist/material-design-icons.css';
import 'font-awesome/css/font-awesome.min.css';

Vue.config.productionTip = false
Vue.use(VueRouter);
Vue.use(Vuex);

const router = new VueRouter({
  mode: 'history',
  base: __dirname,
  routes: [
    {
      path: '/', component: Account,
      children: [
        { path: '', component: HHomeBody },
        { path: 'introduction', component: HIntroduction },
        { path: 'home', component: HHomeBody },
        { path: 'person-info', component: HPersonInfoBody }
      ]
    },
    { path: '/login', component: Login },
  ]
});

new Vue({
  vuetify,
  router,
  render: h => h(App)
}).$mount('#app')
