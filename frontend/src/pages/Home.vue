<script setup lang="ts">
    import { onMounted } from "vue";
    import { useRouter } from "vue-router";
    import { useConfigStore, useProfileStore } from "@/modules/stores";

    const router = useRouter();
    const configStore = useConfigStore();
    const profileStore = useProfileStore();

    onMounted(() => {
        if (configStore.getAvailableIDE.length === 0 || configStore.getWorkspace === "" || profileStore.getAllProfiles.length === 0) {
            console.log("[Vaildator] 全新的配置文件，转到 /onboarding 引导用户初始化");
            router.push("/setup/onboarding");
            return;
        }
        
        if (profileStore.getCurrentProfile === null) {
            console.log("[Vaildator] 当前没有选择 Profile，转到 /setup/profile 选择 Profile");
            router.push("/setup/profile");
            return;
        }
    });
</script>

<template></template>
