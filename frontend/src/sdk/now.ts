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
