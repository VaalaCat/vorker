import React from 'react'
import { Layout as SemiLayout } from '@douyinfe/semi-ui'
import { QueryClient, QueryClientProvider } from '@tanstack/react-query'

const queryClient = new QueryClient()
const { Header, Footer, Sider, Content } = SemiLayout

export const Layout = ({
  header,
  side,
  main,
}: {
  header: React.ReactNode
  side?: React.ReactNode
  main: React.ReactNode
}) => {
  return (
    <QueryClientProvider client={queryClient}>
      <SemiLayout>
        <Header>{header}</Header>
        <SemiLayout className="relative">
          {side && <Sider className="fixed md:relative">{side}</Sider>}
          <Content
            className="flex-1 overflow-scroll"
            style={{ height: '93vh' }}
          >
            {main}
          </Content>
        </SemiLayout>
      </SemiLayout>
    </QueryClientProvider>
  )
}
