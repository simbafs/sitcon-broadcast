'use client'
import { useCounter } from '@/hooks/useCounter'
import { formatCountdown } from './formatTime'
import { Loading } from '@/components/loading'
import { useQuery } from '@/hooks/util/useQuery'

export default function Page() {
	const [name] = useQuery('room', 'R0')
	const counter = useCounter(name)

	return (
		<div className="grid h-screen w-screen place-items-center">
			{counter ? <h1 className="text-[35vw] leading-[0.8]">{formatCountdown(counter.count)}</h1> : <Loading />}
		</div>
	)
}
