import { getReq, postReq } from "@/utils/http"
import { AxiosResponse } from "axios"
import { WorkerItem } from "@/types/workers"

export const GetWorker = (uid: string): Promise<AxiosResponse<any, any>> => {
	return getReq(`/api/workers/${uid}`, {})
}

export const GetWorkers = (offset: number, limit: number): Promise<AxiosResponse<any, any>> => {
	return getReq(`/workers/${offset}/${limit}`, {})
}

export const GetAllWorkers = (): Promise<AxiosResponse<any, any>> => {
	return getReq("/workers", {})
}

export const CreateWorker = (worker: WorkerItem): Promise<AxiosResponse<any, any>> => {
	return postReq("/workers", worker)
}

export const DeleteWorker = (uid: string): Promise<AxiosResponse<any, any>> => {
	return postReq(`/workers/${uid}`, {})
}

export const UpdateWorker = (uid: string, worker: WorkerItem): Promise<AxiosResponse<any, any>> => {
	return postReq(`/workers/${uid}`, worker)
}

export const FlushWorker = (uid: string): Promise<AxiosResponse<any, any>> => {
	return getReq(`/workers/flush/${uid}`, {})
}

export const FlushAllWorkers = (): Promise<AxiosResponse<any, any>> => {
	return getReq(`/workers/flush`, {})
}