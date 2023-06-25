import React from 'react'
import * as monaco from 'monaco-editor'
import Editor, { loader } from '@monaco-editor/react'
import { useAtom } from 'jotai'
import { CodeAtom } from '@/store/workers'

loader.config({
  monaco,
})

export function MonacoEditor({ uid }: { uid: string }) {
  const [codeAtom, setCodeAtom] = useAtom(CodeAtom)
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
