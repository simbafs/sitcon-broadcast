'use client'
import { useRouter } from 'next/navigation'
import { useQueryState } from 'nuqs'
import { FormEvent, Suspense, useState } from 'react'
import { twMerge } from 'tailwind-merge'

function Verify() {
	const [token, setToken] = useState('')
	const [isInvalid, setIsInvalid] = useState(false)
	const [redirect] = useQueryState('redirect')
	const router = useRouter()

	const check = (e: FormEvent) => {
		e.preventDefault()
		fetch('/verify', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json',
			},
			body: JSON.stringify({ token }),
		})
			.then(res => res.json())
			.then(data => {
				if (data.status === 'ok') {
					router.push(redirect || '/')
				} else {
					setIsInvalid(true)
				}
			})
	}

	return (
		<div className="grid h-screen w-screen place-items-center">
			<form
				className="flex min-h-[200px] min-w-[400px] flex-col items-center justify-center gap-2 rounded-lg bg-gray-200 p-2"
				onSubmit={check}
			>
				<label htmlFor="token" className="text-lg">
					Token
				</label>
				<input
					id="token"
					type="text"
					placeholder="Token"
					value={token}
					onChange={e => {
						setToken(e.target.value)
						setIsInvalid(false)
					}}
					className={twMerge(
						'w-2/3 rounded border-2 border-gray-300 p-2 text-center outline-none focus:border-blue-500',
						isInvalid && 'border-red-500 focus:border-red-500',
					)}
				/>
				{isInvalid && <p className="text-red-500">Invalid token</p>}
				<button type="submit" className="w-2/3 rounded bg-blue-500 p-2 text-white">
					Submit
				</button>
			</form>
		</div>
	)
}

export default function Page() {
	return (
		<Suspense>
			<Verify />
		</Suspense>
	)
}
