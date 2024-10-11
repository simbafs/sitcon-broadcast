'use client'
import { useSearchParams } from 'next/navigation'

export default function useQuery(key: string, defaultValue: string = '') {
	const param = useSearchParams()

	let value = param.get(key)

	if (!value) return defaultValue

	return value
}
