import { createApp } from "vue";

import App from "@/App.vue";
import "@/assets/style.css";
import { router } from "@/modules/router";
import { i18n } from "@/modules/I18n";

const app = createApp(App);

app.use(createPinia());
app.use(router);
app.use(i18n);

app.mount("#app");

// Onload Handler
import { useProfileStore } from "@/modules/stores/ProfileStore";
import { GetAllProfile } from "@/modules/wailsjs/go/bind/ProfileBind";
(async function () {
    const profileStore = useProfileStore();
    const allpf = await GetAllProfile();
    profileStore.setProfiles(allpf);
})();

import { useConfigStore } from "@/modules/stores/ConfigStore";
import { GetConfig } from "@/modules/wailsjs/go/bind/ConfigBind";
import { createPinia } from "pinia";
(async function () {
    const configStore = useConfigStore();
    const config = await GetConfig();
    configStore.setConfig(config);
})();

