'use client'
import { Session } from '@/sdk'
import { useState } from 'react'
import { Time } from './time'
import { useNow } from '@/hooks/useNow'

export function Card({ session }: { session: Session }) {
	const [raw, setRaw] = useState(false)
	const now = useNow()
	return (
		<div onDoubleClick={() => setRaw(!raw)} className="m-2 flex flex-col rounded-lg border-2 border-black p-2 overflow-scroll">
			{raw ? (
				<pre>{JSON.stringify(session, null, 2)}</pre>
			) : (
				<>
					<h1 className="text-2xl">{session.title}</h1>
					<p>
						{session.room}: {session.session_id}
					</p>
					<Time time={session.start} />
					<Time time={session.finish ? session.end : now} />
				</>
			)}
			<div className="grow" />
			<hr />
			<p className="text-gray-500">Double click to show raw</p>
		</div>
	)
}
