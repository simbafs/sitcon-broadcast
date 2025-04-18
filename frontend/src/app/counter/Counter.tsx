import { Loading } from '@/components/loading'
import { usePolling } from '@/hooks/usePolling'
import { Get } from '@/sdk/counter'
import { btn } from '@/style/btn'
import { twMerge } from 'tailwind-merge'

export function Counter({ name }: { name: string }) {
	const counter = usePolling(() => Get(name), undefined, {
		interval: 500,
	})

	return (
		<div className="grid grid-cols-4 place-items-center gap-2 md:grid-cols-7">
			<h1 className="col-span-2 md:col-span-1">{name}</h1>
			{counter ? (
				<>
					<h1 className="col-span-2 md:col-span-1">{counter.count}</h1>
					<button className={twMerge(btn(), 'w-full')}>開始</button>
					<button className={twMerge(btn(), 'w-full')}>暫停</button>
					<button className={twMerge(btn(), 'w-full')}>停止</button>
					<button className={twMerge(btn(), 'w-full')}>設定</button>
					<button className={twMerge(btn(), 'col-span-4 w-full md:col-span-1')}>開啟頁面</button>
				</>
			) : (
				<Loading />
			)}
		</div>
	)
}
