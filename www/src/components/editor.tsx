import React, { useLayoutEffect } from 'react'
import Editor, { loader } from '@monaco-editor/react'
import { useAtom } from 'jotai'
import { CodeAtom } from '@/store/workers'
import dynamic from 'next/dynamic'

loader.config({
  paths: {
    vs: 'https://fastly.jsdelivr.net/npm/monaco-editor@0.36.1/min/vs',
  },
})

export function MonacoEditor({ uid }: { uid: string }) {
  const [codeAtom, setCodeAtom] = useAtom(CodeAtom)
  useLayoutEffect(() => {}, [])
  return (
    <div className="flex-1">
      <Editor
        height="60vh"
        onChange={(v) => setCodeAtom(v || '')}
        value={codeAtom}
        defaultLanguage="javascript"
      />
    </div>
  )
}
