export interface WorkerItem {
	UID: string;
	ExternalPath?: string;
	HostName?: string;
	NodeName?: string;
	Port?: number;
	Entry?: string;
	Code: string;
	Name: string;
}