import { Session } from '@/types/card'
import { formatTime } from '@/utils/formatTime'
import { getCurrent } from './getCurrent'
import { useState } from 'react'
import { btn } from '@/varients/btn'
import { twMerge } from 'tailwind-merge'
import { useTime } from '@/components/useTime'
import { mutate } from 'swr'

export function Room({ sessions }: { sessions: Session[] }) {
	const now = getCurrent(sessions)
	const [current, setCurrent] = useState(() => getCurrent(sessions))
	return (
		<>
			<div className="m-4 flex w-full justify-around">
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

			<ul className="m-2 flex w-full flex-col justify-stretch gap-4">
				{/* previous */}
				<EditSession idx={current - 1} sessions={sessions} isCurrent={now == current - 1} key={current - 1} />
				{/* current */}
				<EditSession idx={current} sessions={sessions} isCurrent={now == current} key={current} />
				{/* next */}
				<button className={twMerge(btn())}>新增議程(Unimplemented)</button>
				<EditSession idx={current + 1} sessions={sessions} isCurrent={now == current + 1} key={current + 1} />
			</ul>
		</>
	)
}

function EditSession({ sessions, idx, isCurrent }: { sessions: Session[]; idx: number; isCurrent: boolean }) {
	const session: Session | undefined = sessions[idx]
	const [detail, setDetail] = useState(false)

	// TODO: case of the session with broadcasts
	const setStart = async (start: number) => {
		if (!session) return
		if (session.end <= start) return // if start after end, do nothing
		if (sessions[idx - 1] && sessions[idx - 1].start >= start) return // if start before previous start, do nothing

		const isShorten = start > session.start

		const current = { ...session, start }
		const pre = { ...sessions[idx - 1], end: start }

		// if (isShorten)
		// 	console.log(
		// 		`${idx}(${session.start} -> ${current.start}), ${idx - 1}(${sessions[idx - 1].end} -> ${pre.end})`,
		// 	)
		// else
		// 	console.log(
		// 		`${idx - 1}(${sessions[idx - 1].end} -> ${pre.end}), ${idx}(${session.start} -> ${current.start})`,
		// 	)

		return fetch(`/api/card/${session.room}/${isShorten ? idx : idx - 1}`, {
			method: 'POST',
			body: JSON.stringify(isShorten ? current : pre),
		})
			.then(() =>
				fetch(`/api/card/${session.room}/${isShorten ? idx - 1 : idx}`, {
					method: 'POST',
					body: JSON.stringify(isShorten ? pre : current),
				}),
			)
			// .then(() => console.log('done'))
			.then(() => mutate(`/api/session`))
			.catch(console.error)
	}

	const setEnd = (end: number) => {
		if (!session) return
		if (session.start >= end) return // if end before start, do nothing
		if (sessions[idx + 1] && sessions[idx + 1].end <= end) return // if end after next end, do nothing

		const isShorten = end < session.end

		const current = { ...session, end }
		const next = { ...sessions[idx + 1], start: end }

		// if (isShorten)
		// 	console.log(
		// 		`${idx}(${session.end} -> ${current.end}), ${idx + 1}(${sessions[idx + 1].start} -> ${next.start})`,
		// 	)
		// else
		// 	console.log(
		// 		`${idx + 1}(${sessions[idx + 1].start} -> ${next.start}), ${idx}(${session.end} -> ${current.end})`,
		// 	)

		return fetch(`/api/card/${session.room}/${isShorten ? idx : idx + 1}`, {
			method: 'POST',
			body: JSON.stringify(isShorten ? current : next),
		})
			.then(() =>
				fetch(`/api/card/${session.room}/${isShorten ? idx + 1 : idx}`, {
					method: 'POST',
					body: JSON.stringify(isShorten ? next : current),
				}),
			)
			// .then(() => console.log('done'))
			.then(() => mutate(`/api/session`))
			.catch(console.error)
	}

	const Start = useTime(`Start: ${formatTime(session?.start || 0)}`, session?.start || 0, setStart)
	const End = useTime(`End: ${formatTime(session?.end || 0)}`, session?.end || 0, setEnd)

	const content = session ? (
		<>
			<h1 className="text-2xl">{session.title}</h1>
			<p>{session.speakers.join('、')}</p>
			<p data-detail={detail} className="data-[detail=true]:hidden">
				{formatTime(session.start)} - {formatTime(session.end)}
			</p>
			<div data-detail={detail} className="data-[detail=false]:hidden">
				{Start}
				{End}
			</div>
		</>
	) : (
		<p>Empty</p>
	)
	return (
		<div
			data-current={isCurrent}
			className="rounded-lg bg-gray-100 p-2 shadow-lg data-[current=true]:bg-gray-200 "
			onClick={() => setDetail(!detail)}
		>
			{content}
		</div>
	)
}
