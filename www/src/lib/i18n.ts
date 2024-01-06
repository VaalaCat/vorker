import { atom } from 'nanostores'

const zh = {
  login: '登陆',
  logout: '登出',
  register: '注册',
  username: '用户名',
  password: '密码',
  email: '邮箱',
  submit: '提交',
  code: '代码',
  editor: '编辑器',
  edit: '编辑',
  run: '运行',
  sync: '同步',
  create: '创建',
  delete: '删除',
  open: '打开',
  refresh: '刷新',
  deleteWorker: '删除函数',
  deleteNode: '删除节点',

  workerConfirmDelete: '确认删除函数',
  workerDeleteSuccess: '删除成功',
  workerCreateSuccess: '创建成功',
  workerSaveSuccess: '保存成功',
  workerSyncSuccess: '同步成功',
  nodeDeleteSuccess: '删除成功',
  nodeSyncSuccess: '同步成功',
  backToList: '返回列表',
  noWorkerPrompt: '空空如也',

  notLoggedInPrompt: '没有登陆，正在跳转...',
  loggingOutPrompt: '正在退出登陆...',
  loginSuccess: '登陆成功',
  loginFailed: '用户名或密码错误',
  registerSuccess: '注册成功',
  registerFailed: '注册失败',
}

const en = {
  login: 'Login',
  logout: 'Logout',
  register: 'Register',
  username: 'Username',
  password: 'Password',
  email: 'Email',
  submit: 'Submit',
  code: 'Code',
  editor: 'Editor',
  edit: 'Edit',
  run: 'Run',
  sync: 'Sync',
  create: 'New',
  delete: 'Delete',
  open: 'Open',
  refresh: 'Refresh',
  deleteWorker: 'Delete worker',
  deleteNode: 'Delete Node',

  workerConfirmDelete: 'Confirm delete worker',
  workerDeleteSuccess: 'Delete worker success',
  workerCreateSuccess: 'Create worker success',
  workerSaveSuccess: 'Save worker success',
  workerSyncSuccess: 'Sync worker success',
  nodeDeleteSuccess: 'Delete node success',
  nodeSyncSuccess: 'Sync node success',
  backToList: 'Back',
  noWorkerPrompt: 'No workers here.',

  notLoggedInPrompt: 'Not logged in, redirecting...',
  loggingOutPrompt: 'Logging out...',
  loginSuccess: '',
  loginFailed: 'Incorrect username or password',
  registerSuccess: 'Successfully registered',
  registerFailed: 'Failed to register',
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
