import { Layout } from '@/components/layout'
import { Avatar, Button, ButtonGroup, Nav, TabPane } from '@douyinfe/semi-ui'
import EngineeringIcon from '@mui/icons-material/Engineering'
import SettingsIcon from '@mui/icons-material/Settings'
import { WorkerComponent } from '@/components/workers'
import { LuFunctionSquare } from 'react-icons/lu'

const sideButtons = [
  <Button key="Settings" icon={<SettingsIcon />}>
    Settings
  </Button>,
]

const SideBarComponent = () => {
  return (
    <Nav
      style={{ height: '90vh' }}
      items={[
        { itemKey: 'workers', text: 'Workers', icon: <EngineeringIcon /> },
        { itemKey: 'settings', text: 'Settings', icon: <SettingsIcon /> },
      ]}
      onSelect={(data) => console.log('trigger onSelect: ', data)}
      onClick={(data) => console.log('trigger onClick: ', data)}
    />
  )
}

const HeaderComponent = () => {
  return (
    <Nav mode="horizontal" defaultSelectedKeys={['Home']}>
      <Nav.Header>
        <LuFunctionSquare color="#7f7f7f" style={{ fontSize: 36 }} />
        <span className="text-xl ml-2" style={{ fontFamily: 'trebuchet ms' }}>
          Vorker
        </span>
      </Nav.Header>
      <Nav.Footer>
        <ButtonGroup aria-label="header button">
          <Button type="primary" theme="borderless">
            Home
          </Button>
          <Button type="primary" theme="borderless">
            Admin
          </Button>
        </ButtonGroup>
      </Nav.Footer>
    </Nav>
  )
}

export default function Admin() {
  return (
    <Layout
      header={<HeaderComponent />}
      side={<SideBarComponent />}
      main={<WorkerComponent />}
    />
  )
}
