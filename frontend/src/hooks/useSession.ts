import { GetCurrentSession, GetSession, ZeroSession } from '@/sdk/session'
import { usePolling } from './usePolling'
import { useCallback } from 'react'

export function useSession(room: string, id?: string) {
	const fn = useCallback(() => (id ? GetSession(room, id) : GetCurrentSession(room)), [id, room])
	return usePolling(fn, ZeroSession, { interval: 300 })
}
