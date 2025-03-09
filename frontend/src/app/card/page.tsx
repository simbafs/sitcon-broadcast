'use client' 

import { parseAsString, useQueryState } from 'nuqs'
import { useSession } from '@/hooks/useSession'

export default function Page() {
	const [room] = useQueryState('room', parseAsString.withDefault('R0'))
	const [id] = useQueryState('id')
	const session = useSession(room, id || undefined)

	return <pre>{JSON.stringify(session, null, 2)}</pre>
}
