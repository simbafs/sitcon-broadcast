import { api } from './api'

export type Event = {
	name: string
	script: string
	url: string
}

export const ZeroEvent: Event = {
	name: '',
	script: '',
	url: '',
}

export function GetAll() {
	return api<Event[]>('/event', 'GET')
}

export function SetScript(name: string, script: string) {
	return api<Event>(`/event/${name}`, 'PUT', {
		script,
	})
}

export function CreateEvent(name: string, url: string) {
	return api<Event>('/event', 'POST', {
		name,
		url,
		script: '',
	})
}
