<script setup lang="ts">
    import { computed, ref, watch, type ComponentPublicInstance } from "vue";
    import CUndeline from "@/animation/CUndeline.vue";
    import GWaveText from "@/animation/GWaveText.vue";
    import VSlideIn from "@/animation/VSlideIn.vue";
    import { dIdeLogo } from "@/modules/Static";
    import { useConfigStore } from "@/modules/stores";
    import { SetConfig } from "@/modules/wailsjs/go/bind/ConfigBind";
    import type { IConfig, WaveTextInstance } from "@/types";
    import FlowContainer from "@/components/FlowContainer.vue";

    const configStore = useConfigStore();
    const t2 = ref<WaveTextInstance | null>();

    const onDone = () => {
        setTimeout(() => {
            t2.value?.startTyping();
        }, 50);
    };

    const mSelectedIDE = ref("");
    const mIdeName = ref("");
    const mIdeTarget = ref("");
    const sInvalidInput = computed(
        () => mSelectedIDE.value === "" || mIdeTarget.value === "" || (mSelectedIDE.value === "custom" && (mIdeName.value === "" || mIdeTarget.value === ""))
    );

    watch(
        () => mSelectedIDE.value,
        (newValue) => {
            if (newValue === "custom") {
                mIdeName.value = "";
            }
            mIdeTarget.value = "";
        }
    );

    const handleAddIDE = () => {
        if (sInvalidInput.value) return;
        configStore.addNewIde({
            GUID: crypto.randomUUID(),
            Name: mSelectedIDE.value === "custom" ? mIdeName.value : mSelectedIDE.value,
            LaunchTarget: mIdeTarget.value,
        });
        SetConfig(configStore.getAll as IConfig);
        mSelectedIDE.value = "";
        mIdeName.value = "";
        mIdeTarget.value = "";
    };

    const handleRemoveIDE = (guid: string) => {
        configStore.removeIde(guid);
        SetConfig(configStore.getAll as IConfig);
    };
</script>

<template>
    <main class="mt-4 p-4 pb-8 w-full h-full flex flex-col">
        <section class="flex items-center">
            <CUndeline class="cursor-pointer" style="--cudl-color: var(--color-base-content); --cudl-dur: 0.3s" @click="$router.push('/settings')">
                <GWaveText class="text-xl" content="←  返回上一页" :speed="0.3" start-when-mount @done="onDone" />
            </CUndeline>
            <GWaveText class="text-xl ml-4" content="IDE 设置" :speed="0.3" ref="t2" />
        </section>
        <FlowContainer class="pb-8">
            <VSlideIn direction="top" :duration="0.75" :delay="0.4">
                <div class="overflow-x-auto mt-8">
                    <table class="table">
                        <!-- head -->
                        <thead>
                            <tr>
                                <th></th>
                                <th>名称</th>
                                <th>实例 ID</th>
                                <th>目标启动命令</th>
                                <th>操作</th>
                            </tr>
                        </thead>
                        <tbody>
                            <!-- row 1 -->
                            <tr v-for="ide of configStore.getAvailableIDE" :key="ide.GUID">
                                <th>
                                    <component :is="dIdeLogo[ide.Name]" class="size-6" />
                                </th>
                                <td>{{ ide.Name }}</td>
                                <td>
                                    {{ ide.GUID.split("-")[0] }}
                                </td>
                                <td>
                                    {{ ide.LaunchTarget }}
                                </td>
                                <td>
                                    <button class="btn btn-sm btn-error btn-outline" @click="handleRemoveIDE(ide.GUID)">删除</button>
                                </td>
                            </tr>
                        </tbody>
                    </table>
                </div>
            </VSlideIn>
            <VSlideIn direction="top" :duration="0.75" :distance="32" :delay="0.75">
                <fieldset class="fieldset bg-base-200 rounded-box w-full border border-base-content/10 p-4 mt-4">
                    <legend class="fieldset-legend">增加新 IDE</legend>

                    <fieldset class="fieldset -mt-4 w-full">
                        <label class="label">IDE 类型 <span class="-ml-1 text-error">*</span></label>
                        <select class="select w-full" v-model="mSelectedIDE">
                            <option disabled selected value="">选择一种 IDE</option>
                            <option value="Visual Studio Code">Visual Studio Code</option>
                            <option value="Visual Studio">Visual Studio</option>
                            <option value="Trae">Trae (海外版)</option>
                            <option value="Trae CN">Trae CN (中国版)</option>
                            <option value="JetBrains Webstorm">JetBrains WebStorm</option>
                            <option value="JetBrains PyCharm">JetBrains PyCharm</option>
                            <option value="custom">自定义</option>
                        </select>
                    </fieldset>

                    <fieldset class="fieldset w-full" v-if="mSelectedIDE === 'custom'">
                        <label class="label">自定义 IDE 名称 <span class="mr-auto -ml-1 text-error">*</span></label>
                        <input type="text" class="input w-full" placeholder="例如 Visual Studio Code" v-model="mIdeName" />
                    </fieldset>

                    <fieldset class="fieldset w-full" v-if="mSelectedIDE !== ''">
                        <label class="label">应用程序启动路径 / 命令 / URL Scheme <span class="-ml-1 text-error">*</span></label>
                        <input type="text" class="input w-full" placeholder='例如 "C:\Program Files\Microsoft VS Code\Code.exe %s"' v-model="mIdeTarget" />
                        <label class="label">使用 %s 占位符表示要打开的文件路径，不需要使用引号包裹</label>
                        <label class="label -mt-2">我们建议优先通过命令行启动，例如使用 "code" "trae" 等命令</label>
                        <label class="label -mt-2">如果从来没有安装过 IDE，请<span class="underline cursor-pointer -ml-1">点击这里</span></label>
                    </fieldset>

                    <button v-if="mSelectedIDE !== ''" class="btn w-full" :disabled="sInvalidInput" @click="handleAddIDE">确认添加</button>
                </fieldset>
            </VSlideIn>
        </FlowContainer>
    </main>
</template>
