import { deleteReq, getReq, patchReq, postReq } from "@/utils/http"
import { AxiosResponse } from "axios"
import { WorkerItem } from "@/types/workers"

export const GetWorker = (uid: string): Promise<AxiosResponse<any, any>> => {
	return getReq(`/api/worker/${uid}`, {})
}

export const GetWorkers = (offset: number, limit: number): Promise<AxiosResponse<any, any>> => {
	return getReq(`/api/workers/${offset}/${limit}`, {})
}

export const GetAllWorkers = (): Promise<AxiosResponse<any, any>> => {
	return getReq("/api/workers", {})
}

export const CreateWorker = (worker: WorkerItem): Promise<AxiosResponse<any, any>> => {
	return postReq("/api/worker/create", worker)
}

export const DeleteWorker = (uid: string): Promise<AxiosResponse<any, any>> => {
	return deleteReq(`/api/worker/${uid}`, {})
}

export const UpdateWorker = (uid: string, worker: WorkerItem): Promise<AxiosResponse<any, any>> => {
	return patchReq(`/api/worker/${uid}`, worker)
}

export const FlushWorker = (uid: string): Promise<AxiosResponse<any, any>> => {
	return getReq(`/api/worker/flush/${uid}`, {})
}

export const FlushAllWorkers = (): Promise<AxiosResponse<any, any>> => {
	return getReq(`/api/workers/flush`, {})
}