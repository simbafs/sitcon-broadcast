'use client'
import { usePolling } from '@/hooks/usePolling'
import { GetAll } from '@/sdk/counter'
import { Counter } from '../Counter'

export default function Page() {
	const counters = usePolling(() => GetAll().then(Object.keys), [], {
		interval: 1 * 60 * 1000, // 1 min
	})

	return (
		<div className="mx-20 mt-10 flex flex-col">
			<h1 className="text-2xl">Counters</h1>
			<div className="flex flex-col gap-2">
				{counters.map(c => (
					<Counter key={c} name={c} />
				))}
			</div>
		</div>
	)
}
