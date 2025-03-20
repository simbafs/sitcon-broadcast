'use client'
import { useEvent } from '@/hooks/useChooseEvent'
import { btn } from '@/style/btn'
import { sandbox } from './sandbox'
import { useState } from 'react'
import { useFetch } from '@/hooks/useFetch'
import { twMerge } from 'tailwind-merge'

export default function Page() {
	const [SelectEvent, event, setScript, saveEvent] = useEvent()
	const data = useFetch(event.url)
	const [result, setResult] = useState<any>()

	return (
		<div className="flex h-screen flex-col gap-4 p-4">
			{/* 橫幅活動選擇區域 */}
			<div className="rounded-lg bg-gray-200 p-4 shadow-md">
				<h1 className="text-2xl font-semibold">Refresh Database</h1>
				<SelectEvent />
			</div>

			{/* 三欄內容區域 */}
			<div className="grid h-[calc(100%-6rem)] grow grid-cols-3 gap-4 ">
				{/* 左側 data 顯示 */}
				<div className="flex h-full flex-col overflow-auto rounded-lg bg-gray-100 p-4">
					<h2 className="text-lg font-semibold">Data</h2>
					<pre className="h-full overflow-auto whitespace-pre-wrap break-words">
						{JSON.stringify(data, null, 2)}
					</pre>
				</div>

				{/* 中間 event.script 輸入區 */}
				<div className="flex flex-col gap-2 rounded-lg bg-gray-100 p-4">
					<h2 className="text-lg font-semibold">Script</h2>
					<textarea
						className="flex-1 rounded-lg border p-2"
						value={event.script}
						onChange={e => setScript(e.target.value)}
					/>
					<div className="grid grid-cols-2 gap-2">
						<button onClick={saveEvent} className={btn({ color: 'green' })}>
							Save Script
						</button>
						<button className={btn({ color: 'yellow' })} disabled>
							Save Sessions
						</button>
						<button
							className={twMerge(btn({ color: 'blue' }), 'col-span-2')}
							onClick={() => sandbox(event.script, data).then(setResult)}
						>
							Exec
						</button>
					</div>
				</div>

				{/* 右側 result 顯示區 */}
				<div className="flex h-full flex-col overflow-auto rounded-lg bg-gray-100 p-4">
					<h2 className="text-lg font-semibold">Result</h2>
					<pre className="h-full overflow-auto whitespace-pre-wrap break-words">
						{JSON.stringify(result, null, 2)}
					</pre>
				</div>
			</div>
		</div>
	)
}
