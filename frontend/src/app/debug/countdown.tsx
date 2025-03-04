import { useState } from 'react'
import { input, useResult } from './useResult'
import * as api from '@/sdk/sdk'

function GetAllCountdown() {
	const { Result, set, clear } = useResult()

	return (
		<div>
			<button className={input} onClick={() => api.GetAllCountdown().then(set)}>
				Get
			</button>
			<button className={input} onClick={clear}>
				Clear
			</button>
			<Result />
		</div>
	)
}

function GetCountdownByName() {
	const [name, setName] = useState('')
	const { Result, set, clear } = useResult()

	return (
		<div>
			<input
				className={input}
				type="text"
				placeholder="countdown name"
				value={name}
				onChange={e => setName(e.target.value)}
			/>
			<button className={input} onClick={() => api.GetCountdownByName(name).then(set)}>
				Get
			</button>
			<button className={input} onClick={clear}>
				Clear
			</button>
			<Result />
		</div>
	)
}

function UpdateCountdown() {
	const [name, setName] = useState('')
	const [inittime, setInittime] = useState(0)
	const [time, setTime] = useState(0)
	const [state, setState] = useState<api.CountdownState>(0)
	const { Result, set, clear } = useResult()

	return (
		<div>
			<input
				className={input}
				type="text"
				placeholder="countdown name"
				value={name}
				onChange={e => setName(e.target.value)}
			/>
			<input
				className={input}
				type="number"
				placeholder="init time"
				value={inittime}
				onChange={e => setInittime(parseInt(e.target.value))}
			/>
			<input
				className={input}
				type="number"
				placeholder="time"
				value={time}
				onChange={e => setTime(parseInt(e.target.value))}
			/>
			<select
				className={input}
				value={state}
				onChange={e => setState(parseInt(e.target.value) as api.CountdownState)}
			>
				<option value="0">PAUSE</option>
				<option value="1">COUNTING</option>
			</select>
			<button
				className={input}
				onClick={() =>
					api
						.UpdateCountdown(name, {
							name,
							inittime,
							time,
							state,
						})
						.then(set)
				}
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

export function Countdown() {
	return (
		<div>
			<h1>GetAllCountdown</h1>
			<GetAllCountdown />
			<hr />
			<h1>GetCountdownByName</h1>
			<GetCountdownByName />
			<hr />
			<h1>UpdateCountdown</h1>
			<UpdateCountdown />
			<hr />
		</div>
	)
}
