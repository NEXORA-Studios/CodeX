<script setup lang="ts">
    import { onMounted, ref } from "vue";
    import { Quit, WindowMinimise, WindowToggleMaximise, Environment } from "@/modules/wailsjs/runtime";

    const Architecture = ref("");
    const isFullscreen = ref(false);

    async function toggleMax() {
        WindowToggleMaximise();
        isFullscreen.value = !isFullscreen.value;
    }

    onMounted(async () => {
        Architecture.value = (await Environment()).platform;
    });
</script>

<template>
    <div id="dragable-bar" class="fixed top-0 left-0 w-screen h-8 flex" style="--drag-region: y" @dblclick="toggleMax()">
        <div class="mr-0.5 ml-auto flex group" v-if="Architecture != 'darwin'">
            <button
                class="btn btn-xs rounded-full btn-success scale-60 -mr-1 opacity-0 group-hover:opacity-100 hover:opacity-100 transition-opacity duration-150 ease-in-out"
                tabindex="-1"
                @click="WindowMinimise()">
                <i class="icon-[ic--round-minus] size-4 -mx-2" @click="WindowMinimise()" />
            </button>
            <button
                class="btn btn-xs rounded-full btn-warning scale-60 -mr-1 opacity-0 group-hover:opacity-100 hover:opacity-100 transition-opacity duration-150 ease-in-out"
                tabindex="-1"
                @click="toggleMax">
                <i class="icon-[ic--round-fullscreen] size-4 -mx-2" v-if="!isFullscreen" @click="toggleMax" />
                <i class="icon-[ic--round-fullscreen-exit] size-4 -mx-2" v-else @click="toggleMax" />
            </button>
            <button
                class="btn btn-xs rounded-full btn-error scale-60 opacity-0 group-hover:opacity-100 hover:opacity-100 transition-opacity duration-150 ease-in-out"
                tabindex="-1"
                @click="Quit()">
                <i class="icon-[ic--round-close] size-4 -mx-2" @click="Quit()" />
            </button>
        </div>
        <div class="ml-0.5 mr-auto flex group" v-else>
            <button
                class="btn btn-xs rounded-full btn-error scale-60 -mr-1 opacity-0 group-hover:opacity-100 hover:opacity-100 transition-opacity duration-150 ease-in-out"
                tabindex="-1"
                @click="Quit()">
                <i class="icon-[ic--round-close] size-4 -mx-2" @click="Quit()" />
            </button>
            <button
                class="btn btn-xs rounded-full btn-warning scale-60 -mr-1 opacity-0 group-hover:opacity-100 hover:opacity-100 transition-opacity duration-150 ease-in-out"
                tabindex="-1"
                @click="WindowMinimise()" >
                <i class="icon-[ic--round-minus] size-4 -mx-2" @click="WindowMinimise()" />
            </button>
            <button
                class="btn btn-xs rounded-full btn-success scale-60 -mr-1 opacity-0 group-hover:opacity-100 hover:opacity-100 transition-opacity duration-150 ease-in-out"
                tabindex="-1"
                @click="toggleMax">
                <i class="icon-[ic--round-fullscreen] size-4 -mx-2" v-if="!isFullscreen" @click="toggleMax" />
                <i class="icon-[ic--round-fullscreen-exit] size-4 -mx-2" v-else @click="toggleMax" />
            </button>
        </div>
    </div>
</template>
