<script setup lang="ts">
    import { computed, onMounted, ref, watch } from "vue";
    import { useRouter } from "vue-router";
    // Animations
    import GWaveText from "@/animation/GWaveText.vue";
    import VSlideIn from "@/animation/VSlideIn.vue";
    // Modules
    import { useConfigStore, useProfileStore } from "@/modules/stores";
    import { SetConfig } from "@/modules/wailsjs/go/bind/ConfigBind";
    import { UpsertProfile } from "@/modules/wailsjs/go/bind/ProfileBind";
    import { dIdeLogo } from "@/modules/Static";
    import type { IConfig, IProfile, WaveTextInstance } from "@/types";

    const router = useRouter();
    const configStore = useConfigStore();
    const profileStore = useProfileStore();

    // 文本动画相关
    interface Line {
        content: string;
        class: string;
        speed: number;
        next?: Function;
    }

    const lines: Line[] = [
        { content: "欢迎来到 Code X", class: "text-4xl absolute top-8 left-8", speed: 0.75 },
        { content: "Code X 不做假教学，也不构造封闭环境", class: "text-xl absolute top-24 left-8", speed: 1 },
        { content: "它用清晰的引导、自动化的项目搭建", class: "text-xl absolute top-30 left-8", speed: 1 },
        { content: "配合无缝集成的 IDE，带你进入真实的编程世界", class: "text-xl absolute top-36 left-8", speed: 1 },
        { content: "在这里，每一段学习过程，都是通向独立作品的构建旅程", class: "text-xl absolute top-42 left-8", speed: 1 },
        {
            content: "通过简单的步骤，你可以探索无限的可能！",
            class: "text-xl absolute top-58 left-8",
            speed: 1,
            next: () => {
                stepIndex.value = 1;
            },
        },
    ];

    const refs = ref<(WaveTextInstance | null)[]>([]);

    function handleDone(index: number, line: Line) {
        const next = refs.value[index + 1];
        if (next?.startTyping) {
            next.startTyping();
        } else if (line.next) {
            line.next();
        }
    }

    // 步骤相关
    const stepIndex = ref(0);
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

    const mWorkspace = ref(configStore.getWorkspace);
    const _hardPathRegex = /^(?:[a-zA-Z]:\\(?:[^\\/:*?"<>|\r\n]+\\)*[^\\/:*?"<>|\r\n]+|\/(?:[^\/\r\n]+\/)*[^\/\r\n.]+)$/;
    const sInvalidInput2 = computed(() => mWorkspace.value === "" || !_hardPathRegex.test(mWorkspace.value));
    const sFinishInput2 = computed(() => mWorkspace.value === configStore.getWorkspace);

    const handleWorkspace = () => {
        if (sInvalidInput2.value) return;
        configStore.setWorkspace(mWorkspace.value);
        SetConfig(configStore.getAll as IConfig);
    };

    const mProfileName = ref("");
    const mProfileIDE = ref("");
    const mProfileLang = ref("zh_CN");
    const mProfileTheme = ref("codexdark");

    const sInvalidInput3 = computed(() => mProfileName.value === "" || mProfileIDE.value === "" || mProfileLang.value === "" || mProfileTheme.value === "");

    const handleAddProfile = () => {
        if (sInvalidInput3.value) return;
        const _p = {
            GUID: crypto.randomUUID(),
            Name: mProfileName.value,
            Avatar: false,
            IDE: mProfileIDE.value,
            Settings: {
                Theme: mProfileTheme.value,
                Language: mProfileLang.value,
            },
            CurrentCourseID: {
                ID: "0",
                SubID: "0",
            },
        };
        profileStore.addProfile(_p);
        UpsertProfile(_p as IProfile);
        mProfileName.value = "";
        mProfileIDE.value = "";
    };

    const handleLogin = (guid: string) => {
        profileStore.selectProfile(guid);
        router.push("/");
    };

    // 启动钩子
    onMounted(() => {
        // 检查设置项
        if (configStore.getAvailableIDE.length !== 0 && configStore.getWorkspace !== "" && profileStore.getAllProfiles.length !== 0) {
            console.log("[Vaildator] 当前配置项均已配置，直接登录");
            router.push("/setup/profile");
            return;
        }

        // 开始展示文本
        refs.value[0]?.startTyping();
    });
</script>

<template>
    <main class="mt-4 pb-8 w-full h-full flex gap-8">
        <section class="w-144 h-full relative">
            <!-- @vue-ignore -->
            <GWaveText
                v-for="(line, i) in lines"
                :key="i"
                :ref="(el: WaveTextInstance) => refs[i] = el"
                :content="line.content"
                :class="line.class"
                :speed="line.speed"
                @done="() => handleDone(i, line)" />
            <VSlideIn
                direction="left"
                :duration="0.75"
                :delay="stepIndex === 1 ? 0.75 : 0"
                extra-class="absolute bottom-8 left-8"
                v-if="0 < stepIndex && stepIndex < 3 && ((stepIndex === 1 && configStore.getAvailableIDE.length !== 0) || (stepIndex === 2 && configStore.getWorkspace !== ''))">
                <button class="btn btn-primary" @click="stepIndex += 1">下一步</button>
            </VSlideIn>
        </section>
        <VSlideIn direction="top" :duration="0.75" v-if="stepIndex > 0" class="w-[calc(100%-calc(var(--spacing)*144))] h-full">
            <section class="relative pt-9 flex flex-col gap-8 items-center h-full">
                <ul class="steps w-full">
                    <li class="step step-primary" :class="{ 'step-primary': stepIndex > 0 }">添加 IDE</li>
                    <li class="step" :class="{ 'step-primary': stepIndex > 1 }">设置工作区</li>
                    <li class="step" :class="{ 'step-primary': stepIndex > 2 }">设置用户</li>
                </ul>
                <!-- Step 1 -->
                <VSlideIn direction="top" :duration="1" class="w-full" v-if="stepIndex === 1">
                    <div class="w-full">
                        <fieldset class="fieldset bg-base-200 rounded-box w-full border border-base-content/10 p-4">
                            <legend class="fieldset-legend">目前已配置</legend>
                            <span class="text-center opacity-50" v-if="configStore.getAvailableIDE.length === 0">无</span>
                            <div class="overflow-x-auto" v-else>
                                <table class="table">
                                    <!-- head -->
                                    <thead>
                                        <tr>
                                            <th></th>
                                            <th>名称</th>
                                            <th>实例 ID</th>
                                            <th>目标启动命令</th>
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
                                        </tr>
                                    </tbody>
                                </table>
                            </div>
                        </fieldset>

                        <fieldset class="fieldset bg-base-200 rounded-box w-full border border-base-content/10 p-4">
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
                    </div>
                </VSlideIn>
                <!-- Step 2 -->
                <VSlideIn direction="top" :duration="1" class="w-full" v-if="stepIndex === 2">
                    <div class="w-full">
                        <fieldset class="fieldset bg-base-200 rounded-box w-full border border-base-content/10 p-4">
                            <legend class="fieldset-legend">设置工作区</legend>

                            <label class="label">工作区路径 <span class="-ml-1 text-error">*</span></label>
                            <input type="text" class="input w-full" placeholder="例如 D:\projects\" v-model="mWorkspace" />
                            <label class="label">提供你希望以后用来存放课程文件的地方，需要是绝对路径，末尾不得有 "/" 或 "\"</label>

                            <button class="btn w-full" :disabled="sInvalidInput2 || sFinishInput2" @click="handleWorkspace">确认</button>
                        </fieldset>
                    </div>
                </VSlideIn>
                <!-- Step 3 -->
                <VSlideIn direction="top" :duration="1" class="w-full" v-if="stepIndex === 3">
                    <div class="w-full">
                        <fieldset class="fieldset bg-base-200 rounded-box w-full border border-base-content/10 p-4">
                            <legend class="fieldset-legend">目前已有账户</legend>
                            <span class="text-center opacity-50" v-if="profileStore.getAllProfiles.length === 0">无</span>
                            <div class="overflow-x-auto" v-else>
                                <table class="table">
                                    <!-- head -->
                                    <thead>
                                        <tr>
                                            <th></th>
                                            <th>用户名</th>
                                            <th>IDE</th>
                                            <th>操作</th>
                                        </tr>
                                    </thead>
                                    <tbody>
                                        <!-- row 1 -->
                                        <tr v-for="[index, profile] of Object.entries(profileStore.getAllProfiles)" :key="profile.GUID">
                                            <th>{{ Number(index) + 1 }}</th>
                                            <td>{{ profile.Name }}</td>
                                            <td>
                                                {{ configStore.AvailableIDE.find((i) => i.GUID === profile.IDE)?.Name }} (实例
                                                {{ configStore.AvailableIDE.find((i) => i.GUID === profile.IDE)?.GUID.split("-")[0] ?? "Unknown" }})
                                            </td>
                                            <td>
                                                <button class="btn" @click="handleLogin(profile.GUID)">登录</button>
                                            </td>
                                        </tr>
                                    </tbody>
                                </table>
                            </div>
                        </fieldset>

                        <fieldset class="fieldset bg-base-200 rounded-box w-full border border-base-content/10 p-4">
                            <legend class="fieldset-legend">增加新账户</legend>

                            <fieldset class="fieldset -mt-4 w-full">
                                <label class="label">用户名 <span class="mr-auto -ml-1 text-error">*</span></label>
                                <input type="text" class="input w-full" placeholder="例如 MoYuan-CN" v-model="mProfileName" />
                            </fieldset>

                            <fieldset class="fieldset w-full">
                                <label class="label">使用的 IDE <span class="-ml-1 text-error">*</span></label>
                                <select class="select w-full" v-model="mProfileIDE">
                                    <option disabled selected value="">选择一种 IDE</option>
                                    <option v-for="ide in configStore.getAvailableIDE" :value="ide.GUID">{{ ide.Name }} (实例 {{ ide.GUID.split("-")[0] }})</option>
                                </select>
                            </fieldset>

                            <div class="grid grid-cols-2 gap-x-2">
                                <fieldset class="fieldset w-full">
                                    <label class="label">语言 <span class="-ml-1 text-error">*</span></label>
                                    <select class="select w-full" v-model="mProfileLang">
                                        <option disabled selected value="zh_CN">简体中文</option>
                                    </select>
                                </fieldset>
                                <fieldset class="fieldset w-full">
                                    <label class="label">主题 <span class="-ml-1 text-error">*</span></label>
                                    <select class="select w-full" v-model="mProfileTheme">
                                        <option disabled selected value="codexdark">CodeX Dark</option>
                                    </select>
                                </fieldset>
                            </div>

                            <button class="btn w-full" :disabled="sInvalidInput3" @click="handleAddProfile">确认添加</button>
                        </fieldset>
                    </div>
                </VSlideIn>
            </section>
        </VSlideIn>
    </main>
</template>

<style lang="css" scoped></style>
