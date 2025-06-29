import { createRouter, createWebHistory } from "vue-router";

export const router = createRouter({
    history: createWebHistory(),
    routes: [
        {
            path: "/",
            name: "Home",
            component: () => import("@/pages/Home.vue"),
        },
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
    ],
});
