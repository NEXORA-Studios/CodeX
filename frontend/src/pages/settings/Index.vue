<script setup lang="ts">
    import GWaveText from "@/animation/GWaveText.vue";
    import VSlideIn from "@/animation/VSlideIn.vue";

    const settings_list = [
        {
            title: "IDE 设置",
            desc: "设置可供使用的 IDE",
            uri: "/settings/ide",
        },
        {
            title: "检查开发环境",
            desc: "检查本地开发环境",
            uri: "/settings/environment",
        },
    ];
</script>

<template>
    <main class="mt-4 p-4 pb-8 w-full h-full flex flex-col">
        <GWaveText class="text-3xl mb-8" content="设置" :speed="0.3" start-when-mount />
        <VSlideIn
            class="mb-2 last:mb-0"
            v-for="[index, setting] in Object.entries(settings_list)"
            direction="top"
            :duration="0.3"
            :distance="16"
            :delay="Number(index) * 0.15 + 0.3">
            <div class="w-full card bg-base-100 shadow-sm">
                <div class="card-body py-4 flex-row items-center">
                    <span class="text-lg font-bold">{{ setting.title }}</span>
                    <span class="text-sm text-neutral-content/50">{{ setting.desc }}</span>
                    <button class="ml-auto btn btn-sm btn-primary" @click="$router.push(setting.uri)">进入</button>
                </div>
            </div>
        </VSlideIn>
    </main>
</template>
