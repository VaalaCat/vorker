import { GetNodeResponse, Node } from '@/types/nodes'
import api from './http'

export const getNodes = () => {
	return api.get<GetNodeResponse>('/api/node/all').then((res) => res.data)
}