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

export function SaveEvent(name: string, script: string) {
	return api<Event>(`/event/${name}`, 'PUT', {
		script,
	})
}
