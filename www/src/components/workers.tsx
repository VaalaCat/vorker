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
  Toast,
  Typography,
} from '@douyinfe/semi-ui'
import { useCallback, useEffect, useState } from 'react'
import { MonacoEditor } from './editor'
import { useMutation, useQuery } from '@tanstack/react-query'
import { useAtom } from 'jotai'
import { CodeAtom, VorkerSettingsAtom } from '@/store/workers'
import ColorHash from 'color-hash'
import { Router, useRouter } from 'next/router'
import {
  IconEdit,
  IconHome,
  IconMore,
  IconTreeTriangleDown,
  IconTreeTriangleRight,
} from '@douyinfe/semi-icons'
import { CH } from '@/lib/color'

export function WorkersComponent() {
  // get worker list
  const [code, setCodeAtom] = useAtom(CodeAtom)
  const [workerUID, setWorkerUID] = useState('')
  const [editItem, setEditItem] = useState(DEFAULT_WORKER_ITEM)
  const [appConfAtom] = useAtom(VorkerSettingsAtom)

  const router = useRouter()

  const { data: workers, refetch: reloadWorkers } = useQuery(
    ['getWorkers'],
    () => api.getAllWorkers()
  )

  const { data: worker } = useQuery(['getWorker', workerUID], () => {
    return workerUID ? api.getWorker(workerUID) : null
  })

  const createWorker = useMutation(async () => {
    await api.createWorker(DEFAULT_WORKER_ITEM)
    await reloadWorkers()
    Toast.info('创建成功！')
  })

  const deleteWorker = useMutation(async (uid: string) => {
    await api.deleteWorker(uid)
    await reloadWorkers()
    Toast.warning('删除成功！')
  })

  const flushWorker = useMutation(async () => {
    await api.flushWorker(workerUID)
    Toast.info('同步成功！')
  })

  const handleOpenWorker = useCallback(
    (item: WorkerItem) => {
      window.open(
        `${appConfAtom?.Scheme}://${item.Name}${appConfAtom?.WorkerURLSuffix}`,
        '_blank'
      )
    },
    [appConfAtom?.Scheme, appConfAtom?.WorkerURLSuffix]
  )

  useEffect(() => {
    reloadWorkers()
  }, [reloadWorkers, workerUID])

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
          <Button onClick={() => reloadWorkers()}>同步</Button>
          <Button onClick={() => createWorker.mutate()}>创建</Button>
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
                <span
                  className="text-base"
                  style={{ color: 'var(--semi-color-text-0)' }}
                >
                  {item.Name}
                </span>
                <p className="text-slate-400">{item.UID}</p>
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
                    编辑
                  </Button>
                  <Button
                    icon={<IconTreeTriangleRight />}
                    onClick={() => handleOpenWorker(item)}
                  >
                    运行
                  </Button>
                  <Dropdown
                    // onVisibleChange={(v) => handleVisibleChange(1, v)}
                    menu={[
                      {
                        node: 'item',
                        name: '删除',
                        onClick: () => deleteWorker.mutate(item.UID),
                      },
                      {
                        node: 'item',
                        name: '同步',
                        onClick: () => flushWorker.mutate(),
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
          <Typography>Worker Editor</Typography>
          <div></div>
          <div>
            <MonacoEditor uid={workerUID} />
          </div>
        </div>
      ) : null}
    </div>
  )
}
