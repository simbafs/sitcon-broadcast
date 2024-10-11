import { useEffect, useState } from 'react'
import { useSSE } from './useSSE'

export const PAUSE = 0
export const COUNTING = 1

export type State = typeof PAUSE | typeof COUNTING

export type RoomData = {
	inittime: number
	time: number
	state: State
	name: string
}

export type Countdown = ReturnType<typeof useCountdown>

export function useCountdown(name: string) {
	const [stopUpdate, setStopUpdate] = useState(false)
	const [countdown, setCountdown] = useState<RoomData>({
		inittime: 10,
		time: 0,
		state: 0,
		name: name,
	})
	const latest = useSSE<RoomData>(`countdown-${countdown.name}`).at(-1)

	// get init
	useEffect(() => {
		fetch(`/api/countdown/${name}`)
			.then(res => res.json())
			.then(data => setCountdown(data.room))
			.catch(console.error)
	}, [name])

	// get update
	useEffect(() => {
		if (!latest || stopUpdate) return
		setCountdown(latest)
	}, [latest, stopUpdate])

	useEffect(() => console.log(countdown), [countdown])

	// operations
	const start = () => {
		setStopUpdate(true)
		fetch(`/api/countdown/${name}`, {
			method: 'post',
			body: JSON.stringify({
				...countdown,
				state: COUNTING,
			}),
		})
			.finally(() => setStopUpdate(false))
			.catch(console.error)
	}
	const pause = () => {
		setStopUpdate(true)
		fetch(`/api/countdown/${name}`, {
			method: 'post',
			body: JSON.stringify({
				...countdown,
				state: PAUSE,
			}),
		})
			.finally(() => setStopUpdate(false))
			.catch(console.error)
	}
	const setTime = (time: number) => {
		setStopUpdate(true)
		setCountdown({
			...countdown,
			time: time,
			inittime: time,
		})
		fetch(`/api/countdown/${name}`, {
			method: 'post',
			body: JSON.stringify({
				...countdown,
				time: time,
				inittime: time,
				state: PAUSE,
			}),
		})
			.finally(() => setStopUpdate(false))
			.catch(console.error)
	}
	const reset = () => {
		setStopUpdate(true)
		fetch(`/api/countdown/${name}`, {
			method: 'post',
			body: JSON.stringify({
				...countdown,
				time: countdown.inittime,
				state: PAUSE,
			}),
		})
			.finally(() => setStopUpdate(false))
			.catch(console.error)
	}

	return {
		...countdown,
		start,
		pause,
		setTime,
		reset,
	}
}
