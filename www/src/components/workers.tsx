import * as api from '@/api/workers'
import { DEFAUTL_WORKER_ITEM as DEFAULT_WORKER_ITEM } from '@/types/workers'
import {
  Avatar,
  Button,
  ButtonGroup,
  List,
  Toast,
  Typography,
} from '@douyinfe/semi-ui'
import { useEffect, useState } from 'react'
import { MonacoEditor } from './editor'
import { useMutation, useQuery } from '@tanstack/react-query'
import { useAtom } from 'jotai'
import { CodeAtom, VorkerSettingsAtom } from '@/store/workers'
import ColorHash from 'color-hash'
import { Router, useRouter } from 'next/router'

const CH = new ColorHash()

export function WorkersComponent() {
  // get worker list
  const [code, setCodeAtom] = useAtom(CodeAtom)
  const [workerUID, setWorkerUID] = useState('')
  const [editItem, setEditItem] = useState(DEFAULT_WORKER_ITEM)
  const [appConfAtom] = useAtom(VorkerSettingsAtom)

  const router = useRouter()

  const { data: workers, refetch: reloadWorkers } = useQuery(
    ['getWorkers'],
    () => api.GetAllWorkers()
  )

  const { data: worker } = useQuery(['getWorker', workerUID], () => {
    return workerUID ? api.GetWorker(workerUID) : null
  })

  const createWorker = useMutation(async () => {
    await api.CreateWorker(DEFAULT_WORKER_ITEM)
    await reloadWorkers()
    Toast.info('创建成功！')
  })

  const deleteWorker = useMutation(async (uid: string) => {
    await api.DeleteWorker(uid)
    await reloadWorkers()
    Toast.warning('删除成功！')
  })

  const flushWorker = useMutation(async () => {
    await api.FlushWorker(workerUID)
    Toast.info('刷新成功！')
  })

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
    <div className="w-full m-4">
      <Button onClick={() => reloadWorkers()}>刷新</Button>
      <Button onClick={() => createWorker.mutate()}>创建</Button>
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
              <div>
                <span
                  style={{ color: 'var(--semi-color-text-0)', fontWeight: 500 }}
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
                    onClick={() => {
                      router.push({
                        pathname: '/worker',
                        query: { UID: item.UID },
                      })
                    }}
                  >
                    编辑
                  </Button>
                  <Button onClick={() => deleteWorker.mutate(item.UID)}>
                    删除
                  </Button>
                  <Button onClick={() => flushWorker.mutate}>刷新</Button>
                  <Button
                    onClick={() => {
                      window.open(
                        `${appConfAtom?.Scheme}://${item.Name}${appConfAtom?.WorkerURLSuffix}`,
                        '_blank'
                      )
                    }}
                  >
                    运行
                  </Button>
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
