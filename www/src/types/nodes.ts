import { Color } from "@tremor/react";

export interface Node {
	Name: string
	UID: string
}

export interface PingMap {
	[key: string]: number
}

export interface GetNodeResponse {
	code: number
	msg: string
	data: {
		nodes: Node[]
		ping: PingMap
	}
}

export interface Tracker {
	color: Color;
	tooltip: string;
}

export interface PingMapList {
	[key: string]: number[];
}