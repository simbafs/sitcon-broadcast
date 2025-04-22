import { btn } from '@/style/btn'
import { useState } from 'react'
import { CreateEvent as sdkCreateEvent } from '@/sdk/event'
import { toast } from 'react-toastify'
import { twMerge } from 'tailwind-merge'
import { Card } from './card'

export function CreateEvent() {
	const [name, setName] = useState('SITCON2025')
	const [url, setURL] = useState('https://sitcon.org/2025/sessions.json')

	return (
		<Card>
			<h1 className="text-2xl font-semibold">Create Event</h1>
			<div className="grid grid-cols-3 gap-2">
				<input
					type="text"
					value={name}
					onChange={e => setName(e.target.value)}
					placeholder="Event Name"
					className="rounded-lg border p-2"
				/>
				<input
					type="text"
					value={url}
					onChange={e => setURL(e.target.value)}
					placeholder="Event URL"
					className="rounded-lg border p-2"
				/>
				<button
					className={twMerge(btn({ color: 'blue' }), '')}
					onClick={() =>
						sdkCreateEvent(name, url)
							.then(() => toast('已建立活動，請重新整理頁面'))
							.catch((e: Error) => toast(e.message))
					}
				>
					送出
				</button>
			</div>
		</Card>
	)
}
