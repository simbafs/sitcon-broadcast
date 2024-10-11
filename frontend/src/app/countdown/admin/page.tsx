'use client'
import { type Countdown, COUNTING, PAUSE, useCountdown } from '@/hooks/useCountdown'
import { btn } from '@/varients/btn'
import Link from 'next/link'
import { setEditor, useEditTime } from '@/components/useEditTime'
import { formatTime } from '@/utils/formatTime'

function Row({ countdown, setTimeEditor }: { countdown: Countdown; setTimeEditor: setEditor }) {
	const setTime = () => {
		setTimeEditor(formatTime(countdown.inittime)).then(time => {
			const [m, s] = time.split(':').map(Number)
			const t = m * 60 + s
			countdown.setTime(t)
		})
	}
	return (
		<div className="grid gap-4 grid-cols-1 lg:grid-cols-[2fr_4fr]">
			<div className="grid grid-cols-2 gap-6">
				<div className="flex justify-center items-center">
					<h2 className="text-3xl">{countdown.name}</h2>
				</div>
				<div className="flex justify-center items-center">
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
				<button className={btn({ color: 'yellow' })} onClick={setTime}>
					設定時間
				</button>
				<Link className={btn()} href={`/countdown?room=${countdown.name}`} target="_blank">
					開啟頁面
				</Link>
			</div>
		</div>
	)
}

function Rooms({ setTimeEditor }: { setTimeEditor: setEditor }) {
	const countdowns = [
		useCountdown('R0'),
		useCountdown('R1'),
		useCountdown('R2'),
		useCountdown('R3'),
		useCountdown('S'),
	]

	return (
		<div className="w-full grid gap-[50px]">
			{countdowns.map((c, i) => (
				<Row key={i} countdown={c} setTimeEditor={setTimeEditor} />
			))}
		</div>
	)
}

export default function Page() {
	const [TimeEditor, setTimeEditor] = useEditTime()

	return (
		<>
			<TimeEditor />
			<div className="min-h-screen w-screen py-[100px] px-[50px] lg:px-[100px] flex flex-col justify-center items-center">
				<Rooms setTimeEditor={setTimeEditor} />
			</div>
		</>
	)
}
