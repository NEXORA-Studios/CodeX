<script setup lang="ts">
    import VFade from "@/animation/VFade.vue";
    import Avatar from "./Avatar.vue";
    import { useProfileStore } from "@/modules/stores";

    const profileStore = useProfileStore();
</script>

<template>
    <nav class="flex">
        <img src="@/assets/svg/codex-logo.svg" />
        <img src="@/assets/svg/codex-text.svg" class="ml-4" />
        <VFade>
            <Avatar class="absolute right-8 top-6" v-if="profileStore.getCurrentProfile !== null" />
        </VFade>
    </nav>
</template>
