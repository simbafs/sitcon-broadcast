'use client'

import useQuery from '@/hooks/useQuery'
import useSWR from 'swr'
import { useEffect } from 'react'
import x from '@/img/x.svg'
import Image from 'next/image'
import { twMerge } from 'tailwind-merge'
import { InvalidURL } from './InvalidURL'
import type { Session } from '@/types/card'

export default function Page() {
	const sessionId = useQuery('id')
	const room = useQuery('room')
	const api_url = sessionId != '' ? `/api/card/${sessionId}` : `/api/card/room/${room}`
	const { data, error } = useSWR<Session>(api_url, url => fetch(url).then(res => res.json()), {
		refreshInterval: 500,
	})

	useEffect(() => console.log(data), [data])

	if (room == '' && sessionId == '') return <InvalidURL />

	if (error)
		return (
			<>
				<h1>Error</h1>
				<pre>{JSON.stringify(error, null, 2)}</pre>
			</>
		)

	if (!data) return <h1>Loading...</h1>

	const speaker = data.speakers.join('、')

	return (
		<div className="h-screen w-screen overflow-hidden bg-transparent">
			<div className="flex aspect-[1.8/1] w-[70vw] flex-col bg-[#f7f6f6] shadow-[18px_18px_50px_0px_rgba(0,0,0,0.1)]">
				<div className="flex h-[6vw] items-center justify-end bg-[#406096]">
					<Image src={x} width={18} height={18} alt="Close" className="mr-[2vw] h-[4vw] w-[4vw]" />
				</div>
				<div className="flex grow flex-col px-[6vw] py-[2vw]">
					<h1 className="text-[5vw] text-[#9f3b24]">{data.zh.title}</h1>
					<div className="flex grow flex-wrap items-center">
						<h2 className="text-[4vw] text-[#383839]">
							{data.start}-{data.end}
						</h2>
						<span className="grow" />
						<h2 className={twMerge('text-[#383839]', speaker.length > 60 ? 'text-[3vw]' : 'text-[4vw]')}>
							{speaker}
						</h2>
					</div>
				</div>
			</div>
		</div>
	)
}
