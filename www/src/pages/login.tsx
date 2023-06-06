import { HeaderComponent } from '@/components/header'
import { Layout } from '@/components/layout'
import { LoginComponent } from '@/components/login'
import * as React from 'react'

export default function SignIn() {
  return (
    <Layout
      header={<HeaderComponent />}
      side={<></>}
      main={<LoginComponent></LoginComponent>}
    ></Layout>
  )
}
