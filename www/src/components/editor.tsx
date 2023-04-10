import React from 'react';
import ReactDOM from 'react-dom';

import Editor from '@monaco-editor/react';
import { useAtom } from 'jotai';
import { usernameAtom } from '../store/userState';

export function MonacoEditor() {
	const [username, setUsername] = useAtom(usernameAtom)
	return <div className='flex-1'>
		<Editor height="100vh" onChange={(v) => setUsername(v || '')}
			defaultLanguage="javascript" defaultValue={username} />
	</div>
}