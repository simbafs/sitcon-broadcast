import { Event } from '@/sdk/event'
import { Card } from './card'
import { btn } from '@/style/btn'
import { UpdateAll } from '@/sdk'
import { toast } from 'react-toastify'
import { twMerge } from 'tailwind-merge'
import { sandbox } from './sandbox'

export function EditScript({
	event,
	setScript,
	saveEvent,
	result,
	setResult,
	data,
}: {
	event: Event
	setScript: (script: string) => void
	saveEvent: () => void
	result: any
	setResult: (result: any) => void
	data: any
}) {
	return (
		<Card>
			<div className="flex flex-col gap-2 rounded-lg bg-gray-100 h-full">
				<h2 className="text-lg font-semibold">Script</h2>
				<textarea
					className="text-nowrap flex-1 rounded-lg border p-2"
					value={event?.script}
					onChange={e => setScript(e.target.value)}
				/>
				<div className="grid grid-cols-2 gap-2">
					<button onClick={saveEvent} className={btn({ color: 'green' })}>
						Save Script
					</button>
					<button
						className={btn({ color: 'yellow' })}
						onClick={() =>
							UpdateAll(result)
								.then(() => toast('已更新資料庫'))
								.catch((e: Error) => toast(e.message))
						}
					>
						Save Sessions
					</button>
					<button
						className={twMerge(btn({ color: 'blue' }), 'col-span-2')}
						onClick={() =>
							sandbox(event?.script || '', data)
								.then(setResult)
								.then(() => toast('已執行'))
								.catch((e: Error) => toast(e.message))
						}
					>
						Exec
					</button>
				</div>
			</div>
		</Card>
	)
}
