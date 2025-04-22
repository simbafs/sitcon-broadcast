import { GetNow } from '@/sdk/now'
import { useSSEFetchValue } from './util/useSSE'

export function useNow() {
	return useSSEFetchValue('now', GetNow) || 0
}
