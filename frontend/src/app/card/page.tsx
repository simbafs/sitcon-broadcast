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
		<div className="flex grow flex-col justify-between p-2 font-card">
			<div className="flex items-center justify-between">
				<p className="text-3xl text-[#2540a7]">{card.type}</p>
				<p className="text-2xl text-[#8144b5]">
					{formatTime(card.start)}~{formatTime(card.end)}
				</p>
			</div>
			<p className="text-4xl text-[#000000]">{card.title}</p>
			<p className="text-4xl text-[#917c6a]">{speakers}</p>
		</div>
	)
}

export default function Page() {
	return (
		<>
			<div className="flex min-h-[250px] w-[400px] flex-col bg-[#fdfdfd]">
				<Suspense>
					<Card />
				</Suspense>
			</div>
			<p className="mt-2">400px x 250px</p>
		</>
	)
}
