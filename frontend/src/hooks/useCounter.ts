import { useCallback } from 'react'
import { useSSEFetchValue } from './util/useSSE'
import { Get } from '@/sdk/counter'

export function useCounter(name: string) {
	return useSSEFetchValue(
		`counter/${name}`,
		useCallback(() => Get(name), [name]),
	)
}
