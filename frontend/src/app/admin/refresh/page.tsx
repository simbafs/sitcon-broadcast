'use client'
import { useEvent } from '@/app/admin/refresh/useChooseEvent'
import { useEffect, useState } from 'react'
import { useFetch } from '@/hooks/useFetch'
import { CreateEvent } from './createEvent'
import { ShowJSON } from './showJSON'
import { EditScript } from './editScript'
import { ZeroEvent } from '@/sdk/event'

export default function Page() {
	const [SelectEvent, event, setScript, saveEvent] = useEvent()
	const data = useFetch(event?.url || '')
	const [result, setResult] = useState<any>()

	useEffect(() => console.log({ event }), [event])

	return (
		<div className="flex h-screen flex-col gap-4 p-4">
			{/* 橫幅活動選擇區域 */}
			<div className="grid grid-cols-2 gap-4">
				<SelectEvent />
				<CreateEvent />
			</div>

			{/* 三欄內容區域 */}
			<div className="grid h-[calc(100%-8rem)] grid-cols-3 gap-4 ">
				{/* 左側 data 顯示 */}
				<ShowJSON data={data} title="sessions.json" />

				{/* 中間 event.script 輸入區 */}
				<EditScript
					event={event || ZeroEvent}
					setScript={setScript}
					saveEvent={saveEvent}
					result={result}
					setResult={setResult}
					data={data}
				/>

				{/* 右側 result 顯示區 */}
				<ShowJSON data={result} setData={setResult} title="result"/>
			</div>
		</div>
	)
}
