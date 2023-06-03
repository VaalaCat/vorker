import api from './http'
import { VorkerSettingsProperties, WorkerItem } from '@/types/workers'

export const GetWorker = (uid: string) => {
  return api
    .get<{ data: WorkerItem[] }>(`/api/worker/${uid}`)
    .then((res) => res.data.data?.[0])
}

export const GetWorkers = (offset: number, limit: number) => {
  return api
    .get<{ data: WorkerItem[] }>(`/api/workers/${offset}/${limit}`)
    .then((res) => res.data.data)
}

export const GetAllWorkers = () => {
  return api
    .get<{ data: WorkerItem[] }>('/api/allworkers')
    .then((res) => res.data.data)
}

export const CreateWorker = (worker: WorkerItem) => {
  return api.post('/api/worker/create', worker).then((res) => res.data)
}

export const DeleteWorker = (uid: string) => {
  return api.delete(`/api/worker/${uid}`, {}).then((res) => res.data)
}

export const UpdateWorker = (uid: string, worker: WorkerItem) => {
  return api.patch(`/api/worker/${uid}`, worker).then((res) => res.data)
}

export const FlushWorker = (uid: string) => {
  return api.get(`/api/worker/flush/${uid}`, {}).then((res) => res.data)
}

export const FlushAllWorkers = () => {
  return api.get(`/api/workers/flush`, {}).then((res) => res.data)
}

export const GetAppConfig = () => {
  return api.get<{ data: VorkerSettingsProperties }>
    (`/api/vorker/config`, {}).then((res) => res.data.data)
}