'use client'
import { Counter } from './Counter'
import { useCounters } from '@/hooks/useCounters'
import { Loading } from '@/components/loading'

export default function Page() {
	const counters = useCounters()

	return (
		<div className="mx-20 mt-10 flex flex-col">
			<h1 className="text-2xl">Counters</h1>
			{counters ? (
				<div className="flex flex-col gap-2">
					{Object.keys(counters).map(c => (
						<Counter key={c} name={c} />
					))}
				</div>
			) : (
				<Loading />
			)}
		</div>
	)
}
