import { Layout } from '@/components/layout'
import { Nav } from '@douyinfe/semi-ui'
import { WorkersComponent } from '@/components/workers'
import { HeaderComponent } from '@/components/header'
import { IconWrench, IconPercentage } from '@douyinfe/semi-icons';

const SideBarComponent = () => {
  return (
    <Nav
      style={{ height: '93vh' }}
      items={[
        { itemKey: 'workers', text: 'Workers', icon: <IconPercentage /> },
        { itemKey: 'settings', text: 'Settings', icon: <IconWrench /> },
      ]}
      onSelect={(data) => console.log('trigger onSelect: ', data)}
      onClick={(data) => console.log('trigger onClick: ', data)}
      footer={{ collapseButton: true }}
      defaultSelectedKeys={['workers']}
    />
  )
}

export default function Admin() {
  return (
    <Layout
      header={<HeaderComponent />}
      side={<SideBarComponent />}
      main={<WorkersComponent />}
    />
  )
}
