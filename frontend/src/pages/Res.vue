<script setup lang="ts">
    import GWaveText from "@/animation/GWaveText.vue";
</script>

<template>
    <main class="mt-4 p-4 pb-8 w-full h-full flex flex-col">
        <section class="flex items-center">
            <GWaveText class="text-xl ml-4" content="实用资源推荐" :speed="0.3" start-when-mount />
        </section>
    </main>
</template>
