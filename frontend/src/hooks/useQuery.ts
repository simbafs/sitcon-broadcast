'use client'
import { useSearchParams } from 'next/navigation'

export default function useQuery<T extends string | null>(key: string, defaultValue: T) {
	const param = useSearchParams()

	let value = param.get(key)

	if (!value) return defaultValue

	return value
}
