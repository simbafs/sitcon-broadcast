import { formatTime } from '@/utils/formatTime'
import { useEffect, useState } from 'react'

function getNow() {
	const now = new Date()
	return now.getHours() * 60 + now.getMinutes()
}

export function useTime(title: string, init: number, onChange: (value: number) => void) {
	const [time, setTime] = useState(init)

	useEffect(() => onChange(time), [onChange, time])

	const component = (
		<div>
			<p>
				{title}: {formatTime(time)}
			</p>
			<div>
				<button onClick={() => setTime(time - 5)}>-5</button>
				<button onClick={() => setTime(time - 1)}>-1</button>
				<button onClick={() => setTime(getNow())}>Now</button>
				<button onClick={() => setTime(time + 1)}>+1</button>
				<button onClick={() => setTime(time + 5)}>+5</button>
			</div>
		</div>
	)

	return [component, time] as const
}
