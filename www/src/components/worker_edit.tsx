import {
  Breadcrumb,
  Button,
  ButtonGroup,
  Divider,
  Input,
  TabPane,
  Tabs,
  Toast,
  Typography,
} from '@douyinfe/semi-ui'
import { MonacoEditor } from './editor'
import { DEFAUTL_WORKER_ITEM, WorkerItem } from '@/types/workers'
import * as api from '@/api/workers'
import { useRouter } from 'next/router'
import { useMutation, useQuery } from '@tanstack/react-query'
import { CodeAtom, VorkerSettingsAtom } from '@/store/workers'
import { useAtom } from 'jotai'
import { useEffect, useState } from 'react'
import { IconArticle, IconHome } from '@douyinfe/semi-icons'

export const WorkerEditComponent = () => {
  const router = useRouter()
  const { UID } = router.query
  const [editItem, setEditItem] = useState(DEFAUTL_WORKER_ITEM)
  const [appConfAtom] = useAtom(VorkerSettingsAtom)
  const [code, setCodeAtom] = useAtom(CodeAtom)
  const { Paragraph, Text, Numeral, Title } = Typography

  const { data: worker } = useQuery(['getWorker', UID], () => {
    return UID ? api.getWorker(UID as string) : null
  })

  const updateWorker = useMutation(async () => {
    await api.updateWorker(UID as string, editItem)
    Toast.info('保存成功！')
  })

  useEffect(() => {
    worker && setEditItem(worker)
  }, [UID, worker])

  useEffect(() => {
    if (worker) {
      setEditItem(worker)
      setCodeAtom(Buffer.from(worker.Code, 'base64').toString('utf8'))
    }
  }, [setCodeAtom, worker])

  useEffect(() => {
    if (code && editItem)
      setEditItem((item) => ({
        ...item,
        Code: Buffer.from(code).toString('base64'),
      }))
  }, [code, editItem])

  useEffect(() => {
    worker?.Code
  })

  const workerURL = `${appConfAtom?.Scheme}://${editItem.Name}${appConfAtom?.WorkerURLSuffix}`

  return (
    <div className="m-4 flex flex-col">
      <div className="flex justify-between">
        <div className="flex flex-col gap-1">
          <Breadcrumb compact={false}>
            <Breadcrumb.Item
              href="/admin"
              icon={<IconHome size="small" />}
            ></Breadcrumb.Item>
            <Breadcrumb.Item href="/admin">Workers</Breadcrumb.Item>
            <Breadcrumb.Item href={`/worker?UID=${editItem.UID}`}>
              {editItem.Name}
            </Breadcrumb.Item>
          </Breadcrumb>
          <Title heading={5}>ID</Title>
          <Paragraph copyable={{ content: editItem.UID }} spacing="extended">
            <code>{editItem.UID}</code>
          </Paragraph>
          <Title heading={5}>URL</Title>
          <Paragraph copyable={{ content: workerURL }} spacing="extended">
            <code>{workerURL}</code>
          </Paragraph>
        </div>
        <div>
          <ButtonGroup>
            <Button onClick={() => updateWorker.mutate()}>保存</Button>
            <Button onClick={() => router.push('/admin')}>返回列表</Button>
          </ButtonGroup>
        </div>
      </div>

      <Divider margin={4}></Divider>
      <Tabs>
        <TabPane itemKey="code" tab={<span>代码</span>}>
          {worker ? (
            <div className="flex flex-col m-4">
              <div>
                <MonacoEditor uid={worker.UID} />
              </div>
            </div>
          ) : null}
        </TabPane>
        <TabPane itemKey="config" tab={<span>配置</span>}>
          <div className="flex flex-row w-full">
            <Input
              addonBefore={`${appConfAtom?.Scheme}://`}
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
    </div>
  )
}
