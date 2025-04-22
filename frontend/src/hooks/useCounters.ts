import { Counter } from '@/sdk/counter'
import { useFetch } from './util/useFetch'

export function useCounters() {
	return useFetch<Record<string, Counter>>('/api/counter')
}
