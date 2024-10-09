'use client'

import Image from 'next/image'
import { twMerge } from 'tailwind-merge'
import { useSearchParams } from 'next/navigation'
import x from '@/img/x.svg'
import useSWR from 'swr'
import { formatTime } from '@/utils/formatTime'

export default function Page() {
	const searchParams = useSearchParams()

	const room = searchParams.get('room') || 'R0'
	let idx = searchParams.get('idx')

	if (!idx) {
		// TODO: get id by time
		idx = '0'
	}

	// TODO: use websocket to revalidate
	const { data, error } = useSWR(`/api/card/${room}/${idx}`, (url: string) => fetch(url).then(res => res.json()), {
		refreshInterval: 500,
	})

	if (error)
		return (
			<>
				<h1>Error</h1>
				<pre>{JSON.stringify(error, null, 2)}</pre>
			</>
		)

	if (!data) return <h1>Loading...</h1>

	const speakers = data.speakers.join('、')
	return (
		<div className="w-screen h-screen bg-transparent overflow-hidden">
			<div className="aspect-[1.8/1] w-[70vw] bg-[#f7f6f6] flex flex-col shadow-[18px_18px_50px_0px_rgba(0,0,0,0.1)]">
				<div className="bg-[#406096] h-[6vw] flex justify-end items-center">
					<Image src={x} width={18} height={18} alt="Close" className="h-[4vw] w-[4vw] mr-[2vw]" />
				</div>
				<div className="px-[6vw] py-[2vw] grow flex flex-col">
					<h1 className="text-[#9f3b24] text-[5vw]">{data.title}</h1>
					<div className="flex flex-wrap grow items-center">
						<h2 className="text-[#383839] text-[4vw]">
							{formatTime(data.start)}-{formatTime(data.end)}
						</h2>
						<span className="grow" />
						<h2 className={twMerge('text-[#383839]', speakers.length > 60 ? 'text-[3vw]' : 'text-[4vw]')}>
							{speakers}
						</h2>
					</div>
				</div>
			</div>
		</div>
	)
}
