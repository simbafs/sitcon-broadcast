'use client'

import { Card } from '@/components/card'
import { useSessions } from '@/hooks/useSessions'
import { parseAsString, useQueryState } from 'nuqs'

export default function Page() {
	const [room] = useQueryState('room', parseAsString.withDefault('R0'))
	const sessions = useSessions(room)

	return (
		<div className="flex flex-col items-center m-6">
			<div className=" max-w-2xl">
			    <h1 className="text-5xl font-bold text-center m-10">{room}</h1>
				{sessions.map(session => (
					<Card key={session.session_id} session={session} />
				))}
			</div>
		</div>
	)
}
