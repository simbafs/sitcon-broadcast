import { getNow } from '@/utils/getNow'
import { btn } from '@/varients/btn'
import { twMerge } from 'tailwind-merge'

export function useTime(title: string, time: number, setTime: (value: number) => void) {
	const component = (
		<div
			onClick={e => e.stopPropagation()}
			className="m-2 flex flex-col items-center rounded-md border border-black p-2"
		>
			<p>{title}</p>
			<div className="grid w-full grid-cols-5 gap-2">
				<button className={twMerge(btn({ size: 'md' }))} onClick={() => setTime(time - 5)}>
					-5
				</button>
				<button className={twMerge(btn({ size: 'md' }))} onClick={() => setTime(time - 1)}>
					-1
				</button>
				<button className={twMerge(btn({ size: 'md' }))} onClick={() => setTime(getNow())}>
					Now
				</button>
				<button className={twMerge(btn({ size: 'md' }))} onClick={() => setTime(time + 1)}>
					+1
				</button>
				<button className={twMerge(btn({ size: 'md' }))} onClick={() => setTime(time + 5)}>
					+5
				</button>
			</div>
		</div>
	)

	return component
}
