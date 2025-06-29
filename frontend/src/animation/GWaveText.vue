<template>
    <div class="w-fit" ref="textRef"></div>
</template>

<script setup lang="ts">
    import { onMounted, ref, watch } from "vue";
    import { gsap } from "gsap";

    const props = defineProps<{ content: string; speed: number; startWhenMount?: boolean }>();
    const _content = ref(props.content);
    const emit = defineEmits<{
        (e: "done"): void;
    }>();

    const textRef = ref<HTMLElement | null>(null);

    // 暴露启动函数
    function startTyping() {
        const content = props.content || "Undefined...";

        if (!textRef.value) return;

        function charWrapper(c: string) {
            let char = c;
            if (c === " ") char = "&nbsp;";
            return `<span style="display:inline-block">${char}</span>`;
        }

        // 将文本分割成单个字符
        textRef.value.innerHTML = content
            .split("")
            .map((char) => charWrapper(char))
            .join("");

        const words = textRef.value.querySelectorAll("span");
        gsap.from(words, {
            y: 25,
            opacity: 0,
            stagger: 0.1,
            duration: props.speed,
            ease: "back",
            onComplete() {
                emit("done");
            },
        }).duration(props.speed);
    }

    defineExpose({ startTyping });

    onMounted(() => {
        if (props.startWhenMount) {
            startTyping();
        }
    });

    watch(props, (newProps) => {
        if (newProps.content !== _content.value) {
            startTyping();
        }
    });
</script>
