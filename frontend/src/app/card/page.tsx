'use client'

import { useSession } from '@/hooks/useSession'
import { Card } from '@/components/card'
import { useQuery } from '@/hooks/util/useQuery'

export default function Page() {
	const [room] = useQuery('room', 'R0')
	const [id] = useQuery('id', '')
	const session = useSession(room, id || undefined)

	return session ? <Card session={session} /> : <div>Loading...</div>
}
