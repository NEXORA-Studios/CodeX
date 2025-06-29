<script setup lang="ts">
    import { computed, onMounted } from "vue";
    import { useRouter } from "vue-router";
    import { useConfigStore, useProfileStore } from "@/modules/stores";
    import GWaveText from "@/animation/GWaveText.vue";

    const router = useRouter();
    const configStore = useConfigStore();
    const profileStore = useProfileStore();

    const dUsername = computed(() => {
        if (profileStore.getCurrentProfile) {
            return profileStore.getCurrentProfile.Name;
        } else {
            return `{Unknown}`;
        }
    });

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

<template>
    <main class="mt-4 p-8 w-full h-full flex gap-8">
        <GWaveText :content="`欢迎回来，${dUsername}`" :speed="0.5" :distance="32" start-when-mount />
    </main>
</template>
