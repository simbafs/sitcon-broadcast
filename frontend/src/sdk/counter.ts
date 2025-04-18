import { api } from './api'

export const StatusStopped = 0
export const StatusPause = 1
export const StatusRunning = 2

export type Status = typeof StatusStopped | typeof StatusPause | typeof StatusRunning

export type Counter = {
	name: string
	init: number
	count: number
	status: Status
}

export function GetAll() {
	return api<Counter[]>('/counter', 'GET')
}

export function Get(name: string) {
	return api<Counter>(`/counter/${name}`, 'GET')
}

export function Set(name: string, init: number) {
	return api<Counter>(`/counter/${name}`, 'PUT', {
		init,
	})
}

export function Start(name: string) {
	return api<Counter>(`/counter/${name}/start`, 'PUT')
}

export function Stop(name: string) {
	return api<Counter>(`/counter/${name}/stop`, 'PUT')
}

export function Pause(name: string) {
	return api<Counter>(`/counter/${name}/pause`, 'PUT')
}
