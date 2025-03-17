'use client'

import { parseAsString, useQueryState } from 'nuqs'
import { useSession } from '@/hooks/useSession'
import { Card } from '@/components/card'
import { Suspense } from 'react'

function CardPage() {
	const [room] = useQueryState('room', parseAsString.withDefault('R0'))
	const [id] = useQueryState('id')
	const session = useSession(room, id || undefined)

	return session ? <Card session={session} /> : <div>Loading...</div>
}

export default function Page() {
	return <Suspense>
		<CardPage />
	</Suspense>
}
