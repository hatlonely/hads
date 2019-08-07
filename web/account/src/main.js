import Vue from 'vue';
import Account from './pages/Account';
import vuetify from './plugins/vuetify';
import 'material-design-icons-iconfont/dist/material-design-icons.css';
import 'font-awesome/css/font-awesome.min.css';
import Vuex from 'vuex';
import VueRouter from 'vue-router';
import HHomeBody from "./components/HHomeBody";
import HPersonInfoBody from "./components/HPersonInfoBody";

Vue.config.productionTip = false
Vue.use(VueRouter);
Vue.use(Vuex);

const router = new VueRouter({
  mode: 'history',
  base: __dirname,
  routes: [
    { path: '/', component: HHomeBody },
    { path: '/home', component: HHomeBody },
    { path: '/person-info', component: HPersonInfoBody }
  ]
});

new Vue({
  vuetify,
  router,
  render: h => h(Account)
}).$mount('#app')
