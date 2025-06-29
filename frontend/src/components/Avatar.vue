<script setup lang="ts">
    import { useRouter } from "vue-router";
    import { useProfileStore } from "@/modules/stores";

    const profileStore = useProfileStore();
    const router = useRouter();

    const handleLogout = () => {
        profileStore.selectProfile(null);
        router.push("/setup/profile");
    };
</script>

<template>
    <div class="dropdown dropdown-end">
        <div tabindex="0" role="button" class="btn btn-circle avatar avatar-placeholder">
            <div class="bg-neutral text-neutral-content w-8 rounded-full">
                <i class="icon-[ic--round-person-outline] size-5 -translate-y-0.25 -translate-x-[0.5px]" />
            </div>
        </div>
        <ul tabindex="0" class="menu menu-sm dropdown-content bg-base-100 rounded-box z-1 mt-3 w-52 p-2 shadow">
            <li><a>设置</a></li>
            <li @click="handleLogout"><a>登出</a></li>
        </ul>
    </div>
</template>
