import { useRoute, useRouter } from "vue-router";

type ISkipMethod = "equals" | "startsWith" | "includes";

export function useHandleNavigation() {
    const route = useRoute();
    const router = useRouter();

    const handleNavigation = (path: string, skipMethod?: ISkipMethod, skip?: string) => {
        if (skip) {
            if (skipMethod === "equals" && skip === route.path) return;
            if (skipMethod === "startsWith" && route.path.startsWith(skip)) return;
            if (skipMethod === "includes" && route.path.includes(skip)) return;
        }
        router.push(path);
    };

    return handleNavigation;
}
