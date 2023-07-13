import { HeaderComponent } from '@/components/header'
import { Layout } from '@/components/layout'
import { RegisterComponent } from '@/components/register'
import dynamic from 'next/dynamic'
import * as React from 'react'

export function Register() {
  return (
    <Layout
      header={<HeaderComponent />}
      side={<></>}
      main={<RegisterComponent></RegisterComponent>}
    ></Layout>
  )
}

export default dynamic(() => Promise.resolve(Register), { ssr: false })
