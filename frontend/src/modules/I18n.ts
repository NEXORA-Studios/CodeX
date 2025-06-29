import { createI18n } from "vue-i18n";
import { zh_CN } from "./locale";

export const i18n = createI18n({
    legacy: false,
    globalInjection: true,
    locale: "zh_CN",
    messages: {
        zh_CN,
    },
});
