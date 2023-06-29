import { HeaderComponent } from '@/components/header'
import { Layout } from '@/components/layout'
import { LoginComponent } from '@/components/login'
import dynamic from 'next/dynamic'
import * as React from 'react'

export function SignIn() {
  return (
    <Layout
      header={<HeaderComponent />}
      side={<></>}
      main={<LoginComponent></LoginComponent>}
    ></Layout>
  )
}

export default dynamic(() => Promise.resolve(SignIn), { ssr: false })