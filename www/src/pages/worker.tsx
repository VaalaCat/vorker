import { HeaderComponent } from '@/components/header'
import { Layout } from '@/components/layout'
import { SideBarComponent } from '@/components/sidebar'
import { WorkerEditComponent } from '@/components/worker_edit'
import dynamic from 'next/dynamic'

export function WorkerPage() {
  return (
    <Layout
      header={<HeaderComponent />}
      side={<SideBarComponent selected="workers" />}
      main={<WorkerEditComponent />}
    />
  )
}
export default dynamic(() => Promise.resolve(WorkerPage), { ssr: false })
