export interface WorkerItem {
  UID: string
  ExternalPath?: string
  HostName?: string
  NodeName?: string
  Port?: number
  Entry?: string
  Code: string
  Name: string
}

export interface WorkerItemProperties {
  item: WorkerItem
}

// name is generated on server side
// @ts-expect-error
export const DEFAUTL_WORKER_ITEM: WorkerItem = {
  UID: 'worker',
  Code: btoa(`addEventListener("fetch", event => {
	event.respondWith(new Response("Hello World"));
});`),
}

export interface WorkerEditorProperties {
  item: string
}

export interface VorkerSettingsProperties {
  WorkerURLSuffix: string
  Scheme: string
  EnableRegister: boolean
}
