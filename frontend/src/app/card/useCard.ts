import { useSSE } from '@/hooks/useSSE'
import { ensureSession, GetCurrentSession, GetSessionByID, Session, ZeroSession } from '@/sdk/sdk'
import { parseAsString, useQueryState } from 'nuqs'
import { useEffect, useState } from 'react'

export function useCard() {
	const [room] = useQueryState('room', parseAsString.withDefault('R0'))
	const [id] = useQueryState('id')
	const [card, setCard] = useState(ZeroSession)

	const keepUpdate = id == null
	const sseStr = keepUpdate ? `card-current-${room}` : `card-id-${id}`

	const update = useSSE<Session>(sseStr).at(-1)
	const [error, setError] = useState<Error>()

	// init
	useEffect(() => {
		if (keepUpdate) GetCurrentSession(room).then(setCard, setError)
		else GetSessionByID(id).then(setCard, setError)
	}, [id, keepUpdate, room])

	useEffect(() => {
		if (!update) return
		setCard(ensureSession(update))
	}, [update])

	return [card, error] as const
}
