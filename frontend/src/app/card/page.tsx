'use client'

import Image from 'next/image'
import { twMerge } from 'tailwind-merge'
import x from '@/img/x.svg'
import { formatTime } from '@/utils/formatTime'
import { useSSE } from '@/hooks/useSSE'
import { Suspense, useEffect, useState } from 'react'
import { Session, ZeroSession } from '@/types/card'
import { parseAsString, useQueryState } from 'nuqs'

function useCard() {
	const [room] = useQueryState('room', parseAsString.withDefault('R0'))
	const [idx] = useQueryState('idx')

	const needUpdate = idx === null
	let url = `/api/card/${room}/${idx}`

	if (needUpdate) {
		url = `/api/card/${room}`
	}

	const [card, setCard] = useState(ZeroSession)

	const latest = useSSE<Session>(`card-%{room}`).at(-1)
	const [error, setError] = useState<Error>()

	// init
	useEffect(() => {
		fetch(url)
			.then(res => res.json())
			.then(data => {
				if (data.error) {
					throw new Error(data.error)
				}
				return data
			})
			.then(setCard)
			.catch(e => setError(new Error(`Failed to fetch card: ${e}`)))
	}, [url])

	// update
	// TODO: test if updating works
	useEffect(() => {
		if (!needUpdate) return
		if (!latest) return
		setCard(latest)
	}, [latest, needUpdate])

	return [card, error] as const
}

function Card() {
	const [card, error] = useCard()

	const speakers = card.speakers.join('、')

	// TODO: should I show the error?
	// if (error) {
	// 	return (
	// 		<div className="flex grow flex-col px-[6vw] py-[2vw]">
	// 			<h1 className="text-[5vw] text-[#9f3b24]">Error</h1>
	// 			<pre>{error.message}</pre>
	// 		</div>
	// 	)
	// }

	return (
		<div className="flex grow flex-col px-[6vw] py-[2vw]">
			<h1 className="text-[5vw] text-[#9f3b24]">{card.title}</h1>
			<div className="flex grow flex-wrap items-center">
				<h2 className="text-[4vw] text-[#383839]">
					{formatTime(card.start)}-{formatTime(card.end)}
				</h2>
				<span className="grow" />
				<h2 className={twMerge('text-[#383839]', speakers.length > 60 ? 'text-[3vw]' : 'text-[4vw]')}>
					{speakers}
				</h2>
			</div>
		</div>
	)
}

export default function Page() {
	return (
		<div className="h-screen w-screen overflow-hidden bg-transparent">
			<div className="flex aspect-[1.8/1] w-[70vw] flex-col bg-[#f7f6f6] shadow-[18px_18px_50px_0px_rgba(0,0,0,0.1)]">
				<div className="flex h-[6vw] items-center justify-end bg-[#406096]">
					<Image src={x} width={18} height={18} alt="Close" className="mr-[2vw] h-[4vw] w-[4vw]" />
				</div>
				<Suspense>
					<Card />
				</Suspense>
			</div>
		</div>
	)
}
