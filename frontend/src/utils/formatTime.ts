const to2 = (n: number) => n.toString().padStart(2, '0')

const timezone = 8

export function formatTime(time: Date) {
	if (time instanceof Date == false) {
		// console.log(time)
		time = new Date(time)
	}
	if (time === new Date(0)) return '00:00'
	const hour = time.getUTCHours() + timezone
	const minute = time.getUTCMinutes()
	return `${to2(hour)}:${to2(minute)}`
}

export function formatCountdown(time: number) {
	const hour = Math.floor(time / 60)
	const minute = time % 60
	return `${to2(hour)}:${to2(minute)}`
}
