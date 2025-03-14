import { GetCurrentSession, GetSession, Session } from '@/sdk/session'
import { useEffect, useState } from 'react'
import { toast } from 'react-toastify'

export function useSession(room: string, id?: string) {
	const [session, setSession] = useState<Session>()

	const getSession = (room: string, id?: string) => {
		if (id)
			GetSession(room, id)
				.then(setSession)
				.catch(e => toast(e.message))
		else
			GetCurrentSession(room)
				.then(setSession)
				.catch(e => toast(e.message))
	}

	useEffect(() => getSession(room, id), [room, id])

	useEffect(() => {
		const timer = setInterval(() => getSession(room, id), 1000)
		return () => clearInterval(timer)
	}, [id, room])

	// TODO: get update from SSE or websocket

	return session
}
