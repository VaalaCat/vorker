import api from './http'
import { VorkerSettingsProperties, WorkerItem } from '@/types/workers'

export const getWorker = (uid: string) => {
  return api
    .get<{ data: WorkerItem[] }>(`/api/worker/${uid}`)
    .then((res) => res.data.data?.[0])
}

export const getWorkerCursor = (offset: number, limit: number) => {
  return api
    .get<{ data: WorkerItem[] }>(`/api/workers/${offset}/${limit}`)
    .then((res) => res.data.data)
}

export const getAllWorkers = () => {
  return api
    .get<{ data: WorkerItem[] }>('/api/allworkers')
    .then((res) => res.data.data)
}

export const createWorker = (worker: WorkerItem) => {
  return api.post('/api/worker/create', worker).then((res) => res.data)
}

export const deleteWorker = (uid: string) => {
  return api.delete(`/api/worker/${uid}`, {}).then((res) => res.data)
}

export const updateWorker = (uid: string, worker: WorkerItem) => {
  return api.patch(`/api/worker/${uid}`, worker).then((res) => res.data)
}

export const flushWorker = (uid: string) => {
  return api.get(`/api/worker/flush/${uid}`, {}).then((res) => res.data)
}

export const flushAllWorkers = () => {
  return api.get(`/api/workers/flush`, {}).then((res) => res.data)
}

export const getAppConfig = () => {
  return api
    .get<{ data: VorkerSettingsProperties }>(`/api/vorker/config`, {})
    .then((res) => res.data.data)
}

export const runWorker = (uid: string) => {
  return api.get(`/api/worker/run/${uid}`, {}).then((res) => res.data)
}