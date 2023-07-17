import { Layout } from '@/components/layout'
import { HeaderComponent } from '@/components/header'
import { Inter } from 'next/font/google'
import clsx from 'clsx'
import { Button } from '@douyinfe/semi-ui'
import { useState } from 'react'

const inter = Inter({ subsets: ['latin'] })

const FeatureCard = ({
  icon,
  title,
  desc,
}: {
  icon?: React.ReactNode
  title: string
  desc: React.ReactNode
}) => {
  return (
    <div className="p-4 border-solid border-2 rounded w-full">
      <h3 className="text-lg">
        {icon}
        {title}
      </h3>
      <div>{desc}</div>
    </div>
  )
}

const code = `addEventListener("fetch", (event) => {
	event.respondWith(handler(event));
});

async function handler(event) {
	try {
		let resp = new Response("hello vorker!")
		return resp
	} catch(e) {
		return new Response(e.stack, { status: 500 })
	}
}`

const HomePage = () => {
  const [showResult, setShowResult] = useState(false)

  return (
    <div className={clsx('flex flex-col items-center', inter.className)}>
      <div className="mt-16 text-center">
        <h1 className="text-6xl">Vorker Functions</h1>
        <h3 className="text-2xl">
          Vorker is a lightweight functions platform.
        </h3>
      </div>

      <div className="mt-16 flex flex-col items-center">
        <pre
          className="m-4 p-4 rounded overflow-auto text-sm"
          style={{
            width: '80%',
            minWidth: '480px',
            border: '1px solid #e9ebec',
            background: '#f5f7f8',
          }}
        >
          <code>{code}</code>
        </pre>
        <Button onClick={() => setShowResult(true)}>Run</Button>
      </div>

      {showResult && (
        <div className="mt-8">
          <pre>
            <code>hello vorker!</code>
          </pre>
        </div>
      )}

      <div className="mt-20 features grid grid-cols-1 md:grid-cols-2 gap-4">
        <FeatureCard title="Fast" desc="super fast powered by workerd" />
        <FeatureCard title="Open Source" desc="under ? license" />
        <FeatureCard title="Self Hosted" desc="host on your own machine" />
      </div>
    </div>
  )
}

export default function Home() {
  return (
    <Layout header={<HeaderComponent hasSider={false} />} main={<HomePage />} />
  )
}
