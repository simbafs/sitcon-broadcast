import { api } from './api'

export type Session = {
	start: number // Unix timestamp in second
	end: number // Unix timestamp in second
	session_id: string
	room: string
	next: string // next session ID
	title: string
	data: Record<string, any>
}

export const ZeroSession: Session = {
	start: 0,
	end: 0,
	session_id: '',
	room: '',
	next: '',
	title: '',
	data: {},
}

export function GetAllInRoom(room: string) {
	return api<Session[]>(`/session/${room}/all`, 'GET')
}

export function GetSession(room: string, id: string) {
	return api<Session>(`/session/${room}/${id}`, 'GET')
}

export function GetCurrentSession(room: string) {
	return api<Session>(`/session/${room}`, 'GET')
}

export function ActionNext(room: string, id: string, end: number) {
	return api<Session>(`/session/${room}/${id}`, 'POST', {
		end,
	})
}
