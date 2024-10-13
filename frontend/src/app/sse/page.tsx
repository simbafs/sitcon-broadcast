'use client'
import { useAllSSE } from '@/hooks/useSSE'

export default function Page() {
	const msg = useAllSSE()

	if (!msg) return <div>loading...</div>

	return (
		<>
			<h1 className="text-3xl">Server Sent Events</h1>
			<div className="m-4 flex gap-4">
				{Object.entries(msg).map(([name, msgs]) => (
					<div key={name} className="rounded border border-black p-4">
						<h2>{name}</h2>
						<ul>
							{msgs.map((msg, i) => (
								<li key={i}>{JSON.stringify(msg)}</li>
							))}
						</ul>
					</div>
				))}
			</div>
		</>
	)
}
