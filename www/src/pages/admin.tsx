import { Layout } from '@/components/layout'
import { WorkersComponent } from '@/components/workers'
import { HeaderComponent } from '@/components/header'
import { SideBarComponent } from '@/components/sidebar'

export default function Admin() {
  return (
    <Layout
      header={<HeaderComponent />}
      side={<SideBarComponent selected='workers' />}
      main={<WorkersComponent />}
    />
  )
}
