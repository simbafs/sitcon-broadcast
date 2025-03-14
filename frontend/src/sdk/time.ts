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
		year: date.getFullYear(),
		month: date.getMonth() + 1,
		day: date.getDate(),
		hours: date.getHours(),
		minutes: date.getMinutes(),
		seconds: date.getSeconds(),
	}
}

export function constructTime(t: Time) {
	const { year, month, day, hours, minutes, seconds } = t
	const timezome = new Date().getTimezoneOffset() / 60
	return Date.UTC(year, month - 1, day, hours + timezome, minutes, seconds) / 1000
}
