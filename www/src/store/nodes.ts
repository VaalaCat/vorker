import { PingMap, PingMapList } from '@/types/nodes'
import { atom } from 'nanostores'

export const $nodeStatus = atom<PingMapList>({})
