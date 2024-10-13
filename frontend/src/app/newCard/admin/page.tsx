'use client'
import { Sessions } from '@/types/card'
import { btn } from '@/varients/btn'
import { useState } from 'react'
import { Room } from './room'
import useSWR from 'swr'
import { twMerge } from 'tailwind-merge'

export default function Page() {
	const [room, setRoom] = useState('R0')

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

			<Room sessions={data[room]} key={room} />
		</div>
	)
}
