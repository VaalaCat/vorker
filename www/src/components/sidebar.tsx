import { Nav } from '@douyinfe/semi-ui'
import { IconWrench, IconPercentage, IconHourglass } from '@douyinfe/semi-icons'
import { useRouter } from 'next/router';

export const SideBarComponent = ({ selected }: { selected: string }) => {
  const router = useRouter()
  const routeMap = {
    'workers': '/admin',
    'status': '/nodes',
    'settings': '/admin',
  } as any;
  return (
    <Nav
      style={{ height: '93vh' }}
      items={[
        { itemKey: 'workers', text: 'Workers', icon: <IconPercentage /> },
        { itemKey: 'status', text: 'Status', icon: < IconHourglass /> },
        // { itemKey: 'settings', text: 'Settings', icon: <IconWrench /> },
      ]}
      onSelect={(data) => console.log('trigger onSelect: ', data)}
      onClick={(data) => { window.location.assign(routeMap[data.itemKey || ""]) }}
      footer={{ collapseButton: true }}
      defaultIsCollapsed={true}
      defaultSelectedKeys={[selected]}
    />
  )
}
