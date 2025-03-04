'use client'

import Image from 'next/image'
import { twMerge } from 'tailwind-merge'
import x from './x.svg'
import { formatTime } from '@/utils/formatTime'
import { useSSE } from '@/hooks/useSSE'
import { Suspense, useEffect, useState } from 'react'
import { parseAsString, useQueryState } from 'nuqs'
import { ensureSession, GetCurrentSession, GetSessionByID, Session, ZeroSession } from '@/sdk/sdk'

function useCard() {
	const [room] = useQueryState('room', parseAsString.withDefault('R0'))
	const [id] = useQueryState('id')
	const [card, setCard] = useState(ZeroSession)

	const keepUpdate = id == null
	const sseStr = keepUpdate ? `card-current-${room}` : `card-id-${id}`

	const update = useSSE<Session>(sseStr).at(-1)
	const [error, setError] = useState<Error>()

	// init
	useEffect(() => {
		if (keepUpdate) GetCurrentSession(room).then(setCard, setError)
		else GetSessionByID(id).then(setCard, setError)
	}, [id, keepUpdate, room])

	useEffect(() => {
		if (!update) return
		setCard(ensureSession(update))
	}, [update])

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
