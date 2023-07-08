import { Button, Nav } from '@douyinfe/semi-ui'
import {
  IconWrench,
  IconPercentage,
  IconHourglass,
  IconSidebar,
  IconMenu,
} from '@douyinfe/semi-icons'
import { useRouter } from 'next/router'
import { atom } from 'nanostores'
import { useStore } from '@nanostores/react'
import cn from 'clsx'
import { useEffect, useState } from 'react'

export const $expandSidebar = atom(false)

export const SideBarComponent = ({ selected }: { selected: string }) => {
  const router = useRouter()
  const expand = useStore($expandSidebar)
  const [invisible, setInvisible] = useState(expand)
  const routeMap = {
    workers: '/admin',
    status: '/nodes',
    settings: '/admin',
  } as any
  useEffect(() => {
    if (!expand) {
      setTimeout(() => setInvisible(!expand), 150)
    } else {
      setInvisible(!expand)
    }
  }, [expand])
  return (
    <>
      <Nav
        className={cn('z-10 fixed md:relative md:visible', { invisible })}
        style={{ height: '93vh' }}
        items={[
          { itemKey: 'workers', text: 'Workers', icon: <IconPercentage /> },
          { itemKey: 'status', text: 'Status', icon: <IconHourglass /> },
          // { itemKey: 'settings', text: 'Settings', icon: <IconWrench /> },
        ]}
        onSelect={(data) => console.log('trigger onSelect: ', data)}
        onClick={(data) => {
          window.location.assign(routeMap[data.itemKey || ''])
        }}
        footer={{ collapseButton: true }}
        isCollapsed={!expand}
        onCollapseChange={() => $expandSidebar.set(!expand)}
        defaultSelectedKeys={[selected]}
      />
    </>
  )
}
