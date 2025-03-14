import { api } from './api'

export function GetNow() {
	return api<number>('/now', 'GET')
}

export function SetNow(now: number) {
	return api<number>('/now', 'POST', {
		now,
	})
}

export function ResetNow() {
	return api<number>('/now', 'DELETE')
}

export type Now = {
	year: number
	month: number
	day: number
	hours: number
	minutes: number
	seconds: number
}

export function parseNow(now: number): Now {
	const date = new Date(now * 1000)
	return {
		year: date.getUTCFullYear(),
		month: date.getUTCMonth() + 1,
		day: date.getUTCDate(),
		hours: date.getUTCHours() + 8,
		minutes: date.getUTCMinutes(),
		seconds: date.getUTCSeconds(),
	}
}

export function constructNow(now: Now) {
	const { year, month, day, hours, minutes, seconds } = now
	return Date.UTC(year, month - 1, day, hours - 8, minutes, seconds) / 1000
}

