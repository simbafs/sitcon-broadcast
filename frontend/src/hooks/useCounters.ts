import { Counter, GetAll } from '@/sdk/counter'
import { useEffect, useState } from 'react'

export function useCounters() {
	const [counters, setCounters] = useState<Record<string, Counter>>({})
	useEffect(() => {
		GetAll().then(setCounters)
	}, [])
	return counters
}
