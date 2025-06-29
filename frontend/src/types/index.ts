import type { ComponentPublicInstance } from "vue";
import { common } from "@/modules/wailsjs/go/models";

export interface IProfile extends common.Profile {}
export interface IConfig extends common.IConfig {}

export type WaveTextInstance = ComponentPublicInstance<{ startTyping: () => void }>;
