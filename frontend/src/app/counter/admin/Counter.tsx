import { Loading } from '@/components/loading'
import { useCounter } from '@/hooks/useCounter'
import { Reset, SetInit, Start, Stop } from '@/sdk/counter'
import { btn } from '@/style/btn'
import { formatCounter } from '../formatTime'

export function Counter({ name }: { name: string }) {
	const counter = useCounter(name)

	return (
		<div className="grid grid-cols-4 place-items-center gap-2 md:grid-cols-7">
			<h1 className="col-span-2 md:col-span-1">{name}</h1>
			{counter ? (
				<>
					<h1 className="col-span-2 md:col-span-1">{formatCounter(counter.count)}</h1>
					<button onClick={() => Start(name)} className={btn({ class: 'w-full' })}>
						開始
					</button>
					<button onClick={() => Stop(name)} className={btn({ class: 'w-full' })}>
						停止
					</button>
					<button onClick={() => Reset(name)} className={btn({ class: 'w-full' })}>
						重設
					</button>
					<button
						onClick={() => {
							console.log('h')
							const t = prompt('請輸入初始值（秒）')
							console.log(t)
							if (t) SetInit(name, +t)
						}}
						className={btn({ class: 'w-full' })}
					>
						設定
					</button>
					<a
						href={`/counter?room=${name}`}
						target="_blank"
						className={btn({ class: 'col-span-4 w-full text-center md:col-span-1' })}
					>
						開啟頁面
					</a>
				</>
			) : (
				<Loading />
			)}
		</div>
	)
}
