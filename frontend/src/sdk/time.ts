export type Time = {
	year: number
	month: number
	day: number
	hours: number
	minutes: number
	seconds: number
}

export function parseTime(t: number): Time {
	const date = new Date(t * 1000)
	return {
		year: date.getUTCFullYear(),
		month: date.getUTCMonth() + 1,
		day: date.getUTCDate(),
		hours: date.getUTCHours() + 8,
		minutes: date.getUTCMinutes(),
		seconds: date.getUTCSeconds(),
	}
}

export function constructTime(t: Time) {
	const { year, month, day, hours, minutes, seconds } = t
	return Date.UTC(year, month - 1, day, hours - 8, minutes, seconds) / 1000
}
