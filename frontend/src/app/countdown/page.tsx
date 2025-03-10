'use client'
import { useSSE } from '@/hooks/useSSE'
import { Room } from '@/sdk/sdk'
import { formatCountdown } from '@/utils/formatTime'
import { parseAsString, useQueryState } from 'nuqs'
import { Suspense, useEffect, useState } from 'react'

export default function Page() {
	return (
		<Suspense>
			<Countdown />
		</Suspense>
	)
}
function Countdown() {
	const [room] = useQueryState('room', parseAsString.withDefault('R0'))
	const [time, setTime] = useState(0)

	const latest = useSSE<Room>(`countdown-${room}`).at(-1)

	// init
	useEffect(() => {
		fetch(`/api/countdown/${room}`)
			.then(res => res.json())
			.then(data => setTime(data.time))
			.catch(console.error)
	}, [room])

	// update
	useEffect(() => {
		if (latest === undefined) return
		setTime(latest.time)
	}, [latest])

	return (
		<div className="grid h-screen w-screen place-items-center">
			<h1 className="text-[35vw] leading-[0.8]">{formatCountdown(time || 0)}</h1>
		</div>
	)
}
