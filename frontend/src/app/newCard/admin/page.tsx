'use client'
import useSWR from 'swr'

export default function Page() {
	const { data, error } = useSWR('/api/session', (url: string) => fetch(url).then(res => res.json()))

	if (error)
		return (
			<>
				<h1>Error</h1>
				<pre>{JSON.stringify(error, null, 2)}</pre>
			</>
		)
	if (!data) return <h1>Loading...</h1>

	return <pre>{JSON.stringify(data, null, 2)}</pre>
}
