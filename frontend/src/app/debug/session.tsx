import { useEffect, useState } from 'react'
import { input, useResult } from './useResult'
import * as api from '@/sdk/sdk'

function GetAllSessions() {
	const [room, setRoom] = useState('R0')
	const { Result, set, clear } = useResult()

	return (
		<div>
			<select className={input} value={room} onChange={e => setRoom(e.target.value)}>
				<option value="">All</option>
				<option value="R0">R0</option>
				<option value="R1">R1</option>
				<option value="R2">R2</option>
				<option value="R3">R3</option>
				<option value="S">S</option>
			</select>
			<button
				className={input}
				onClick={() =>
					api.GetAllSessions().then(ss => {
						if (room) set(ss.filter(s => s.room == room || s.broadcast.includes(room)))
						else set(ss)
					}, set)
				}
			>
				Get
			</button>
			<button className={input} onClick={clear}>
				Clear
			</button>
			<Result />
		</div>
	)
}

function GetSessionByID() {
	const [id, setID] = useState('')
	const { Result, set, clear } = useResult()

	return (
		<div>
			<input
				className={input}
				type="text"
				placeholder="session id"
				value={id}
				onChange={e => setID(e.target.value)}
			/>
			<button className={input} onClick={() => api.GetSessionByID(id).then(set, set)}>
				Get
			</button>
			<button className={input} onClick={clear}>
				Clear
			</button>
			<Result />
		</div>
	)
}

function GetCurrentSession() {
	const [room, setRoom] = useState('R0')
	const { Result, set, clear } = useResult()

	useEffect(() => console.log(room), [room])

	return (
		<div>
			<select className={input} value={room} onChange={e => setRoom(e.target.value)}>
				<option value="R0">R0</option>
				<option value="R1">R1</option>
				<option value="R2">R2</option>
				<option value="R3">R3</option>
				<option value="S">S</option>
			</select>
			<button className={input} onClick={() => api.GetCurrentSession(room).then(set, set)}>
				Get
			</button>
			<button className={input} onClick={clear}>
				Clear
			</button>
			<Result />
		</div>
	)
}

function UpdateSession() {
	const [room, setRoom] = useState('R0')
	const [id, setID] = useState('')
	const [start, setStart] = useState('')
	const [end, setEnd] = useState('')
	const { Result, set, clear } = useResult()

	return (
		<div>
			<select className={input} value={room} onChange={e => setRoom(e.target.value)}>
				<option value="R0">R0</option>
				<option value="R1">R1</option>
				<option value="R2">R2</option>
				<option value="R3">R3</option>
				<option value="S">S</option>
			</select>
			<input className={input} type="text" placeholder="id" value={id} onChange={e => setID(e.target.value)} />
			<input className={input} type="datetime-local" value={start} onChange={e => setStart(e.target.value)} />
			<input className={input} type="datetime-local" value={end} onChange={e => setEnd(e.target.value)} />
			<button
				className={input}
				onClick={() => api.UpdateSession(room, id, new Date(start), new Date(end)).then(set, set)}
			>
				Update
			</button>
			<button className={input} onClick={clear}>
				Clear
			</button>

			<Result />
		</div>
	)
}

export function Session() {
	return (
		<div>
			<h1>GetAllSessions</h1>
			<GetAllSessions />
			<hr />
			<h1>GetSessionByID</h1>
			<GetSessionByID />
			<hr />
			<h1>GetCurrentSession</h1>
			<GetCurrentSession />
			<hr />
			<h1>UpdateSession</h1>
			<UpdateSession />
			<hr />
		</div>
	)
}
