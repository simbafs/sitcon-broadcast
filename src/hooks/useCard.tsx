import { useState, useEffect } from 'react'
import data from '../data.json'

export type Session = {
	title: string
	start: string
	end: string
	speakers: string[]
	finish: boolean
}

export type Room = keyof typeof data

type ScheduleData = {
	[key in Room]: Session[]
}

export function useCard(room: Room) {
	const [sessions, setSessions] = useState<Session[]>([])

	useEffect(() => {
		const storageKey = `sessions-${room}`
		try {
			const savedSessions = localStorage.getItem(storageKey)
			if (savedSessions) {
				setSessions(JSON.parse(savedSessions))
				return
			}
		} catch (error) {
			console.error('Failed to load sessions from localStorage', error)
		}

		// Fallback to data.json
		const roomSessions = (data as ScheduleData)[room]
		const initialSessions = roomSessions.map(s => ({ ...s, finish: !!s.finish }))
		setSessions(initialSessions)
		try {
			localStorage.setItem(storageKey, JSON.stringify(initialSessions))
		} catch (error) {
			console.error('Failed to save sessions to localStorage', error)
		}
	}, [room])

	const currentIndex = sessions.findIndex(s => !s.finish)
	const session =
		currentIndex !== -1 ? sessions[currentIndex] : sessions.length > 0 ? sessions[sessions.length - 1] : undefined

	const next = () => {
		console.log('next')
		if (currentIndex === -1) {
			return
		}

		const now = new Date().toISOString()
		const updatedSessions = [...sessions]

		updatedSessions[currentIndex] = {
			...updatedSessions[currentIndex],
			end: now,
			finish: true,
		}

		if (currentIndex + 1 < updatedSessions.length) {
			updatedSessions[currentIndex + 1] = {
				...updatedSessions[currentIndex + 1],
				start: now,
			}
		}

		setSessions(updatedSessions)
		try {
			localStorage.setItem(`sessions-${room}`, JSON.stringify(updatedSessions))
		} catch (error) {
			console.error('Failed to save updated sessions to localStorage', error)
		}
	}

	const clear = () => {
		console.log('clear')
		const storageKey = `sessions-${room}`
		try {
			localStorage.removeItem(storageKey)
			const roomSessions = (data as ScheduleData)[room]
			const initialSessions = roomSessions.map(s => ({ ...s, finish: !!s.finish }))
			setSessions(initialSessions)
		} catch (error) {
			console.error('Failed to clear sessions from localStorage', error)
		}
	}

	return { session, next, clear }
}
