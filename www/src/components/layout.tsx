import { useAtom } from "jotai"
import { usernameAtom } from "@/store/userState"
import React, { useState } from "react"
import { Button } from "@mui/material"


export const Layout = ({ header, side, main }: { header: React.ReactNode, side: React.ReactNode, main: React.ReactNode }) => {
	return (
		<main className="flex flex-col">
			<div className="flex flex-row fixed h-7 mt-0">{header}</div>
			<div className='flex flex-row mt-7'>
				<div className='w-32 h-full flex-none'>{side}</div>
				{main}
			</div>
		</main>
	)
}