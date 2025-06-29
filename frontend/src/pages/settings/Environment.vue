<script setup lang="ts">
    import { onMounted, ref } from "vue";
    import CUndeline from "@/animation/CUndeline.vue";
    import GWaveText from "@/animation/GWaveText.vue";
    import VSlideIn from "@/animation/VSlideIn.vue";
    import { dEnvLogo } from "@/modules/Static";
    import type { WaveTextInstance } from "@/types";
    import * as $Env from "@/modules/wailsjs/go/bind/EnvBind";

    const t2 = ref<WaveTextInstance | null>(null);
    const refs = ref<(WaveTextInstance | null)[]>([]);
    const refs2 = ref<(WaveTextInstance | null)[]>([]);

    const devenv = ref<{ [key: string]: string }>({
        nodejs: "",
        npm: "",
        yarn: "",
        pnpm: "",
        go: "",
        python: "",
        rustup: "",
        rustc: "",
        cargo: "",
    });

    async function getEnvs() {
        const _env = devenv.value;
        // NodeJS
        const node = await $Env.GetNodeVersion();
        _env["nodejs"] = node.replace("v", "");
        // NPM
        const npm = await $Env.GetNPMVersion();
        _env["npm"] = npm.replace("v", "");
        // Yarn
        const yarn = await $Env.GetYarnVersion();
        _env["yarn"] = yarn.replace("v", "");
        // PNPM
        const pnpm = await $Env.GetPNPMVersion();
        _env["pnpm"] = pnpm.replace("v", "");
        // Golang
        const go = await $Env.GetGoVersion();
        _env["go"] = go.replace("go version go", "");
        // Python
        const python = await $Env.GetPythonVersion();
        _env["python"] = python.replace("Python ", "");
        // Rust
        const rustup = await $Env.GetRustupVersion();
        _env["rustup"] = rustup.split("\n")[0].replace("rustup ", "");
        // RustC
        const rustc = await $Env.GetRustVersion();
        _env["rustc"] = rustc.split("\n")[0].replace("rustc ", "");
        // Cargo
        const cargo = await $Env.GetCargoVersion();
        _env["cargo"] = cargo.split("\n")[0].replace("cargo ", "");
    }

    const onDone1 = () => {
        setTimeout(() => {
            t2.value?.startTyping();
        }, 50);
    };
    const onDone2 = () => {
        setTimeout(() => {
            refs.value.forEach((ref) => {
                ref?.startTyping();
            });
            refs2.value.forEach((ref) => {
                ref?.startTyping();
            });
        }, 50);
    };

    onMounted(() => {
        setTimeout(() => {
            getEnvs();
        }, 1500);
    });
</script>

<template>
    <main class="mt-4 p-4 pb-8 w-full h-full flex flex-col">
        <section class="flex items-center">
            <CUndeline class="cursor-pointer" style="--cudl-color: var(--color-base-content); --cudl-dur: 0.3s" @click="$router.push('/settings')">
                <GWaveText class="text-xl" content="←  返回上一页" :speed="0.3" start-when-mount @done="onDone1()" />
            </CUndeline>
            <GWaveText class="text-xl ml-4" content="检查开发环境" :speed="0.3" ref="t2" @done="onDone2()" />
        </section>
        <VSlideIn direction="top" :duration="0.75" :delay="0.6 - 0.5">
            <div class="overflow-x-auto overflow-y-hidden mt-8 w-3/5 mx-auto">
                <table class="table">
                    <thead>
                        <tr>
                            <th></th>
                            <th>开发环境</th>
                            <th class="w-100">版本</th>
                        </tr>
                    </thead>
                    <tbody>
                        <tr v-for="(value, key) in devenv" :key="key">
                            <td>
                                <component :is="dEnvLogo[key]" class="size-6" />
                            </td>
                            <td>
                                <!-- @vue-ignore -->
                                <GWaveText class="-mr-16" :ref="(el: WaveTextInstance) => refs[Object.keys(devenv).indexOf(key)] = el" :content="key" :speed="0.5" />
                            </td>
                            <td>
                                <!-- @vue-ignore -->
                                <GWaveText
                                    class="-mr-16"
                                    :ref="(el: WaveTextInstance) => refs2[Object.keys(devenv).indexOf(key)] = el"
                                    :content="value.length === 0 ? '获取中...' : value === 'error' ? '未安装 / 读取版本时发生错误' : value"
                                    :speed="0.5" />
                            </td>
                        </tr>
                    </tbody>
                </table>
            </div>
        </VSlideIn>
    </main>
</template>
