import { CreateWorker, DeleteWorker, FlushWorker, GetWorker, GetWorkers, UpdateWorker } from "@/api/workers";
import { DEFAUTL_WORKER_ITEM } from "@/types/workers";
import { WorkerItem, WorkerItemProperties } from "@/types/workers";
import { Alert, AlertTitle, Box, Button, ButtonGroup, Card, Stack, Typography } from "@mui/material";
import { useEffect, useState } from "react";
import { MonacoEditor } from "./editor";
import { useAtom } from "jotai";
import { CodeAtom } from "@/store/workers";

export function WorkerComponent() {
	// get workerlist list
	const [items, setItems] = useState<WorkerItem[] | undefined>([])
	const [alert, setAlert] = useState(false)
	const [alertTitle, setAlertTitle] = useState("")
	const [editItemUID, setEditItemUID] = useState("")
	const [editItem, setEditItem] = useState(DEFAUTL_WORKER_ITEM)
	const [codeAtom, setCodeAtom] = useAtom(CodeAtom)

	const getWorker = (UID: string) => {
		GetWorker(UID)
			.then(resp => {
				setEditItem(resp?.data.data)
			})
			.catch(err => {
				console.log(err)
			})
	}

	const getWorkers = () => {
		GetWorkers(0, 10).then((resp) => {
			setItems((resp.data.data as WorkerItem[]));
		})
	}

	const createWorker = () => {
		CreateWorker(DEFAUTL_WORKER_ITEM).then((resp) => {
			console.log(`createwokers: ${JSON.stringify(resp)}`);
		}).then(() => { getWorkers() })
	}

	const deleteWorker = (uid: string) => {
		DeleteWorker(uid).then(() => {
			setAlertTitle('Worker deleted')
			setAlert(true);
			setTimeout(() => setAlert(false), 3000);
		}).then(() => { getWorkers() })
	}

	const flushWorker = (uid: string) => {
		FlushWorker(uid).then(() => {
			setAlertTitle('Worker flushed');
			setAlert(true);
			setTimeout(() => setAlert(false), 3000);
		})
	}

	const updateWorker = (uid: string, workerInfo: WorkerItem) => {
		UpdateWorker(uid, workerInfo).then(() => {
			setAlertTitle('Worker updated');
			setAlert(true);
			setTimeout(() => setAlert(false), 3000);
		})
	}

	function WorkerItemComponent({ item }: WorkerItemProperties) {
		return <div className="flex flex-row w-full p-2">
			<div className="flex flex-col">
				<div><Typography>{item.Name}</Typography></div>
			</div>
			<div className="ml-auto">
				<ButtonGroup variant="outlined" orientation="horizontal" size="small">
					<Button onClick={() => { setEditItemUID(item.UID); getWorker(item.UID) }}>Edit</Button>
					<Button onClick={() => deleteWorker(item.UID)}>Delete</Button>
					<Button onClick={() => flushWorker(item.UID)}>Flush</Button>
				</ButtonGroup>
			</div>
		</div>
	}

	function WorkerItemEditComponent() {
		return <div className="flex flex-col w-full m-4 h-full">
			<Typography>Worker Editor</Typography>
			<Button onClick={() => setEditItemUID("")}>返回</Button>
			<div>
				<MonacoEditor />
			</div>
		</div>
	}

	if (editItemUID !== "") {
		return <WorkerItemEditComponent />
	}

	return <div className="w-full m-4">
		<Button onClick={getWorkers}>刷新</Button>
		<Button onClick={createWorker}>创建</Button>
		{alert && <Alert severity="info">
			<AlertTitle>Info</AlertTitle>{alertTitle}
		</Alert>}
		{items?.map((worker) => (
			<div className="my-2">
				<Card variant="outlined">
					<WorkerItemComponent item={worker} />
				</Card>
			</div>
		))}
	</div>
}