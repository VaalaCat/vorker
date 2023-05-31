import React from 'react'
import Editor from '@monaco-editor/react'
import { useAtom } from 'jotai'
import { CodeAtom } from '@/store/workers'
import { WorkerEditorProperties } from '@/types/workers'

export function MonacoEditor({ uid }: { uid: string }) {
  const [codeAtom, setCodeAtom] = useAtom(CodeAtom)
  return (
    <div className="flex-1">
      <Editor
        height="100vh"
        onChange={(v) => setCodeAtom(v || '')}
        value={codeAtom}
        defaultLanguage="javascript"
      />
    </div>
  )
}
