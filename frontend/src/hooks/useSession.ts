import { GetCurrentSession, GetSession, Session } from '@/sdk/session'
import { useEffect, useState } from 'react'

export function useSession(room: string, id?: string) {
	const [session, setSession] = useState<Session>()

	useEffect(() => {
		if (id) GetSession(room, id).then(setSession)
		else GetCurrentSession(room).then(setSession)
	}, [room, id])

	// TODO: get update from SSE or websocket

	return session
}
