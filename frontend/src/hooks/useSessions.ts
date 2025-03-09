import { GetAllInRoom, Session } from '@/sdk/session'
import { useState, useEffect } from 'react'

export function useSessions(room: string) {
	const [sessions, setSessions] = useState<Session[]>([])

	useEffect(() => {
		GetAllInRoom(room).then(setSessions)
	}, [room])

	// TODO: get update from SSE or websocket

	return sessions
}
