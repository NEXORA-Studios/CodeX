import { defineStore } from "pinia";
import type { IConfig } from "@/types";

interface IConfigStoreState {
    AvailableIDE: IConfig["AvailableIDE"];
    Workspace: IConfig["Workspace"];
}

export const useConfigStore = defineStore("config", {
    state: (): IConfigStoreState => ({
        AvailableIDE: [],
        Workspace: "",
    }),
    getters: {
        getAll: (state) => ({
            AvailableIDE: state.AvailableIDE,
            Workspace: state.Workspace,
        }),
        getAvailableIDE: (state) => state.AvailableIDE,
        getWorkspace: (state) => state.Workspace,
    },
    actions: {
        addNewIde(ide: IConfigStoreState["AvailableIDE"][0]) {
            this.AvailableIDE.push(ide);
        },
        setWorkspace(workspace: IConfigStoreState["Workspace"]) {
            this.Workspace = workspace;
        },
        setConfig(config: IConfigStoreState) {
            this.AvailableIDE = config.AvailableIDE;
            this.Workspace = config.Workspace;
        },
    },
});
