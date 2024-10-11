'use client'
import { createContext, useContext, useEffect, ReactNode, useReducer } from 'react'

type Events = Record<string, string[]>

const SSEContext = createContext<Events>({})

export const SSEProvider = ({ children, url }: { children: ReactNode; url: string }) => {
	const [events, updateEvents] = useReducer((state: Events, action: { name: string; data: string }) => {
		if (action.name in state) {
			return {
				...state,
				[action.name]: [...state[action.name], action.data],
			}
		} else {
			return {
				...state,
				[action.name]: [action.data],
			}
		}
	}, {})

	useEffect(() => {
		console.log({ url })
		const eventSource = new EventSource(url)

		eventSource.onmessage = event => {
			try {
				const parsedData: { name: string; data: string } = JSON.parse(event.data)
				updateEvents(parsedData)
			} catch (err) {
				console.error('Error parsing SSE data:', err)
			}
		}

		eventSource.onerror = err => {
			console.error('EventSource failed:', err)
			eventSource.close()
		}

		return () => {
			eventSource.close()
		}
	}, [url])

	return <SSEContext.Provider value={events}> {children} </SSEContext.Provider>
}

export const useSSE = (name: string) => {
	const events = useContext(SSEContext)
	return events[name] || null
}
