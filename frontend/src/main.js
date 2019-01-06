// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import VueClipboard from 'vue-clipboard2'
import VueMathjax from 'vue-mathjax'

Vue.use(VueClipboard)
Vue.use(VueMathjax)

Vue.config.productionTip = false

// let axiosInstance = axios.create({
//   baseURL: 'https://localhost:13',
//   #<{(| other custom settings |)}>#
// });
//
// module.exports = axiosInstance;

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  components: { App },
  template: '<App/>'
})
