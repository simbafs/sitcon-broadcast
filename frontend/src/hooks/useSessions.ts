import { GetAllInRoom } from '@/sdk/session'
import { usePolling } from './usePolling'

export function useSessions(room: string) {
	return usePolling(() => GetAllInRoom(room), [])
}
