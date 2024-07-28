import {
  Breadcrumb,
  Button,
  ButtonGroup,
  Divider,
  Input,
  Notification,
  Select,
  TabPane,
  Tabs,
  Toast,
  Typography,
} from '@douyinfe/semi-ui'
import { DEFAUTL_WORKER_ITEM, WorkerItem } from '@/types/workers'
import * as api from '@/api/workers'
import { useRouter } from 'next/router'
import { useMutation, useQuery } from '@tanstack/react-query'
import { $code, $vorkerSettings } from '@/store/workers'
import { useStore } from '@nanostores/react'
import { useEffect, useState } from 'react'
import { IconArticle, IconHome } from '@douyinfe/semi-icons'
import { getNodes } from '@/api/nodes'
import dynamic from 'next/dynamic'
import { i18n } from '@/lib/i18n'
import { TemplateEditor } from './template_editor'

const MonacoEditor = dynamic(
  import('./editor').then((m) => m.MonacoEditor),
  { ssr: false }
)

export const WorkerEditComponent = () => {
  const router = useRouter()
  const { UID } = router.query
  const [editItem, setEditItem] = useState(DEFAUTL_WORKER_ITEM)
  const [templateContent, setTemplateContent] = useState('')

  const appConf = useStore($vorkerSettings)
  const code = useStore($code)
  const { Paragraph, Text, Numeral, Title } = Typography
  const { data: resp } = useQuery(['getNodes'], () => getNodes())

  const { data: worker } = useQuery(['getWorker', UID], () => {
    return UID ? api.getWorker(UID as string) : null
  })

  const updateWorker = useMutation(async () => {
    await api.updateWorker(UID as string, editItem)
    Toast.info(i18n('workerSaveSuccess'))
  })

  const runWorker = useMutation(async (UID: string) => {
    let resp = await api.runWorker(UID)
    let raw_resp = JSON.stringify(resp)
    let run_resp = Buffer.from(resp?.data?.run_resp, 'base64').toString('utf8')
    let opts = {
      title: 'worker run result',
      content: (
        <>
          <Paragraph spacing="extended">
            <code className="overflow-scroll w-full">
              {run_resp.length > 100
                ? run_resp.slice(0, 100) + '......'
                : run_resp.length == 0
                  ? 'data is undefined, raw resp: ' + raw_resp
                  : run_resp}
            </code>
          </Paragraph>
          <div className="flex flex-row justify-end">
            <Text>copy to see full content</Text>
            <Paragraph
              copyable={{ content: run_resp }}
              spacing="extended"
              className="justify-end"
            />
          </div>
        </>
      ),
      duration: 10,
    }
    Notification.info({ ...opts, position: 'bottomRight' })
  })

  useEffect(() => {
    worker && setEditItem(worker)
  }, [UID, worker])

  useEffect(() => {
    if (worker) {
      setEditItem(worker)
      $code.set(Buffer.from(worker.Code, 'base64').toString('utf8'))
      if (worker.Template) setTemplateContent(worker.Template)
      else { setTemplateContent(DEFAUTL_WORKER_ITEM.Template) }
    }
  }, [worker])

  useEffect(() => {
    if (code && editItem)
      setEditItem((item) => ({
        ...item,
        Template: templateContent,
        Code: Buffer.from(code).toString('base64'),
      }))
  }, [code, editItem, templateContent])

  useEffect(() => {
    worker?.Code
  })

  const workerURL = `${appConf?.Scheme}://${editItem.Name}${appConf?.WorkerURLSuffix}`

  return (
    <div className="m-4 flex flex-col">
      <div className="flex flex-row justify-between">
        <div className="flex flex-col">
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
        </div>
        <div className="flex flex-col">
          <ButtonGroup>
            <Button
              onClick={() => {
                window.location.assign('/admin')
              }}
            >
              Back
            </Button>
            <Button onClick={() => updateWorker.mutate()}>Save</Button>
          </ButtonGroup>
        </div>
      </div>
      <div className="flex flex-row gap-1">
        <div className="columns-1 md:columns-2">
          <div></div>
          <Title heading={5}>ID</Title>
          <Paragraph copyable={{ content: editItem.UID }} spacing="extended">
            <code>{editItem.UID}</code>
          </Paragraph>
          <Title heading={5}>URL</Title>
          <Paragraph copyable={{ content: workerURL }} spacing="extended">
            <code>{workerURL}</code>
          </Paragraph>
        </div>
      </div>

      <Divider margin={4}></Divider>
      <Tabs
        tabBarExtraContent={
          <Button
            theme="borderless"
            onClick={() => runWorker.mutate(editItem.UID)}
          >
            Run
          </Button>
        }
      >
        <TabPane
          itemKey="code"
          style={{ overflow: 'initial' }}
          tab={<span>Code</span>}
        >
          {worker ? (
            <div className="flex flex-col my-1">
              <div>
                <MonacoEditor uid={worker.UID} />
              </div>
            </div>
          ) : null}
        </TabPane>
        <TabPane itemKey="config" tab={<span>Config</span>}>
          <div className="flex flex-col">
            <div className="flex flex-row m-2">
              <p className="self-center">Entry: </p>
              <div className="grid grid-cols-1 lg:grid-cols-2">
                <Input
                  addonBefore={
                    <p className="invisible w-0 sm:visible sm:w-auto">
                      {appConf?.Scheme}://
                    </p>
                  }
                  addonAfter={
                    <p className="invisible w-0 sm:visible w-25">
                      {appConf?.WorkerURLSuffix}
                    </p>
                  }
                  value={editItem.Name}
                  defaultValue={worker?.Name}
                  onChange={(value) => {
                    if (worker) {
                      setEditItem((item) => ({ ...item, Name: value }))
                    }
                  }}
                />
              </div>
            </div>
            <div className="flex flex-row m-2">
              <p className="self-center">Node: </p>
              <Select
                placeholder="请选择节点"
                style={{ width: 180 }}
                optionList={resp?.data.nodes.map((node) => {
                  return {
                    label: node.Name,
                    value: node.Name,
                  }
                })}
                value={editItem.NodeName}
                onChange={(value) => {
                  if (worker) {
                    setEditItem((item) => ({
                      ...item,
                      NodeName: value as string,
                    }))
                  }
                }}
              ></Select>
            </div>
          </div>
        </TabPane>
        <TabPane
          itemKey="template"
          style={{ overflow: 'initial' }}
          tab={<span>Template</span>}
        >
          <TemplateEditor content={templateContent} setContent={setTemplateContent} />
        </TabPane>
      </Tabs>
    </div>
  )
}
