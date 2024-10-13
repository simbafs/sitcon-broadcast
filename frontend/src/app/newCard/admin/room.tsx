import { Session } from '@/types/card'
import { formatTime } from '@/utils/formatTime'
import { getNow } from './getNow'
import { useState } from 'react'
import { btn } from '@/varients/btn'
import { twMerge } from 'tailwind-merge'

export function Room({ sessions }: { sessions: Session[] }) {
	const now = getNow(sessions)
	const [current, setCurrent] = useState(() => getNow(sessions))
	return (
		<>
			<div className="m-4 flex w-full justify-around">
				<button className={btn()} onClick={() => current - 1 >= 0 && setCurrent(current - 1)}>
					↑
				</button>
				<button className={btn()} onClick={() => setCurrent(now)}>
					Now
				</button>
				<button className={btn()} onClick={() => current + 1 < sessions.length && setCurrent(current + 1)}>
					↓
				</button>
			</div>

			<ul className="m-2 flex w-full flex-col justify-stretch gap-4">
				{/* previous */}
				<EditSession session={sessions[current - 1]} isCurrent={now == current - 1} key={current - 1} />
				{/* current */}
				<EditSession session={sessions[current]} isCurrent={now == current} key={current} />
				{/* next */}
				<button className={twMerge(btn())}>新增議程(Unimplemented)</button>
				<EditSession session={sessions[current + 1]} isCurrent={now == current + 1} key={current + 1} />
			</ul>
		</>
	)
}

function EditSession({ session, isCurrent }: { session?: Session; isCurrent: boolean }) {
	const content = session ? (
		<>
			<h1>{session.title}</h1>
			<p>
				{formatTime(session.start)} - {formatTime(session.end)}
			</p>
		</>
	) : (
		<p>Empty</p>
	)
	return (
		<div
			data-isCurrent={isCurrent}
			className="rounded-lg bg-gray-100 p-2 shadow-lg data-[isCurrent=true]:bg-gray-200 "
		>
			{content}
		</div>
	)
}
