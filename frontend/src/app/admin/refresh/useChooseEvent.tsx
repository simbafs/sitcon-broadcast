import { Event, GetAll, SetScript } from '@/sdk/event'
import { useEffect, useState } from 'react'
import { toast } from 'react-toastify'

export function useEvent() {
	const [events, setEvents] = useState<Event[]>([])
	const [currentEvent, setCurrentEvent] = useState<Event | undefined>()

	useEffect(() => {
		GetAll()
			.then(e => {
				setEvents(e)
				setCurrentEvent(e[0])
			})
			.catch(() => toast('無法取得活動'))
	}, [])

	const saveEvent = () => {
		if (!currentEvent) return
		SetScript(currentEvent.name, currentEvent.script)
			.then(() => toast('已儲存'))
			.catch(() => toast('無法儲存'))
	}
	const setScript = (script: string) =>
		setCurrentEvent(e => {
			if (!e) return e
			return { ...e, script }
		})

	return [
		() => (
			<select onChange={e => setCurrentEvent(events[+e.target.value])}>
				{events.map((e, i) => (
					<option key={i} value={i}>
						{e.name}
					</option>
				))}
			</select>
		),
		currentEvent,
		setScript,
		saveEvent,
	] as const
}
