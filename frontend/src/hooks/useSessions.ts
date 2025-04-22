import { GetAllInRoom } from '@/sdk/session'
import { useSSEFetchValue } from './util/useSSE'
import { useCallback } from 'react'

export function useSessions(room: string) {
	return useSSEFetchValue(
		`room/${room}`,
		useCallback(() => GetAllInRoom(room), [room]),
	)
}
