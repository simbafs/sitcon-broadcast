const to2 = (n: number) => n.toString().padStart(2, '0')
export function formatTime(time: Date) {
	const hour = time.getHours()
	const minute = time.getMinutes()
	return `${to2(hour)}:${to2(minute)}`
}

export function formatCountdown(time: number) {
	const hour = Math.floor(time / 60)
	const minute = time % 60
	return `${to2(hour)}:${to2(minute)}`
}
