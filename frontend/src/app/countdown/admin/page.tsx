'use client'
import useRoom, { COUNTING, PAUSE, RoomData, type Room } from '@/hooks/useRoom'
import { btn } from '@/varients/btn'
import Link from 'next/link'
import { setEditor, useEditTime } from '@/components/useEditTime'
import useSWR from 'swr'
import { formatTime } from '@/utils/formatTime'
import { useInterval } from 'usehooks-ts'

function Row({ room, setTimeEditor }: { room: Room; setTimeEditor: setEditor }) {
	const setTime = () => {
		setTimeEditor(formatTime(room.inittime)).then(time => {
			const [m, s] = time.split(':').map(Number)
			const t = m * 60 + s
			room.setTime(t)
		})
	}
	return (
		<div className="grid gap-4 grid-cols-1 lg:grid-cols-[2fr_4fr]">
			<div className="grid grid-cols-2 gap-6">
				<h2 className="text-center text-3xl">{room.name}</h2>
				<p>{formatTime(room.time)}</p>
			</div>
			<div className="grid grid-cols-5 gap-6">
				<button
					className={btn({ color: room.state === PAUSE ? 'green' : 'normal' })}
					onClick={room.start}
					disabled={room.state === COUNTING}
				>
					開始
				</button>
				<button
					className={btn({ color: room.state === COUNTING ? 'red' : 'normal' })}
					onClick={room.pause}
					disabled={room.state === PAUSE}
				>
					暫停
				</button>
				<button className={btn({ color: 'yellow' })} onClick={room.reset}>
					重設
				</button>
				<button className={btn({ color: 'yellow' })} onClick={setTime}>
					設定時間
				</button>
				<Link className={btn()} href={`/countdown?id=${room.id}`} target="_blank">
					開啟頁面
				</Link>
			</div>
		</div>
	)
}

// a component to show server time with API /api/now
function ServerTime() {
	const { data, error } = useSWR<{ now: number }>('/api/now/', url => fetch(url).then(res => res.json()), {
		refreshInterval: 500,
	})

	if (error) {
		return <h1>Error: {error.message}</h1>
	}

	if (!data) {
		return <h1>Loading...</h1>
	}

	return <h1 className="mt-10 text-2xl">現在時間: {formatTime(data.now)}</h1>
}

function Rooms({ setTimeEditor }: { setTimeEditor: setEditor }) {
	const rooms = [useRoom(0), useRoom(1), useRoom(2), useRoom(3), useRoom(4)]

	useInterval(() => {
		fetch('/api/room/', {
			cache: 'no-cache',
		})
			.then(res => res.json())
			.then((data: { rooms: RoomData[] }) => {
				data.rooms.forEach((room, i) => {
					rooms[i].updateRoom({
						time: room.time,
						state: room.state,
					})
				})
			})
	}, 500)

	return (
		<div className="w-full grid gap-[50px]">
			{rooms.map((r, i) => (
				<Row key={i} room={r} setTimeEditor={setTimeEditor} />
			))}
		</div>
	)
}

export default function Page() {
	const [TimeEditor, setTimeEditor] = useEditTime()

	return (
		<>
			<TimeEditor />
			<div className="min-h-screen w-screen py-[100px] px-[50px] lg:px-[100px] flex flex-col justify-center items-center">
				<Rooms setTimeEditor={setTimeEditor} />
				<ServerTime />
			</div>
		</>
	)
}
