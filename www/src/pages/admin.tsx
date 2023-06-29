'use client'
import { Layout } from '@/components/layout'
import { WorkersComponent } from '@/components/workers'
import { HeaderComponent } from '@/components/header'
import { SideBarComponent } from '@/components/sidebar'
import dynamic from 'next/dynamic'

export function Admin() {
  return (
    <Layout
      header={<HeaderComponent />}
      side={<SideBarComponent selected="workers" />}
      main={<WorkersComponent />}
    />
  )
}

export default dynamic(() => Promise.resolve(Admin), { ssr: false })
