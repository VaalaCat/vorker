import { Button, ButtonGroup, Divider, Input, TabPane, Tabs, Typography } from "@douyinfe/semi-ui"
import { MonacoEditor } from "./editor"
import { DEFAUTL_WORKER_ITEM, WorkerItem } from "@/types/workers"
import * as api from '@/api/workers'
import { useRouter } from "next/router"
import { useMutation, useQuery } from "@tanstack/react-query"
import { CodeAtom, VorkerSettingsAtom } from "@/store/workers"
import { useAtom } from "jotai"
import { useEffect, useState } from "react"
import { Base64 } from "js-base64"

export const WorkerEditComponent = () => {
	const router = useRouter();
	const { UID } = router.query
	const [editItem, setEditItem] = useState(DEFAUTL_WORKER_ITEM)
	const [appConfAtom] = useAtom(VorkerSettingsAtom)
	const [code, setCodeAtom] = useAtom(CodeAtom)
	const { Paragraph, Text, Numeral, Title } = Typography;

	const { data: worker } = useQuery(['getWorker', UID], () => {
		return UID ? api.GetWorker(UID as string) : null
	})

	const updateWorker = useMutation(async () => {
		await api.UpdateWorker(UID as string, editItem)
	})

	useEffect(() => {
		worker && setEditItem(worker)
	}, [UID])

	useEffect(() => {
		if (worker) {
			setEditItem(worker)
			setCodeAtom(Base64.decode(worker.Code))
		}
	}, [worker])

	useEffect(() => {
		if (code && editItem) setEditItem((item) => ({ ...item, Code: Base64.encode(code) }))
	}, [code])

	useEffect(() => {
		worker?.Code
	})

	return <div className="m-4 flex flex-col">
		<Title heading={5}>ID</Title>
		<Paragraph copyable spacing="extended">{editItem.UID}</Paragraph>
		<Title heading={5}>URL</Title>
		<Paragraph copyable spacing="extended">{`${appConfAtom?.Scheme}://${editItem.Name}${appConfAtom?.WorkerURLSuffix}`}</Paragraph>
		<Divider margin={4}></Divider>
		<Tabs tabBarExtraContent={
			<ButtonGroup>
				<Button onClick={() => updateWorker.mutate()}
					style={{
						justifySelf: 'flex-end',
					}}
				>保存</Button>
				<Button onClick={() => {
					router.push('/admin')
				}}>返回列表</Button>
			</ButtonGroup>
		}>
			<TabPane
				itemKey="code" tab={<span>代码</span>} >
				{worker ? (
					<div className="flex flex-col m-4" >
						<div></div>
						<div>
							<MonacoEditor uid={worker.UID} />
						</div>
					</div >
				) : null}
			</TabPane>
			<TabPane
				itemKey="config"
				tab={<span>配置</span>}
			>
				<div className="flex flex-row w-full">
					<Input addonBefore={`${appConfAtom?.Scheme}://`}
						addonAfter={`${appConfAtom?.WorkerURLSuffix}`}
						style={{ width: '30%' }}
						defaultValue={worker?.Name}
						onChange={(value) => {
							if (worker) {
								setEditItem((item) => ({ ...item, Name: value }))
							}
						}}
					/>
				</div>
			</TabPane>
		</Tabs>

	</div >
}