'use client'
import useQuery from '@/hooks/useQuery'
import { formatTime } from '@/utils/formatTime'
import { useState } from 'react'
import { useInterval } from 'usehooks-ts'

export default function Home() {
	const roomid = useQuery('id', '0')
	const [time, setTime] = useState(0)

	useInterval(() => {
		fetch(`/api/now/${roomid}/`)
			.then(res => res.json())
			.then((data: { now: number }) => {
				setTime(data.now)
			})
	}, 200)

	return (
		<div className="w-screen h-screen grid place-items-center">
			<h1 className="text-[35vw] leading-[0.8]">{formatTime(time)}</h1>
		</div>
	)
}
