import * as api from '@/api/workers'
import { DEFAUTL_WORKER_ITEM } from '@/types/workers'
import { WorkerItem, WorkerItemProperties } from '@/types/workers'
import {
  Avatar,
  Button,
  ButtonGroup,
  Card,
  List,
  Typography,
} from '@douyinfe/semi-ui'
import { useEffect, useState } from 'react'
import { MonacoEditor } from './editor'
import { useMutation, useQuery } from '@tanstack/react-query'
import { useAtom } from 'jotai'
import { CodeAtom } from '@/store/workers'
// @ts-ignore
import ColorHash from 'color-hash'

const CH = new ColorHash()

export function WorkersComponent() {
  // get workerlist list
  const [code, setCodeAtom] = useAtom(CodeAtom)
  const [workerUID, setWorkerUID] = useState('')
  const [editItem, setEditItem] = useState(DEFAUTL_WORKER_ITEM)

  const { data: workers, refetch: reloadWorkers } = useQuery(
    ['getWorkers'],
    () => api.GetAllWorkers()
  )
  const { data: worker } = useQuery(['getWorker', workerUID], () => {
    return workerUID ? api.GetWorker(workerUID) : null
  })
  const createWorker = useMutation(async () => {
    await api.CreateWorker(DEFAUTL_WORKER_ITEM)
    await reloadWorkers()
  })

  const deleteWorker = useMutation(async (uid: string) => {
    await api.DeleteWorker(uid)
    await reloadWorkers()
  })

  const flushWorker = useMutation(async () => {
    await api.FlushWorker(workerUID)
  })

  const updateWorker = useMutation(async () => {
    await api.UpdateWorker(workerUID, editItem)
  })

  useEffect(() => {
    console.log('reload workers')
    reloadWorkers()
  }, [workerUID])
  useEffect(() => {
    if (worker) {
      console.log('update item and code', atob(worker.Code))
      setEditItem(worker)
      setCodeAtom(atob(worker.Code))
    }
  }, [worker])
  useEffect(() => {
    console.log('update item code', code)
    if (code && editItem) setEditItem((item) => ({ ...item, Code: btoa(code) }))
  }, [code])

  return (
    <div className="w-full m-4 h-full">
      <Button onClick={() => reloadWorkers()}>刷新</Button>
      <Button onClick={() => createWorker.mutate()}>创建</Button>
      <Button onClick={() => setWorkerUID('')}>返回</Button>
      <Button onClick={() => updateWorker.mutate()}>保存</Button>
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
                      setWorkerUID(item.UID)
                    }}
                  >
                    Edit
                  </Button>
                  <Button onClick={() => deleteWorker.mutate(item.UID)}>
                    Delete
                  </Button>
                  <Button onClick={() => flushWorker.mutate}>Flush</Button>
                </ButtonGroup>
              </ButtonGroup>
            }
          />
        )}
      />
      {workerUID ? (
        <div className="flex flex-col w-full m-4 h-full">
          <Typography>Worker Editor</Typography>
          <div></div>
          <div className='h-full'>
            <MonacoEditor uid={workerUID} />
          </div>
        </div>
      ) : null}
    </div>
  )
}
