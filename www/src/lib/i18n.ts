import { atom } from 'nanostores'

const zh = {
  login: '',
  logout: '',
  register: '',
  username: '',
  password: '',
  email: '',
  submit: '',
  code: '',
  editor: '',
  edit: '编辑',
  run: '',
  sync: '',
  create: '',
  delete: '',
  deleteWorker: '',

  workerConfirmDelete: '',
  workerDeleteSuccess: '',
  workerCreateSuccess: '',
  workerSaveSuccess: '',
  workerSyncSuccess: '',
  backToList: '',
  noWorker: '',

  notLoggedInPrompt: '',
  loggingOutPrompt: '',
  loginSuccess: '',
  loginFailed: '',
  registerSuccess: '',
  registerFailed: '',
}

const en = {
  login: '',
  logout: '',
  register: '',
  username: '',
  password: '',
  email: '',
  submit: '',
  code: '',
  editor: '',
  edit: '编辑',
  run: '',
  sync: '',
  create: '',
  delete: '',
  deleteWorker: '',

  workerConfirmDelete: '',
  workerDeleteSuccess: '',
  workerCreateSuccess: '',
  workerSaveSuccess: '',
  workerSyncSuccess: '',
  backToList: '',
  noWorker: '',

  notLoggedInPrompt: '',
  loggingOutPrompt: '',
  loginSuccess: '',
  loginFailed: '',
  registerSuccess: '',
  registerFailed: '',
}

export type Key = keyof typeof zh & keyof typeof en

export const i18n = (k: Key) => {
  return { zh, en }[$lang.get().split('-')[0] || 'en']![k] || k
}

export const t = new Proxy(
  {},
  {
    get(_, k) {
      return i18n(k as Key)
    },
  }
) as Record<Key, string>

export const $lang = atom(
  typeof navigator !== 'undefined' ? navigator.language : 'en'
)
