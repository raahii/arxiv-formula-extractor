import Vue from "vue";
import VueRouter from "vue-router";
import TopPage from "@/views/TopView";

Vue.use(VueRouter);

const routes = [
  {
    path: "/",
    name: "TopPage",
    component: TopPage,
  },
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes,
});

export default router;
