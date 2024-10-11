'use client'
import { useSSE } from '@/hooks/useSSE'

export default function Page() {
	const msg = useSSE('time')

	return JSON.stringify(msg, null, 2)
}
