import { useEffect, useState } from 'react'
import { useSSE } from './useSSE'
import { COUNTING, GetCountdownByName, PAUSE, Room, UpdateCountdown } from '@/sdk/sdk'

// export const PAUSE = 0
// export const COUNTING = 1
//
// export type State = typeof PAUSE | typeof COUNTING
//
// export type RoomData = {
// 	inittime: number
// 	time: number
// 	state: State
// 	name: string
// }

export type Countdown = ReturnType<typeof useCountdown>

export function useCountdown(name: string) {
	const [stopUpdate, setStopUpdate] = useState(false)
	const [countdown, setCountdown] = useState<Room>({
		inittime: 10,
		time: 0,
		state: 0,
		name: name,
	})
	const latest = useSSE<Room>(`countdown-${name}`).at(-1)

	useEffect(() => console.log(latest), [latest])

	// get init
	useEffect(() => {
		GetCountdownByName(name).then(setCountdown).catch(console.error)
	}, [name])

	// get update
	useEffect(() => {
		if (!latest || stopUpdate) return
		setCountdown(latest)
	}, [latest, stopUpdate])

	// useEffect(() => console.log(countdown), [countdown])

	// operations
	const start = () => {
		setStopUpdate(true)
		UpdateCountdown(name, {
			...countdown,
			state: COUNTING,
		})
			.finally(() => setStopUpdate(false))
			.catch(console.error)
	}
	const pause = () => {
		setStopUpdate(true)
		UpdateCountdown(name, {
			...countdown,
			state: PAUSE,
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
		UpdateCountdown(name, {
			...countdown,
			time: time,
			inittime: time,
			state: PAUSE,
		})
			.finally(() => setStopUpdate(false))
			.catch(console.error)
	}
	const reset = () => {
		setStopUpdate(true)
		UpdateCountdown(name, {
			...countdown,
			time: countdown.inittime,
			state: PAUSE,
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
