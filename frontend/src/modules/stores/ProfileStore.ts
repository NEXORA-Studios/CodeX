import { defineStore } from "pinia";
import type { IProfile } from "@/types";

interface IProfileItem {
    GUID: IProfile["GUID"];
    Name: IProfile["Name"];
    Avatar: IProfile["Avatar"];
    IDE: IProfile["IDE"];
    Settings: IProfile["Settings"];
    CurrentCourseID: IProfile["CurrentCourseID"];
}

interface IProfileStoreState {
    profiles: IProfileItem[];
    current_profile: string | null;
}

export const useProfileStore = defineStore("profile", {
    state: (): IProfileStoreState => ({
        profiles: [],
        current_profile: null,
    }),

    getters: {
        getAllProfiles: (state) => state.profiles,
        getCurrentProfile: (state): IProfileItem | null => state.profiles.find((p) => p.GUID === state.current_profile) || null,
    },

    actions: {
        addProfile(profiles: IProfileStoreState["profiles"][0]) {
            this.profiles.push(profiles);
        },
        setProfiles(profiles: IProfileStoreState["profiles"]) {
            this.profiles = profiles;
        },
        selectProfile(guid: IProfileStoreState["current_profile"]) {
            this.current_profile = guid;
        },
    },
    
    persist: true,
});
