<template>
    <div class="w-fit" ref="textRef"></div>
</template>

<script setup lang="ts">
    import { onMounted, ref } from "vue";
    import { gsap } from "gsap";

    const props = defineProps<{ content: string; speed: number; startWhenMount?: boolean }>();
    const emit = defineEmits<{
        (e: "done"): void;
    }>();

    const textRef = ref<HTMLElement | null>(null);
    const content = props.content || "Undefined...";

    // 暴露启动函数
    function startTyping() {
        if (!textRef.value) return;

        // 将文本分割成单个字符
        textRef.value.innerHTML = content
            .split("")
            .map((char) => `<span style="display:inline-block">${char}</span>`)
            .join("");

        const words = textRef.value.querySelectorAll("span");
        gsap.from(words, {
            y: 25,
            opacity: 0,
            stagger: 0.1,
            ease: "back",
            duration: props.speed,
            onComplete() {
                emit("done");
            },
        });
    }

    defineExpose({ startTyping });

    onMounted(() => {
        if (props.startWhenMount) {
            startTyping();
        }
    });
</script>
