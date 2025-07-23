import { useState, useEffect } from 'react'
import { Display } from 'controly'

interface UseDisplayResult {
	display: Display | null
	id: string
	showQRcode: boolean
}

export const useDisplay = (): UseDisplayResult => {
	const [display, setDisplay] = useState<Display | null>(null)
	const [id, setID] = useState('')
	const [showQRcode, setShowQRcode] = useState(true)

	useEffect(() => {
		const urlParams = new URLSearchParams(window.location.search)
		const initialId = urlParams.get('id')

		const displayInstance = new Display({
			serverUrl: 'wss://controly.1li.tw/ws',
			commandUrl: `${window.location.origin}${window.location.pathname}/command.json`,
			id: initialId || undefined,
		})

		displayInstance.on('subscribed', payload => {
			setShowQRcode(payload.count === 0)
		})
		displayInstance.on('unsubscribed', payload => {
			setShowQRcode(payload.count === 0)
		})
		displayInstance.on('open', newId => {
			setID(newId)
			// Update URL without reloading
			const newUrl = new URL(window.location.href)
			newUrl.searchParams.set('id', newId)
			window.history.replaceState({}, '', newUrl.toString())
		})
		displayInstance.on('error', err => console.error('Display Error:', err))

		displayInstance.connect()
		setDisplay(displayInstance)

		return () => {
			display?.disconnect()
		}
		// eslint-disable-next-line react-hooks/exhaustive-deps
	}, [])

	return { display, id, showQRcode }
}
