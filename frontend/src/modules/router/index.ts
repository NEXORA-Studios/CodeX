import { createRouter, createWebHistory } from "vue-router";

export const router = createRouter({
    history: createWebHistory(),
    routes: [
        // 初始化
        {
            path: "/setup/onboarding",
            name: "Onboarding",
            component: () => import("@/pages/setup/Onboarding.vue"),
        },
        {
            path: "/setup/profile",
            name: "Profile",
            component: () => import("@/pages/setup/Profile.vue"),
        },
        // 主页
        {
            path: "/",
            name: "Home",
            component: () => import("@/pages/Index.vue"),
        },
        // 设置
        {
            path: "/settings",
            name: "Settings",
            component: () => import("@/pages/settings/Index.vue"),
        },
        {
            path: "/settings/ide",
            name: "SettingsIDE",
            component: () => import("@/pages/settings/IDE.vue"),
        },
        {
            path: "/settings/environment",
            name: "SettingsEnvironment",
            component: () => import("@/pages/settings/Environment.vue"),
        },
        // 实用资源
        {
            path: "/resources",
            name: "Resources",
            component: () => import("@/pages/Res.vue"),
        },
    ],
});

export { useHandleNavigation } from "./utils";
