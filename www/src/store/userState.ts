import { UserInfo } from '@/types/body'
import { atom, useAtom } from 'jotai'
import { userInfo } from 'os'

// Create your atoms and derivatives
export const UserAtom = atom<UserInfo | undefined>({} as UserInfo)