'use client'
import { useRouter } from 'next/navigation'
import { useState } from 'react'

export default function Page() {
	const router = useRouter()
	const [token, setToken] = useState('')
	const [valid, setValid] = useState<null | boolean>(null)

	return (
		<div className="flex h-screen w-full items-center justify-center gap-4 bg-gray-100">
			<form
				onSubmit={e => {
					e.preventDefault()
					fetch(`/verify?token=${token}`)
						.then(res => res.json())
						.then(data => {
							if (data.error) {
								setValid(false)
							} else {
								router.push('/')
							}
						})
				}}
				className="flex w-80 flex-col items-center justify-center gap-4 rounded-xl bg-white p-4 shadow-lg"
			>
				<h1 className="w-full text-center text-2xl">Token</h1>
				<input
					type="text"
					value={token}
					onChange={e => setToken(e.target.value)}
					className="w-full rounded border-2 border-black p-4"
				/>
				{valid === false && <p>Invalid token</p>}

				<button type="submit" className="w-full rounded border-2 border-black p-4">
					Check
				</button>
			</form>
		</div>
	)
}
