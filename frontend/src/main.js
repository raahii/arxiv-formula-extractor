import Vue from "vue";
import App from "./App.vue";
import router from "./router";
import VueClipboard from "vue-clipboard2";
import VueMathjax from "vue-mathjax";

Vue.use(VueClipboard);
Vue.use(VueMathjax);

Vue.config.productionTip = false;

new Vue({
  router,
  render: (h) => h(App),
}).$mount("#app");
