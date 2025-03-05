import { useSSE } from '@/hooks/useSSE'
import { ensureSpecial, GetSpecialByID, Special, ZeroSpecial } from '@/sdk/sdk'
import { useEffect, useState } from 'react'

export function useSpecial(id: string) {
	const [special, setSpecial] = useState<Special>(ZeroSpecial)

	// get initial data
	useEffect(() => {
		GetSpecialByID(id).then(ensureSpecial).then(setSpecial)
	}, [id])

	const update = useSSE<string>(`special-${id}`).at(-1)

	// get update from sse
	useEffect(() => {
		if (!update) return
		setSpecial(ensureSpecial(JSON.parse(update)))
	}, [update])

	return special
}
