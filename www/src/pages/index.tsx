import { Inter } from 'next/font/google'
import { MonacoEditor } from '../components/editor';
import { Layout } from '../components/layout';

const inter = Inter({ subsets: ['latin'] })

export default function Home() {
  return (
    <Layout><MonacoEditor /></Layout>
  )
}
