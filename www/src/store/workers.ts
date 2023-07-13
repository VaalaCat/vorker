import { VorkerSettingsProperties } from '@/types/workers'
import { atom } from 'nanostores'

export const $code = atom('')

export const $vorkerSettings = atom<undefined | VorkerSettingsProperties>(
  undefined
)
