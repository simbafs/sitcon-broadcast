'use client'

import { useSession } from '@/hooks/useSession'
import { ActionNext } from '@/sdk'
import { btn } from '@/style/btn'
import { parseAsString, useQueryState } from 'nuqs'
import { useNow } from '@/hooks/useNow'

export default function Page() {
	const [room] = useQueryState('room', parseAsString.withDefault('R0'))
	const session = useSession(room)
	const now = useNow()

	const next = () => {
		if (!session) return
		ActionNext(session.room, session.session_id, now)
	}

	return (
		<div className="grid h-screen grid-rows-3">
			<pre>{JSON.stringify({ ...session, end: Math.max(session?.end || 0, now) }, null, 2)}</pre>
			<button onClick={next} className={btn()}>
				下一個
			</button>
			<button className={btn()}>撤銷</button>
		</div>
	)
}
