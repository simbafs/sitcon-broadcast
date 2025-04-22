import { GetAllInRoom } from '@/sdk/session'
import { useSSEFetchValue } from './useSSE'
import { useCallback } from 'react'

export function useSessions(room: string) {
	return useSSEFetchValue(
		`room/${room}`,
		useCallback(() => GetAllInRoom(room), [room]),
	)
	// return usePolling(() => GetAllInRoom(room), [])
}
