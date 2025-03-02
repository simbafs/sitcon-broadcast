'use client'
import { createContext, useContext, useEffect, ReactNode, useReducer, useState } from 'react'

type Events = Record<string, any[]>

const SSEContext = createContext<Events>({})

export const SSEProvider = ({
	children,
	url,
	maxLength = 100,
}: {
	children: ReactNode
	url: string
	maxLength?: number
}) => {
	const [events, updateEvents] = useReducer((state: Events, action: { name: string; data: any }) => {
		if (action.name in state) {
			return {
				...state,
				[action.name]: [...state[action.name], action.data].slice(-maxLength),
			}
		} else {
			return {
				...state,
				[action.name]: [action.data].slice(-maxLength),
			}
		}
	}, {})

	useEffect(() => {
		const eventSource = new EventSource(url)

		eventSource.onmessage = event => {
			try {
				const parsedData: { name: string; data: any } = JSON.parse(event.data)
				updateEvents(parsedData)
			} catch (err) {
				console.error('Error parsing SSE data:', err)
			}
		}

		eventSource.onerror = err => {
			console.error('EventSource failed:', err)
		}

		eventSource.onopen = () => {
			console.log('SSE connected')
		}

		return () => {
			eventSource.close()
		}
	}, [url])

	return <SSEContext.Provider value={events}> {children} </SSEContext.Provider>
}

export function useSSE<T>(name: string): T[] {
	const events = useContext(SSEContext)

	return events[name] || []
}

export function useAllSSE() {
	return useContext(SSEContext)
}

export function useSSEFetch<T>(name: string, init: () => Promise<T>, deps: any[] = []) {
	const [data, setData] = useState<T | undefined>(undefined)
	const latest = useSSE<T>(name).at(-1)

	useEffect(() => {
		init().then(setData)
		// eslint-disable-next-line react-hooks/exhaustive-deps
	}, [...deps])

	return latest || data
}
