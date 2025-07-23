import type { Display } from 'controly'
import { useCard, type Room } from '../hooks/useCard'
import { useEffect, useState } from 'react'

function formatTime(date: string) {
	const d = new Date(date)
	return d.toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' })
}

export function Card({ display }: { display: Display }) {
	const [room, setRoom] = useState<Room>('R0')
	const { session, next, clear } = useCard(room)

	useEffect(() => {
		display.command('next', next)
		display.command('clear', clear)
		display.command('set_room', args => setRoom(args.value))
	}, [display, next, clear])

	if (!session) {
		return <span>Error: No session data available.</span>
	}

	return (
		<div className="rounded-lg shadow-lg p-6 w-[400px] min-h-[300px]">
			<h1 className="text-3xl font-bold mb-4">{session.title}</h1>
			<p className="text-xl text-gray-500 mb-2">
				{formatTime(session.start)} - {formatTime(session.end)}
			</p>
			<p className="text-xl font-medium">{session.speakers.join('、')}</p>
		</div>
	)
}
