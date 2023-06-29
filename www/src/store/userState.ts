import { UserInfo } from '@/types/body'
import { atom } from 'nanostores'

// Create your atoms and derivatives
export const $user = atom<UserInfo | undefined>({} as UserInfo)
