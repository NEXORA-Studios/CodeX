<template>
    <div class="CUnderline w-fit">
        <slot></slot>
    </div>
</template>

<style lang="css" scoped>
    .CUnderline {
        background: linear-gradient(to right, var(--cudl-color) 0%, var(--cudl-color) 100%) no-repeat;
        background-size: 0 2px;
        background-position: right bottom;
        transition: background-size var(--cudl-dur);
    }
    .CUnderline:hover {
        background-position: left bottom;
        background-size: 100% 2px;
    }
</style>
