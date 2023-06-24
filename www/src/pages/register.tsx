import { HeaderComponent } from '@/components/header'
import { Layout } from '@/components/layout'
import { RegisterComponent } from '@/components/register'
import * as React from 'react'

export default function Register() {
  return (
    <Layout
      header={<HeaderComponent />}
      side={<></>}
      main={<RegisterComponent></RegisterComponent>}
    ></Layout>
  )
}
