<script setup lang="ts">
    import VFade from "@/animation/VFade.vue";
    import Avatar from "./Avatar.vue";
    import { useProfileStore } from "@/modules/stores";
    import { useHandleNavigation } from "@/modules/router";

    const profileStore = useProfileStore();
    const handleNavigation = useHandleNavigation();
</script>

<template>
    <nav class="flex">
        <img src="@/assets/svg/codex-logo.svg" />
        <img src="@/assets/svg/codex-text.svg" class="ml-4" />
        <div role="tablist" class="tabs tabs-border ml-12">
            <a role="tab" class="tab" :class="{ 'tab-active': $route.path === '/' }" @click="handleNavigation('/')">主页</a>
            <a role="tab" class="tab" :class="{ 'tab-active': $route.path.startsWith('/workspace') }" @click="handleNavigation('/workspace', 'startsWith', '/workspace')">工作区</a>
            <a role="tab" class="tab" :class="{ 'tab-active': $route.path.startsWith('/settings') }" @click="handleNavigation('/settings', 'startsWith', '/settings')">设置</a>
        </div>
        <div role="tablist" class="tabs tabs-border absolute right-24">
            <a role="tab" class="tab" :class="{ 'tab-active': $route.path.startsWith('/resources') }" @click="handleNavigation('/resources', 'startsWith', '/resources')">
                实用资源
            </a>
        </div>
        <VFade>
            <Avatar class="absolute right-8 top-8" v-if="profileStore.getCurrentProfile !== null" />
        </VFade>
    </nav>
</template>
