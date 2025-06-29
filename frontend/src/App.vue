<script lang="ts" setup>
    import VFade from "./animation/VFade.vue";
    import DecorationBar from "./components/DecorationBar.vue";
    import Navigator from "./components/Navigator.vue";
</script>

<template>
    <DecorationBar />
    <div id="innerframe" class="p-8 w-screen h-screen bg-base-200 overflow-hidden">
        <Navigator />
        <RouterView v-slot="{ Component }">
            <VFade>
                <component :is="Component" />
            </VFade>
        </RouterView>
    </div>
</template>

