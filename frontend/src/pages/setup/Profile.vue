<script setup lang="ts">
    import { onMounted } from "vue";
    import { useRouter } from "vue-router";

    import GWaveText from "@/animation/GWaveText.vue";
    import VSlideIn from "@/animation/VSlideIn.vue";
    import { useConfigStore, useProfileStore } from "@/modules/stores";

    const router = useRouter();
    const configStore = useConfigStore();
    const profileStore = useProfileStore();

    const handleLogin = (guid: string) => {
        profileStore.selectProfile(guid);
        router.push("/");
    };

    onMounted(() => {
        if (profileStore.current_profile) {
            console.log("[Vaildator] 用户已登录");
            router.push("/");
        }
    });
</script>

<template>
    <main class="mt-4 pb-8 w-full h-full flex flex-col justify-center items-center gap-8">
        <div class="w-full px-32">
            <GWaveText content="选择你的 CodeX 档案" :speed="0.5" start-when-mount class="w-full text-4xl text-center mb-8" />
            <VSlideIn direction="top" :distance="32" :duration="0.5" :delay="1.5">
                <table class="table">
                    <!-- head -->
                    <thead>
                        <tr>
                            <th></th>
                            <th>用户名</th>
                            <th>用户 GUID</th>
                            <th>IDE</th>
                            <th>操作</th>
                        </tr>
                    </thead>
                    <tbody>
                        <!-- row 1 -->
                        <tr v-for="[index, profile] of Object.entries(profileStore.getAllProfiles)" :key="profile.GUID">
                            <th>{{ Number(index) + 1 }}</th>
                            <td>{{ profile.Name }}</td>
                            <td>{{ profile.GUID }}</td>
                            <td>
                                {{ configStore.AvailableIDE.find((i) => i.GUID === profile.IDE)?.Name }} (实例
                                {{ configStore.AvailableIDE.find((i) => i.GUID === profile.IDE)?.GUID.split("-")[0] }})
                            </td>
                            <td>
                                <button class="btn" @click="handleLogin(profile.GUID)">登录</button>
                            </td>
                        </tr>
                    </tbody>
                </table>
            </VSlideIn>
        </div>
    </main>
</template>
