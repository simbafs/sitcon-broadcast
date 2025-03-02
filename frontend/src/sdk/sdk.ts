// GET /card
// GET /card/:id
// GET /card/current/:room
// PUT /card/:id
//
// GET /countdown
// GET /countdown/:name
// PUT /countdown/:name
//
// GET /now
// PUT /now
// Delete /now

type Method = 'POST' | 'GET' | 'PUT' | 'DELETE'

export function parseJSONWithDates(jsonString: string) {
	const data = JSON.parse(jsonString, (key, value) => {
		if (typeof value === 'string' && value.match(/^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}\.\d+([+-]\d{2}:\d{2}|Z)$/)) {
			return new Date(value)
		}
		return value
	})
	return data
}

async function api<T extends any>(path: string, method: Method, body?: any) {
	let input = `/api`

	if (!path.startsWith('/')) input += '/' + path
	else input += path

	if (!input.endsWith('/')) input += '/'

	return fetch(`/api${path}`, {
		method,
		body: JSON.stringify(body),
	})
		.then(res => res.text())
		.then(parseJSONWithDates)
		.then(body => {
			if (body.error) throw new Error(body.error)
			return body as T
		})
}

export async function GetNow() {
	return api<{ now: Date }>('/now', 'GET').then(data => data.now)
}

export async function SetNow(now: Date) {
	return api('/now', 'POST', {
		now,
	})
}

export async function ResetNow() {
	return api('/now', 'DELETE')
}

export type Session = {
	id: string
	title: string
	type: string
	speakers: string[]
	room: string
	broadcast: string[]
	start: Date
	end: Date
	slido: string
	slide: string
	hackmd: string
}

export const ZeroSession: Session = {
	id: '',
	title: '',
	type: '',
	speakers: [],
	room: '',
	broadcast: [],
	start: new Date(0),
	end: new Date(0),
	slido: '',
	slide: '',
	hackmd: '',
}

function ensureSession(session: Partial<Session>): Session {
	return {
		...ZeroSession,
		...session,
	}
}

export async function GetAllSessions() {
	return api<Session[]>('/card', 'GET').then(sessions => sessions.map(ensureSession))
}

export async function GetSessionByID(id: string) {
	return api<Session>(`/card/${id}`, 'GET').then(ensureSession)
}

export async function GetCurrentSession(room: string) {
	return api<Session>(`/card/current/${room}`, 'GET').then(ensureSession)
}

export async function UpdateSession(id: string, start: Date, end: Date) {
	return api(`/card/${id}`, 'PUT', { start, end })
}

export const PAUSE = 0
export const COUNTING = 1
export type CountdownState = typeof PAUSE | typeof COUNTING

export type Room = {
	inittime: number
	time: number
	state: CountdownState
	name: string
}

export async function GetAllCountdown() {
	return api<Room[]>(`/countdown`, 'GET')
}

export async function GetCountdownByName(name: string) {
	return api<Room>(`/countdown/${name}`, 'GET')
}

export async function UpdateCountdown(name: string, updated: Room) {
	console.log('update countdown', name, updated)
	return api(`/countdown/${name}`, 'PUT', updated)
}
