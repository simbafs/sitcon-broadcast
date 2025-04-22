import { GetCurrentSession, GetSession, ZeroSession } from '@/sdk/session'
import { useCallback } from 'react'
import { useSSEFetchValue } from './useSSE'

export function useSession(room: string, id?: string) {
	const fn = useCallback(() => (id ? GetSession(room, id) : GetCurrentSession(room)), [id, room])
	return useSSEFetchValue(`room/${room}`, fn) || ZeroSession
}
