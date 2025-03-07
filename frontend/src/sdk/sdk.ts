// GET    /api/card/
// GET    /api/card/:id
// GET    /api/card/current/:room
// PUT    /api/card/:room/:id
//
// GET    /api/now/
// PUT    /api/now/
// DELETE /api/now/
//
// GET    /api/countdown/
// GET    /api/countdown/:name
// PUT    /api/countdown/:name

type Method = 'POST' | 'GET' | 'PUT' | 'DELETE'

export function parseJSONWithDates(jsonString: string) {
	const data = JSON.parse(jsonString, (key, value) => {
		if (typeof value === 'string' && value.match(/\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}/)) {
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
	return api<{ now: Date }>('/now', 'GET').then(data => new Date(data.now))
}

export async function SetNow(now: Date) {
	return api('/now', 'PUT', { now })
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

export function ensureSession(session: Partial<Session>): Session {
	return {
		...ZeroSession,
		...session,
	}
}

function convertDate(s: Session) {
	s.start = new Date(s.start)
	s.end = new Date(s.end)
	return s
}

export async function GetAllSessions() {
	return api<Session[]>('/card', 'GET').then(sessions => sessions.map(ensureSession).map(convertDate))
}

export async function GetSessionByID(id: string) {
	return api<Session>(`/card/${id}`, 'GET').then(ensureSession).then(convertDate)
}

export async function GetCurrentSession(room: string) {
	return api<Session>(`/card/current/${room}`, 'GET').then(ensureSession).then(convertDate)
}

export async function UpdateSession(room: string, id: string, start: Date, end: Date) {
	return api(`/card/${room}/${id}`, 'PUT', {
		start: new Date(start).toISOString(),
		end: new Date(end).toISOString(),
	})
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
	return api(`/countdown/${name}`, 'PUT', updated)
}

export type Special = {
	title: string
	speakers: string
	titleStyle: string
	speakersStyle: string
}

export const ZeroSpecial: Special = {
	title: '',
	speakers: '',
	titleStyle: '',
	speakersStyle: '',
}

export function ensureSpecial(special: Partial<Special>): Special {
	return {
		...ZeroSpecial,
		...special,
	}
}

export async function GetAllSpecial() {
	return api<Special>(`/special`, 'GET')
}

export async function GetSpecialByID(id: string) {
	return api<Special>(`/special/${id}`, 'GET')
}

export async function UpdateSpecial(id: string, updated: string) {
	return api<Special>(`/special/${id}`, 'PUT', updated)
}
