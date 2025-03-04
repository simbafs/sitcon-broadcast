import { ensureSession, GetAllSessions, Session } from '@/sdk/sdk'
import { useEffect, useState } from 'react'
import { useSSE } from '../../../hooks/useSSE'

// TODO: only store the sessions that are in the room
export function useSessions(room: string) {
	const [sessions, setSessions] = useState<Session[]>([])
	const [error, setError] = useState<Error>()

	// get initial state
	useEffect(() => {
		GetAllSessions().then(setSessions).catch(setError)
	}, [room])

	// handle udpates
	const currentArr = useSSE<Session>(`card-${room}`)
	useEffect(() => {
		const current = Object.fromEntries(currentArr.map(s => [s.id, ensureSession(s)]))
		setSessions(ss => ss.map(s => (s.id in current ? current[s.id] : s)))
	}, [currentArr])

	const s = sessions.filter(s => s.room == room || s.broadcast.includes(room))

	return [s, error] as const
}
