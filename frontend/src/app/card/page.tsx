'use client'

import { formatTime } from '@/utils/formatTime'
import { useSSE } from '@/hooks/useSSE'
import { Suspense, useEffect, useState } from 'react'
import { parseAsString, useQueryState } from 'nuqs'
import { ensureSession, GetCurrentSession, GetSessionByID, Session, ZeroSession } from '@/sdk/sdk'
import QRCode from 'react-qr-code'
import Image from 'next/image'

import slido from './slido.png'

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

function Card({ card }: { card: Session }) {
	const speakers = card.speakers.join('、')

	return (
		<div className="font-card flex grow flex-col justify-between p-2">
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

function Slido({ card }: { card: Session }) {
	if (!card.slido) return <p>No Slido</p>
	return (
		<div className="grid grid-cols-[40px_1fr] place-content-center gap-2 bg-[#9698c8] h-full w-full p-2">
			<Image src={slido} alt="slido" width={40} height={40} />
			<div className="flex items-center">
				<p className="text-white text-3xl font-card">Slido</p>
			</div>
			<p />
			<p className="text-[#37ca34] text-3xl font-card">#{card.slido}</p>
		</div>
	)
}

function Slides({ card }: { card: Session }) {
	if (!card.slide) return <p>No Image</p>
	return <QRCode value={card.slide} />
}

function Hackmd({ card }: { card: Session }) {
	if (!card.hackmd) return <p>No Hackmd</p>
	return <QRCode value={card.hackmd} />
}

function SessionCard() {
	const [card, _] = useCard()

	return (
		<>
			<div className="w-fit border-4 border-green-500">
				<div className="flex min-h-[250px] w-[400px] flex-col bg-[#fefefe]">
					<Card card={card} />
				</div>
			</div>
			<p className="mt-2">400px x 250px</p>

			<div className=" w-fit border-4 border-green-500 mt-2">
				<div className="h-[100px] w-[350px]">
					<Slido card={card} />
				</div>
			</div>
			<p className="mt-2">350px x 100px</p>
			<div className="mt-2 grid w-fit grid-cols-2 gap-2">
				<div className=" w-fit border-4 border-green-500">
					<div className="grid h-[256px] w-[256px] place-content-center">
						<Slides card={card} />
					</div>
				</div>
				<div className="w-fit border-4 border-green-500">
					<div className="grid h-[256px] w-[256px] place-content-center">
						<Hackmd card={card} />
					</div>
				</div>
				<p>256px x 256px</p>
				<p>256px x 256px</p>
			</div>
		</>
	)
}

export default function Page() {
	return (
		<Suspense>
			<SessionCard />
		</Suspense>
	)
}
