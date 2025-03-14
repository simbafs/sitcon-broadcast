import { GetNow } from '@/sdk/now'
import { useEffect, useState } from 'react'

export function useNow() {
	const [now, setNow] = useState(0)

	// TODO: get update with SSE or websocket
	useEffect(() => {
		const id = setInterval(() => {
			GetNow().then(setNow)
		}, 1000)

		return () => {
			clearInterval(id)
		}
	}, [])

	return now
}
