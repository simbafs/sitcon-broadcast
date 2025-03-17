import { GetNow } from '@/sdk/now'
import { usePolling } from './usePolling'

export function useNow() {
	return usePolling(GetNow, 0)
}
