'use client'
import { btn } from '@/varients/btn'
import { Room } from './room'
import { twMerge } from 'tailwind-merge'
import { useSSEFetch } from '@/hooks/useSSE'
import { formatTime } from '@/utils/formatTime'
import { parseAsString, useQueryState } from 'nuqs'
import { Suspense } from 'react'
import { GetNow } from '@/sdk/sdk'
import { useSessions } from '@/app/card/admin/useSessions'

function Admin() {
	const [room, setRoom] = useQueryState('room', parseAsString.withDefault('R0'))
	const now = useSSEFetch('now', GetNow)

	const [sessions, error] = useSessions(room)

	if (error)
		return (
			<>
				<h1>Error</h1>
				<pre>{JSON.stringify(error, null, 2)}</pre>
			</>
		)
	if (!sessions) return <h1>Loading...</h1>

	return (
		<div className="flex min-h-screen w-screen flex-col items-center px-8 py-4">
			<div className="flex w-full gap-4">
				<select
					value={room}
					onChange={e => setRoom(e.target.value)}
					className={twMerge(btn({ size: '4xl' }), 'grow')}
				>
					{['R0', 'R1', 'R2', 'R3', 'S'].map(r => (
						<option key={r} value={r}>
							{r}
						</option>
					))}
				</select>
				<a
					href={`/card?room=${room}`}
					className={twMerge(btn({ size: '2xl' }), 'grid place-items-center')}
					target="_blank"
				>
					<span>開啟字卡</span>
				</a>
			</div>
			{now !== undefined && <h1 className="mt-4 text-4xl">{formatTime(now)}</h1>}

			<Room sessions={sessions} room={room} key={room} />
		</div>
	)
}

export default function Page() {
	return (
		<Suspense>
			<Admin />
		</Suspense>
	)
}
