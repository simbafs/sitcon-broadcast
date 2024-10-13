'use client'
import { type Countdown, COUNTING, PAUSE, useCountdown } from '@/hooks/useCountdown'
import { btn } from '@/varients/btn'
import Link from 'next/link'
import { formatTime } from '@/utils/formatTime'
import { edit, useEditor } from '@/components/useEditTime'
import { useState } from 'react'
import { twMerge } from 'tailwind-merge'

function MyEditor({
	isOpen,
	initValue,
	callback,
}: {
	isOpen: boolean
	initValue: number
	callback: (value: number, ok: boolean) => void
}) {
	const [hour, setHour] = useState(Math.floor(initValue / 60).toString())
	const [minute, setMinute] = useState((initValue % 60).toString())

	return (
		<div
			className={twMerge(
				isOpen ? 'grid' : 'hidden',
				'fixed left-0 top-0 h-screen w-screen place-items-center bg-black/50',
			)}
			onClick={() => callback(0, false)}
		>
			<form
				className="flex w-80 flex-col items-center gap-4 rounded-lg bg-white px-8 py-16"
				onClick={e => e.stopPropagation()}
				onSubmit={e => {
					e.preventDefault()
					callback(Number(hour) * 60 + Number(minute), true)
				}}
			>
				<h1>Set Time</h1>
				<div className="w-fulla flex gap-4">
					<input
						className="w-full rounded-lg border-2 border-gray-500 p-1 outline-none focus:border-blue-500"
						type="txt"
						value={hour}
						onChange={e => setHour(e.target.value)}
						inputMode="decimal"
						onFocus={e => e.target.select()}
						autoFocus
						tabIndex={1}
					/>
					:
					<input
						className="w-full rounded-lg border-2 border-gray-500 p-1 outline-none focus:border-blue-500"
						type="txt"
						value={minute}
						onChange={e => setMinute(e.target.value)}
						inputMode="decimal"
						onFocus={e => e.target.select()}
						tabIndex={2}
					/>
				</div>
				<div className="flex w-full gap-4">
					<button
						className="w-full rounded-md bg-blue-500 p-2 font-bold text-white"
						type="submit"
						tabIndex={3}
					>
						Save
					</button>
					<button
						className="w-full rounded-md bg-gray-500 p-2 font-bold text-white"
						onClick={() => callback(0, false)}
						type="button"
						tabIndex={4}
					>
						Cancel
					</button>
				</div>
			</form>
		</div>
	)
}

function Row({ countdown, edit }: { countdown: Countdown; edit: edit<number> }) {
	return (
		<div className="grid grid-cols-1 gap-4 lg:grid-cols-[2fr_4fr]">
			<div className="grid grid-cols-2 gap-6">
				<div className="flex items-center justify-center">
					<h2 className="text-3xl">{countdown.name}</h2>
				</div>
				<div className="flex items-center justify-center">
					<p className="text-3xl">{formatTime(countdown.time)}</p>
				</div>
			</div>
			<div className="grid grid-cols-5 gap-6">
				<button
					className={btn({ color: countdown.state === PAUSE ? 'green' : 'normal' })}
					onClick={countdown.start}
					disabled={countdown.state === COUNTING}
				>
					開始
				</button>
				<button
					className={btn({ color: countdown.state === COUNTING ? 'red' : 'normal' })}
					onClick={countdown.pause}
					disabled={countdown.state === PAUSE}
				>
					暫停
				</button>
				<button className={btn({ color: 'yellow' })} onClick={countdown.reset}>
					重設
				</button>
				<button
					className={btn({ color: 'yellow' })}
					onClick={() => edit(countdown.inittime).then(countdown.setTime)}
				>
					設定時間
				</button>
				<Link className={btn()} href={`/countdown?room=${countdown.name}`} target="_blank">
					開啟頁面
				</Link>
			</div>
		</div>
	)
}

function Rooms({ edit }: { edit: edit<number> }) {
	const countdowns = [
		useCountdown('R0'),
		useCountdown('R1'),
		useCountdown('R2'),
		useCountdown('R3'),
		useCountdown('S'),
	]

	return (
		<div className="grid w-full gap-[50px]">
			{countdowns.map((c, i) => (
				<Row key={i} countdown={c} edit={edit} />
			))}
		</div>
	)
}

export default function Page() {
	const [Editor, edit] = useEditor<number>(MyEditor, 0)
	return (
		<div className="flex min-h-screen w-screen flex-col items-center justify-center px-[50px] py-[100px] lg:px-[100px]">
			<Rooms edit={edit} />
			<Editor />
		</div>
	)
}
