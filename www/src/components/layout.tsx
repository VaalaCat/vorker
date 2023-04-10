import { useAtom } from "jotai"
import { usernameAtom } from "@/store/userState"
import { useState } from "react"


export const Layout = ({ children }: React.PropsWithChildren<{}>) => {
	const [username, setUsername] = useAtom(usernameAtom)
	const [show, setShow] = useState(false)
	return (
		<main>
			<div className='flex flex-row'>
				<div className='w-32 h-full flex-none' style={{ animation: 'test 1s ease-in-out' }}>
					<p>
						<span>{username}</span>
					</p>
					{show && <div>111</div>}
				</div>
				{children}
			</div>
		</main>
	)
}