import { HeaderComponent } from '@/components/header'
import { Layout } from '@/components/layout'
import { SideBarComponent } from '@/components/sidebar'
import { WorkerEditComponent } from '@/components/worker_edit'

export default function WorkerPage() {
  return (
    <Layout
      header={<HeaderComponent />}
      side={<SideBarComponent />}
      main={<WorkerEditComponent />}
    />
  )
}
