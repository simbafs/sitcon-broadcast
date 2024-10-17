'use client'
import { edit, useEditor } from '@/components/useEditTime'
import { Countdown, COUNTING, PAUSE, useCountdown } from '@/hooks/useCountdown'
import { formatTime } from '@/utils/formatTime'
import { btn } from '@/varients/btn'
import Link from 'next/link'
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
		<div className="grid grid-cols-6 gap-4">
			<div className="col-span-3 flex items-center justify-center lg:col-span-6">
				<h2 className="text-3xl">{countdown.name}</h2>
			</div>
			<div className="col-span-3 flex items-center justify-center lg:col-span-1">
				<p className="text-3xl">{formatTime(countdown.time)}</p>
			</div>
			<button
				className={twMerge(
					btn({ color: countdown.state === PAUSE ? 'green' : 'normal' }),
					'col-span-3 lg:col-span-1',
				)}
				onClick={countdown.start}
				disabled={countdown.state === COUNTING}
			>
				開始
			</button>
			<button
				className={twMerge(
					btn({ color: countdown.state === COUNTING ? 'red' : 'normal' }),
					'col-span-3 lg:col-span-1',
				)}
				onClick={countdown.pause}
				disabled={countdown.state === PAUSE}
			>
				暫停
			</button>
			<button
				className={twMerge(btn({ color: 'yellow' }), 'col-span-3 sm:col-span-2 lg:col-span-1')}
				onClick={countdown.reset}
			>
				重設
			</button>
			<button
				className={twMerge(btn({ color: 'yellow' }), 'col-span-3 sm:col-span-2 lg:col-span-1')}
				onClick={() => edit(countdown.inittime).then(countdown.setTime)}
			>
				設定時間
			</button>
			<Link
				className={twMerge(btn(), 'col-span-6 sm:col-span-2 lg:col-span-1')}
				href={`/countdown?room=${countdown.name}`}
				target="_blank"
			>
				開啟頁面
			</Link>
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
		<div className="flex flex-col gap-12 divide-y">
			{countdowns.map((c, i) => (
				<Row key={i} countdown={c} edit={edit} />
			))}
		</div>
	)
}

export default function Page() {
	const [Editor, edit] = useEditor<number>(MyEditor, 0)
	return (
		<div className="p-4 lg:px-20 ">
			<Rooms edit={edit} />
			<Editor />
		</div>
	)
}
