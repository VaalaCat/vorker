import React from 'react'
import Editor from '@monaco-editor/react'
import { useAtom } from 'jotai'
import { CodeAtom } from '@/store/workers'

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
