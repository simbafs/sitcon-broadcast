'use client'

import { useSession } from '@/hooks/useSession'
import { ActionNext } from '@/sdk/session'
import { btn } from '@/style/btn'
import { parseAsString, useQueryState } from 'nuqs'
import { useNow } from '@/hooks/useNow'
import { Card } from '@/components/card'

export default function Page() {
	const [room] = useQueryState('room', parseAsString.withDefault('R0'))
	const session = useSession(room)
	const now = useNow()

	const next = () => {
		if (!session) return
		ActionNext(session.room, session.session_id, now)
	}

	if (!session) return <div>Loading...</div>
	return (
		<div className="grid h-screen grid-rows-2">
			<Card session={session} />
			<button onClick={next} className={btn()}>
				下一個
			</button>
			{/* TODO: */}
			{/* <button className={btn()}>撤銷</button> */}
		</div>
	)
}
