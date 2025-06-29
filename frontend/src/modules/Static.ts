import { defineAsyncComponent, type Component } from "vue";

export const dIdeLogo: { [key: string]: Component } = {
    "Visual Studio Code": defineAsyncComponent(() => import("@/icons/VisualStudioCode.vue")),
    Trae: defineAsyncComponent(() => import("@/icons/Trae.vue")),
    "Trae CN": defineAsyncComponent(() => import("@/icons/Trae.vue")),
};

export const dEnvLogo: { [key: string]: Component } = {
    "nodejs": defineAsyncComponent(() => import("@/icons/NodeJS.vue")),
    "npm": defineAsyncComponent(() => import("@/icons/NPM.vue")),
    "pnpm": defineAsyncComponent(() => import("@/icons/PNPM.vue")),
    "yarn": defineAsyncComponent(() => import("@/icons/Yarn.vue")),
    "go": defineAsyncComponent(() => import("@/icons/Go.vue")),
    "python": defineAsyncComponent(() => import("@/icons/Python.vue")),
    "rustup": defineAsyncComponent(() => import("@/icons/Rust.vue")),
    "rustc": defineAsyncComponent(() => import("@/icons/Rust.vue")),
    "cargo": defineAsyncComponent(() => import("@/icons/Rust.vue")),
}
