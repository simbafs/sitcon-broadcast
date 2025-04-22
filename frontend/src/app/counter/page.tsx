'use client'
import { useCounter } from '@/hooks/useCounter'
import { parseAsString, useQueryState } from 'nuqs'
import { Suspense } from 'react'
import { formatCountdown } from './formatTime'
import { Loading } from '@/components/loading'

function CounterPage() {
	const [name] = useQueryState('name', parseAsString.withDefault('R0'))
	const counter = useCounter(name)

	return (
		<div className="grid h-screen w-screen place-items-center">
			{counter ? <h1 className="text-[35vw] leading-[0.8]">{formatCountdown(counter.count)}</h1> : <Loading />}
		</div>
	)
}

export default function Page() {
	return (
		<Suspense>
			<CounterPage />
		</Suspense>
	)
}
