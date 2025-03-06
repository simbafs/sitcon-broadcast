'use client'

import { Suspense } from 'react'
import Image from 'next/image'
import QRCode from 'react-qr-code'
import { StaticImport } from 'next/dist/shared/lib/get-img-props'
import { formatTime } from '@/utils/formatTime'
import { Session, Special } from '@/sdk/sdk'
import { useCard } from './useCard'

import hackmd from './hackmd.svg'

// TODO: a fucntion to choose icon by the link
import slido from './slido.svg'
import slides from './slides.svg'
import ppt from './ptt.svg'
import google from './google.svg'
import { twMerge } from 'tailwind-merge'
import { useSpecial } from './useSpecial'

function Card({ card, special }: { card: Session; special: Special }) {
	let speakers = card.speakers.join('、')

	if (special.title) card.title = special.title
	if (special.speakers) speakers = special.speakers

	return (
		<div className="font-card flex h-[200px] grow flex-col justify-between overflow-hidden bg-white p-2">
			<div className="flex items-center justify-between">
				<p className="text-3xl text-[#2540a7]">{card.type}</p>
				<p className="text-2xl text-[#8144b5]">
					{formatTime(card.start)}~{formatTime(card.end)}
				</p>
			</div>
			<p className={twMerge('text-3xl text-[#000000]', special.titleStyle)}>{card.title}</p>
			<pre className={twMerge('text-3xl text-[#917c6a]', special.speakersStyle)}>{speakers}</pre>
		</div>
	)
}

function Slido({ card }: { card: Session }) {
	const content = card.slido ? (
		<>
			<Image src={slido} alt="slido" width={40} height={40} />
			<div className="flex items-center">
				<p className="font-card text-3xl text-white">Slido</p>
			</div>
			<p />
			<p className="font-card text-3xl text-[#37ca34]">#{card.slido}</p>
		</>
	) : (
		<div />
	)

	return <div className="grid h-[100px] w-full grid-cols-[40px_1fr] place-content-center gap-2 p-2">{content}</div>
}

function QR({ Icon, title, link }: { Icon: StaticImport; title: string; link: string }) {
	return (
		<div className="flex flex-col items-center gap-4">
			<div className="flex gap-4 ">
				<Image src={Icon} alt="slides" width={40} height={40} />
				<p className="font-card text-3xl text-white">{title}</p>
			</div>
			<div className="grid place-items-center bg-white p-2">
				<QRCode value={link} size={180} />
			</div>
		</div>
	)
}

function Slides({ card }: { card: Session }) {
	return card.slide ? <QR Icon={google} title="簡報" link={card.slide} /> : <div />
}

function Hackmd({ card }: { card: Session }) {
	return card.hackmd ? <QR Icon={hackmd} title="共筆" link={card.hackmd} /> : <div />
}

function SessionCard() {
	const [card, _] = useCard()
	const special = useSpecial(card?.id)

	// WARN: remove this
	// card.slide = 'https://sitcon.org'

	return (
		<div className="w-[420px]">
			<Card card={card} special={special} />
			<div className="h-[151px]" />
			<Slido card={card} />
			<div className="mt-4 grid grid-cols-2">
				<Slides card={card} />
				<Hackmd card={card} />
			</div>
		</div>
	)
}

export default function Page() {
	return (
		<Suspense>
			<SessionCard />
		</Suspense>
	)
}
