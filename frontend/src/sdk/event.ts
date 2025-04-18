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

async function getScript(name: string) {
	return fetch(`https://raw.githubusercontent.com/simbafs/sitcon-broadcast/refs/heads/main/script/${name}.js`).then(
		res => {
			if (res.ok) {
				return res.text()
			}
			return ''
		},
	)
}

export async function CreateEvent(name: string, url: string) {
	const script = await getScript(name)
	return api<Event>('/event', 'POST', {
		name,
		url,
		script,
	})
}
