import { VorkerSettingsProperties } from "@/types/workers";
import { atom } from "jotai";

export const CodeAtom = atom("")

export const VorkerSettingsAtom =
	atom<undefined | VorkerSettingsProperties>(undefined)