import { createApp } from "vue";
import "./style.css";
import App from "./App.vue";
import router from "./router";

// primevue
import PrimeVue from "primevue/config";
import "primevue/resources/themes/aura-light-purple/theme.css";

const app = createApp(App);

app.use(PrimeVue);
app.use(router);

app.mount("#app");
