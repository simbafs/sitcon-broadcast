import { createContext, useContext, useEffect, ReactNode, useRef, useState, useCallback } from 'react'

type Callback = (data: any) => void
type Handlers = Map<string, Set<Callback>>

const SSEContext = createContext<{
	addHandler: (topic: string, callback: Callback) => void
	removeHandler: (topic: string, callback: Callback) => void
}>({
	addHandler: () => {},
	removeHandler: () => {},
})

export const SSEProvider = ({ children, url }: { children: ReactNode; url: string; maxLength?: number }) => {
	const handlersRef = useRef<Handlers>(new Map())

	const addHandler = (topic: string, callback: Callback) => {
		if (!handlersRef.current.has(topic)) {
			handlersRef.current.set(topic, new Set())
		}
		handlersRef.current.get(topic)!.add(callback)
	}

	const removeHandler = (topic: string, callback: Callback) => {
		if (handlersRef.current.has(topic)) {
			handlersRef.current.get(topic)!.delete(callback)
			if (handlersRef.current.get(topic)!.size === 0) {
				handlersRef.current.delete(topic)
			}
		}
	}

	useEffect(() => {
		const eventSource = new EventSource(url)

		eventSource.onmessage = event => {
			try {
				const parsedData: { topic: string[]; data: any } = JSON.parse(event.data)
				const { topic, data } = parsedData
				const callbacks = handlersRef.current.get(topic[0])
				if (callbacks) {
					callbacks.forEach(callback => {
						callback(data)
					})
				}

				const allHandlers = handlersRef.current.get('__all__')
				if (allHandlers) {
					allHandlers.forEach(callback => {
						callback({ topic, data })
					})
				}
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

	return (
		<SSEContext.Provider
			value={{
				addHandler,
				removeHandler,
			}}
		>
			{children}
		</SSEContext.Provider>
	)
}

export function useSSE(topic: string, callback: Callback) {
	const { addHandler, removeHandler } = useContext(SSEContext)

	useEffect(() => {
		addHandler(topic, callback)

		return () => {
			removeHandler(topic, callback)
		}
	}, [topic, addHandler, removeHandler, callback])
}

export function useSSEFetch<T>(topic: string, callback: Callback, init: () => Promise<T>) {
	useSSE(topic, callback)

	useEffect(() => {
		init().then(data => callback(data))
	}, [callback, init, topic])
}

export function useSSEValue<T>(topic: string) {
	const [value, setValue] = useState<T>()
	useSSE(topic, setValue)
	return value
}

export function useSSEFetchValue<T>(topic: string, init: () => Promise<T>) {
	const [value, setValue] = useState<T>()
	useSSEFetch(topic, setValue, init)
	return value
}

export function useAll(callback: Callback) {
	const { addHandler, removeHandler } = useContext(SSEContext)

	useEffect(() => {
		addHandler('__all__', callback)

		return () => {
			removeHandler('__all__', callback)
		}
	}, [addHandler, removeHandler, callback])
}
