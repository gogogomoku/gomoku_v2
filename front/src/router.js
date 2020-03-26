import Vue from "vue";
import VueRouter from "vue-router";

import App from "./App.vue";

Vue.use(VueRouter);

export const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes: [
    { path: "/", name: "home", component: App },
    { path: "/match/:id", name: "match", component: App }
  ]
});
