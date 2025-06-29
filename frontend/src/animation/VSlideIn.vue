<template>
    <Transition :name="transitionName" appear>
        <div
            :style="{
                '--distance': distance + 'px',
                '--duration': duration + 's',
                '--delay': delay + 's',
            }"
            :class="extraClass">
            <slot />
        </div>
    </Transition>
</template>

<script setup lang="ts">
    import { computed } from "vue";

    const props = defineProps<{
        direction?: "left" | "right" | "top" | "bottom";
        duration?: number;
        delay?: number;
        distance?: number;
        extraClass?: string;
    }>();

    const direction = props.direction ?? "left";
    const distance = props.distance ?? 100;
    const duration = props.duration ?? 0.5;
    const delay = props.delay ?? 0;

    const transitionName = computed(() => `fade-slide-${direction}`);
</script>

<style scoped>
    /* 通用过渡 */
    div {
        transition: opacity var(--duration) ease var(--delay), transform var(--duration) ease var(--delay);
        will-change: transform, opacity;
    }

    /* Fade + Slide 进入 */
    .fade-slide-left-enter-from,
    .fade-slide-left-leave-to {
        opacity: 0;
        transform: translateX(calc(-1 * var(--distance)));
    }

    .fade-slide-right-enter-from,
    .fade-slide-right-leave-to {
        opacity: 0;
        transform: translateX(var(--distance));
    }

    .fade-slide-top-enter-from,
    .fade-slide-top-leave-to {
        opacity: 0;
        transform: translateY(calc(-1 * var(--distance)));
    }

    .fade-slide-bottom-enter-from,
    .fade-slide-bottom-leave-to {
        opacity: 0;
        transform: translateY(var(--distance));
    }

    /* 保持进入后为正常状态 */
    .fade-slide-left-enter-to,
    .fade-slide-right-enter-to,
    .fade-slide-top-enter-to,
    .fade-slide-bottom-enter-to {
        opacity: 1;
        transform: translate(0, 0);
    }

    /* 进入和离开过渡都开启 */
    .fade-slide-left-enter-active,
    .fade-slide-left-leave-active,
    .fade-slide-right-enter-active,
    .fade-slide-right-leave-active,
    .fade-slide-top-enter-active,
    .fade-slide-top-leave-active,
    .fade-slide-bottom-enter-active,
    .fade-slide-bottom-leave-active {
        transition: opacity var(--duration) ease var(--delay), transform var(--duration) ease var(--delay);
    }
</style>
