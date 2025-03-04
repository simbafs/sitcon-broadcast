import { formatTime } from '@/utils/formatTime'
import { useState } from 'react'
import { btn } from '@/varients/btn'
import { EditTime } from '@/app/card/admin/useTime'
import { useMediaQuery } from 'usehooks-ts'
import { twMerge } from 'tailwind-merge'
import { Session, UpdateSession } from '@/sdk/sdk'
import { getCurrent } from './getCurrent'

export function Room({ sessions, room }: { sessions: Session[]; room: string }) {
	const now = getCurrent(sessions) // get current session id
	const [current, setCurrent] = useState(() => getCurrent(sessions))

	return (
		<div className="mt-4 grid w-full gap-4 md:grid-cols-[100px_1fr]">
			<div className="grid w-full grid-cols-3 gap-2 md:grid-cols-1">
				<button className={btn({ size: 'xl' })} onClick={() => current - 1 >= 0 && setCurrent(current - 1)}>
					↑
				</button>
				<button className={btn({ size: 'xl' })} onClick={() => setCurrent(now)}>
					Now
				</button>
				<button
					className={btn({ size: 'xl' })}
					onClick={() => current + 1 < sessions.length && setCurrent(current + 1)}
				>
					↓
				</button>
			</div>

			<ul className="flex w-full flex-col justify-stretch gap-4">
				{/* previous */}
				<EditSession
					idx={current - 1}
					room={room}
					sessions={sessions}
					isCurrent={now == current - 1}
					key={current - 1}
				/>
				{/* current */}
				<EditSession idx={current} room={room} sessions={sessions} isCurrent={now == current} key={current} />
				{/* <button className={twMerge(btn())}>新增議程(Unimplemented)</button> */}
				{/* next */}
				<EditSession
					idx={current + 1}
					room={room}
					sessions={sessions}
					isCurrent={now == current + 1}
					key={current + 1}
				/>
			</ul>
		</div>
	)
}

function EditSession({
	idx,
	room,
	sessions,
	isCurrent,
}: {
	idx: number
	room: string
	sessions: Session[]
	isCurrent: boolean
}) {
	// return <pre>{JSON.stringify({ s: sessions[idx] }, null, 2)}</pre>
	const session = sessions[idx] as Session | undefined
	const [detail, setDetail] = useState(false)
	const isMd = useMediaQuery('(min-width: 768px)')

	// TODO: case of the session with broadcasts
	const setStart = (start: Date) => {
		if (!session) return
		UpdateSession(room, session.id, start, session.end).catch(console.error)
	}

	const setEnd = (end: Date) => {
		if (!session) return
		UpdateSession(room, session.id, session.start, end).catch(console.error)
	}

	const speakers = session && session.speakers ? session.speakers.join('、') : ''

	const content = session ? (
		<>
			<div className="flex w-full">
				<div className="grow">
					<h1 className="text-2xl">{session.title}</h1>
					<p>{speakers}</p>
					{session.broadcast.length > 0 && <p>轉播：{session.broadcast.join('、')}</p>}
				</div>
				<div>
					<a
						href={`/card?id=${session.id}`}
						className={twMerge(btn({ size: '2xl' }), 'grid place-items-center')}
						target="_blank"
					>
						<span>開啟字卡</span>
					</a>
				</div>
			</div>
			<p data-show={!detail && !isMd} className="data-[show=false]:hidden">
				{formatTime(session.start)} - {formatTime(session.end)}
			</p>
			<div data-show={detail || isMd} className="grid data-[show=false]:hidden md:grid-cols-2">
				<EditTime title="start" time={session.start} setTime={setStart} />
				<EditTime title="end" time={session.end} setTime={setEnd} />
			</div>
		</>
	) : (
		<div className="grid h-full w-full place-items-center">
			<p>Empty</p>
		</div>
	)
	return (
		<div
			data-current={isCurrent}
			className="min-h-[100px] rounded-lg bg-gray-100 p-2 shadow-lg data-[current=true]:bg-gray-200 md:min-h-[200px]"
			onClick={() => setDetail(!detail)}
		>
			{content}
		</div>
	)
}
