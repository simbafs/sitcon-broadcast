import { formatTime } from '@/utils/formatTime'
import { btn } from '@/varients/btn'
import { twMerge } from 'tailwind-merge'

function addMinute(base: Date, offset: number) {
	base = new Date(base)
	base.setMinutes(base.getMinutes() + offset)
	return base
}

export function EditTime({ title, time, setTime }: { title: string; time: Date; setTime: (value: Date) => void }) {
	const component = (
		<div
			onClick={e => e.stopPropagation()}
			className="m-2 flex flex-col items-center rounded-md border-2 border-gray-400 p-2"
		>
			<p>
				{title}: {formatTime(time)}
			</p>
			<div className="grid w-full grid-cols-5 gap-2">
				<button className={twMerge(btn({ size: 'md' }))} onClick={() => setTime(addMinute(time, -5))}>
					-5
				</button>
				<button className={twMerge(btn({ size: 'md' }))} onClick={() => setTime(addMinute(time, -1))}>
					-1
				</button>
				<button className={twMerge(btn({ size: 'md' }))} onClick={() => setTime(new Date())}>
					Now
				</button>
				<button className={twMerge(btn({ size: 'md' }))} onClick={() => setTime(addMinute(time, +1))}>
					+1
				</button>
				<button className={twMerge(btn({ size: 'md' }))} onClick={() => setTime(addMinute(time, +5))}>
					+5
				</button>
			</div>
		</div>
	)

	return component
}
