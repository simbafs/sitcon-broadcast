'use client'
import { Sessions } from '@/types/card'
import { btn } from '@/varients/btn'
import { useState } from 'react'
import { Room } from './room'
import useSWR from 'swr'
import { twMerge } from 'tailwind-merge'
import { useSSEFetch } from '@/hooks/useSSE'
import { formatTime } from '@/utils/formatTime'

export default function Page() {
	const [room, setRoom] = useState('R0')
	const now = useSSEFetch('now', () => fetch('/api/now').then(res => res.json()))

	const { data, error } = useSWR<Sessions>('/api/session', (url: string) => fetch(url).then(res => res.json()))

	if (error)
		return (
			<>
				<h1>Error</h1>
				<pre>{JSON.stringify(error, null, 2)}</pre>
			</>
		)
	if (!data) return <h1>Loading...</h1>

	return (
		<div className="flex min-h-screen w-screen flex-col items-center px-8 py-4">
			<select
				value={room}
				onChange={e => setRoom(e.target.value)}
				className={twMerge(btn({ size: '4xl' }), 'w-full')}
			>
				{Object.keys(data).map(r => (
					<option key={r} value={r}>
						{r}
					</option>
				))}
			</select>
			{now !== undefined && <h1 className="mt-4 text-4xl">{formatTime(now)}</h1>}

			<Room sessions={data[room]} key={room} />
		</div>
	)
}
