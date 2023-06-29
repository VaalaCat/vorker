import { getNodes } from '@/api/nodes'
import {
  Node,
  PingMap,
  PingMapList,
  Tracker as TrackerType,
} from '@/types/nodes'
import { useMutation, useQuery } from '@tanstack/react-query'
import { useEffect, useRef, useState } from 'react'
import { Card, Title, Tracker, Flex, Text, Color } from '@tremor/react'
import { Button } from '@douyinfe/semi-ui'
import * as nodesApi from '@/api/nodes'
import { useStore } from '@nanostores/react'
import { $nodeStatus } from '@/store/nodes'

export function NodesComponent() {
  const [nodelist, setNodelist] = useState<Node[]>([])
  const [rerenderID, rerender] = useState(0)
  const { data: resp, refetch: reloadNodes } = useQuery(
    ['getNodes'],
    async () => {
      const res = await getNodes()
      rerender(Math.random())
      return res
    },
    {
      refetchInterval: 5000,
    }
  )
  const nodeStatus = useStore($nodeStatus)

  useEffect(() => {
    setNodelist(resp?.data.nodes || [])
    if (resp?.data) {
      const v = Object.entries(resp.data.ping).map(([k, v]) => {
        let t = $nodeStatus.get()[k] || Array.from({ length: 50 }, () => -1)
        if (t.length > 50) {
          t.shift()
        }
        let s = [k, [...t, resp?.data.ping[k] || 0]]
        return s
      })
      const a = Object.fromEntries(v) as PingMapList
      $nodeStatus.set(a)
    }
  }, [reloadNodes, resp?.data, rerenderID])

  return (
    <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3">
      {nodelist?.map((node) => (
        <div key={node.UID}>
          <NodeComponent node={node} ping={nodeStatus[node.Name]} />
        </div>
      ))}
    </div>
  )
}

export function NodeComponent({ node, ping }: { node: Node; ping: number[] }) {
  const syncNodes = useMutation(async (nodeName: string) => {
    await nodesApi.syncNodes(nodeName)
  })

  const data = ping.map((v, i) => {
    if (v >= 1000) {
      return { color: 'rose', tooltip: `${v}ms` }
    }
    if (v >= 100) {
      return { color: 'yellow', tooltip: `${v}ms` }
    }
    if (v === -1) {
      return { color: 'gray', tooltip: 'N/A' }
    }
    return { color: 'emerald', tooltip: `${v}ms` }
  }) as TrackerType[]

  const sla = parseFloat(
    ((1 - ping.filter((v) => v >= 500).length / ping.length) * 100).toFixed(2)
  )
  const validValue = ping.filter((v) => v !== -1)
  const avg = parseFloat(
    (validValue.reduce((a, b) => a + b, 0) / validValue.length).toFixed(2)
  )

  return (
    <div className="m-2">
      <Card>
        <div className="flex flex-row justify-between">
          <Title>{node.Name}</Title>
          <Button
            theme="borderless"
            className="justify-end"
            onClick={() => syncNodes.mutate(node.Name)}
          >
            {' '}
            Sync{' '}
          </Button>
        </div>
        <Text>ID {node.UID}</Text>
        <Flex justifyContent="end" className="mt-4">
          <Text>Uptime {sla}%</Text>
          <Text className="ml-2">Avg. {avg}ms</Text>
        </Flex>
        <Tracker data={data} className="mt-2" />
      </Card>
    </div>
  )
}
