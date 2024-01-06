import * as api from '@/api/workers'
import {
  DEFAUTL_WORKER_ITEM as DEFAULT_WORKER_ITEM,
  WorkerItem,
} from '@/types/workers'
import {
  Avatar,
  Breadcrumb,
  Button,
  ButtonGroup,
  Dropdown,
  List,
  Modal,
  Notification,
  Tag,
  Toast,
  Typography,
} from '@douyinfe/semi-ui'
import { useCallback, useEffect, useState } from 'react'
import { useMutation, useQuery } from '@tanstack/react-query'
import { useStore } from '@nanostores/react'
import { $code, $vorkerSettings } from '@/store/workers'
import ColorHash from 'color-hash'
import { Router, useRouter } from 'next/router'
import {
  IconEdit,
  IconHome,
  IconLink,
  IconMore,
  IconTreeTriangleDown,
  IconTreeTriangleRight,
} from '@douyinfe/semi-icons'
import { CH } from '@/lib/color'
import { it } from 'node:test'
import dynamic from 'next/dynamic'
import { t } from '@/lib/i18n'

const MonacoEditor = dynamic(
  import('./editor').then((m) => m.MonacoEditor),
  { ssr: false }
)

export function WorkersComponent() {
  // get worker list
  const code = useStore($code)
  const [workerUID, setWorkerUID] = useState('')
  const [editItem, setEditItem] = useState(DEFAULT_WORKER_ITEM)
  const { Paragraph, Text, Numeral, Title } = Typography
  const appConf = useStore($vorkerSettings)

  const router = useRouter()

  const { data: workers, refetch: reloadWorkers } = useQuery(
    ['getWorkers'],
    () => api.getAllWorkers()
  )

  const { data: worker } = useQuery(['getWorker', workerUID], () => {
    return workerUID ? api.getWorker(workerUID) : null
  })

  const createWorker = useMutation(async () => {
    await api.createWorker({ ...DEFAULT_WORKER_ITEM })
    await reloadWorkers()
    Toast.info(t.workerCreateSuccess)
  })

  const deleteWorker = useMutation(async (uid: string) => {
    await api.deleteWorker(uid)
    await reloadWorkers()
    Toast.warning(t.workerDeleteSuccess)
  })

  const flushWorker = useMutation(async (uid: string) => {
    await api.flushWorker(uid)
    Toast.info(t.workerSyncSuccess)
  })

  const flushAllWorkers = useMutation(async () => {
    await api.flushAllWorkers()
    Toast.info(t.workerSyncSuccess)
  })

  const handleOpenWorker = useCallback(
    (item: WorkerItem) => {
      window.open(
        `${appConf?.Scheme}://${item.Name}${appConf?.WorkerURLSuffix}`,
        '_blank'
      )
    },
    [appConf?.Scheme, appConf?.WorkerURLSuffix]
  )

  const handleDeleteWorker = useCallback(
    (item: WorkerItem) => {
      Modal.warning({
        title: t.deleteWorker,
        content: (
          <span className="break-all">
            确定要删除 {item.Name} (ID: <code>{item.UID}</code>) 吗
          </span>
        ),
        centered: true,
        onOk: () => deleteWorker.mutate(item.UID),
      })
    },
    [deleteWorker]
  )

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
    reloadWorkers()
  }, [reloadWorkers, workerUID])

  useEffect(() => {
    if (worker) {
      setEditItem(worker)
      $code.set(Buffer.from(worker.Code, 'base64').toString('utf8'))
    }
  }, [worker])

  useEffect(() => {
    if (code && editItem)
      setEditItem((item) => ({
        ...item,
        Code: Buffer.from(code).toString('base64'),
      }))
  }, [code, editItem])

  return (
    <div className="m-4">
      <div className="flex justify-between">
        <Breadcrumb>
          <Breadcrumb.Item
            href="/admin"
            icon={<IconHome size="small" />}
          ></Breadcrumb.Item>
          <Breadcrumb.Item href="/admin">Workers</Breadcrumb.Item>
        </Breadcrumb>
        <ButtonGroup>
          <Button onClick={() => reloadWorkers()}>{t.refresh}</Button>
          <Button onClick={() => flushAllWorkers.mutate()}>{t.sync}</Button>
          <Button onClick={() => createWorker.mutate()}>{t.create}</Button>
        </ButtonGroup>
      </div>
      <List
        dataSource={workers}
        renderItem={(item) => (
          <List.Item
            header={
              <Avatar
                shape="square"
                style={{ background: CH.hex(item.UID) }}
                className="pointer-events-none"
              >
                {item.Name.slice(0, 2).toUpperCase()}
              </Avatar>
            }
            main={
              <div className="flex flex-col justify-between h-12">
                <div className="flex flex-row w-full">
                  <span
                    className="text-base"
                    style={{ color: 'var(--semi-color-text-0)' }}
                  >
                    {item.Name}
                  </span>
                </div>
                <p className="text-slate-400">
                  Node: <Tag>{item.NodeName} </Tag>
                </p>
              </div>
            }
            extra={
              <ButtonGroup theme="borderless">
                <ButtonGroup theme="borderless">
                  <Button
                    icon={<IconEdit />}
                    onClick={() => {
                      router.push({
                        pathname: '/worker',
                        query: { UID: item.UID },
                      })
                    }}
                  >
                    {t.edit}
                  </Button>
                  <Button
                    icon={<IconTreeTriangleRight />}
                    // onClick={() => handleOpenWorker(item)}
                    onClick={() => runWorker.mutate(item.UID)}
                  >
                    {t.run}
                  </Button>
                  <Button
                    icon={<IconLink />}
                    onClick={() => handleOpenWorker(item)}
                  >
                    {t.open}
                  </Button>
                  <Dropdown
                    // onVisibleChange={(v) => handleVisibleChange(1, v)}
                    menu={[
                      {
                        node: 'item',
                        name: t.delete,
                        onClick: () => handleDeleteWorker(item),
                      },
                      {
                        node: 'item',
                        name: t.sync,
                        onClick: () => flushWorker.mutate(item.UID),
                      },
                    ]}
                    trigger="click"
                    position="bottomRight"
                  >
                    <Button theme="borderless" icon={<IconMore />}></Button>
                  </Dropdown>
                </ButtonGroup>
              </ButtonGroup>
            }
          />
        )}
      />
      {workerUID ? (
        <div className="flex flex-col w-full m-4">
          <Typography>{t.editor}</Typography>
          <div></div>
          <div>
            <MonacoEditor uid={workerUID} />
          </div>
        </div>
      ) : null}
    </div>
  )
}
