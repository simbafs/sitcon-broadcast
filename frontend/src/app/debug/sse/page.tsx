'use client'
import { useSSE } from '@/hooks/util/useSSE'
import { useState } from 'react'

export default function Page() {
	const [events, setData] = useState<string[]>([])

	useSSE('__all__', data => {
		if (data.topic.includes('now')) return
		setData(prev => [...prev, JSON.stringify(data, null, 2)].slice(-10))
	})

	return <pre>{events.join('\n---\n')}</pre>
}
