import { useEffect, useState } from 'react'

export function useNow(t = 0) {
	const [now, setNow] = useState(0)

	useEffect(() => {
		const id = setInterval(() => {
			setNow(Math.floor(Date.now() / 1000))
		}, 1000)
		return () => {
			clearInterval(id)
		}
	}, [])

	return Math.max(t, now)
}
