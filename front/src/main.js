import Vue from "vue";
import App from "./App.vue";
import store from "./store/store";
import "./assets/tailwind.css";
import { library } from "@fortawesome/fontawesome-svg-core";
import { faCrown, faRobot, faUser } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";

library.add(faRobot, faUser, faCrown);

Vue.component("font-awesome-icon", FontAwesomeIcon);

Vue.config.productionTip = false;

new Vue({
  store,
  render: h => h(App)
}).$mount("#app");
