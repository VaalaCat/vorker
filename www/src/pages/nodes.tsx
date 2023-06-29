import { HeaderComponent } from '@/components/header'
import { Layout } from '@/components/layout'
import { NodesComponent } from '@/components/nodes'
import { SideBarComponent } from '@/components/sidebar'
import dynamic from 'next/dynamic'

export function Register() {
  return (
    <Layout
      header={<HeaderComponent />}
      side={<SideBarComponent selected="status" />}
      main={<NodesComponent />}
    ></Layout>
  )
}
export default dynamic(() => Promise.resolve(Register), { ssr: false })
