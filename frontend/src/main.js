import Vue from 'vue';
import axios from 'axios';
import VueAxios from 'vue-axios';
import App from './App.vue';

Vue.use(VueAxios, axios);

//   {
//     path: '',
//     component: HelloWorld,
//   },
// ];
// const routes = [
//     path: '',
//     component: MainPage,
//   },
//   {
//     path: '/login',
//     component: LoginUser,
//   },
// ];

// const router = new VueRouter({
//   mode: 'history',
// });

Vue.config.productionTip = false;

new Vue({
  render: (h) => h(App),
}).$mount('#app');
