const to2 = (n: number) => n.toString().padStart(2, '0')
export function formatTime(time: number) {
	const hour = Math.floor(time / 60)
	const minute = time % 60
	return `${to2(hour)}:${to2(minute)}`
}
