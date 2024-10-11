'use client'
import { RoomData } from '@/hooks/useCountdown'
import useQuery from '@/hooks/useQuery'
import { useSSE } from '@/hooks/useSSE'
import { formatTime } from '@/utils/formatTime'
import { useEffect, useState } from 'react'

export default function Home() {
	const room = useQuery('room', 'R0')
	const [time, setTime] = useState(0)

	const latest = useSSE<RoomData>(`countdown-${room}`).at(-1)

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
		<div className="w-screen h-screen grid place-items-center">
			<h1 className="text-[35vw] leading-[0.8]">{formatTime(time || 0)}</h1>
		</div>
	)
}
