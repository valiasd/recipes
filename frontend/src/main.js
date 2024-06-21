import {createApp} from 'vue'
import { router } from "./router/mainRouter";
import App from "./App.vue";

createApp(App)
    .use(router)
    .mount('#app')

