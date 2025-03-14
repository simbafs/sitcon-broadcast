'use client'

import { parseAsString, useQueryState } from 'nuqs'
import { useSession } from '@/hooks/useSession'
import { Card } from '@/components/card'

export default function Page() {
	const [room] = useQueryState('room', parseAsString.withDefault('R0'))
	const [id] = useQueryState('id')
	const session = useSession(room, id || undefined)

	return session ? <Card session={session} /> : <div>Loading...</div>
}
