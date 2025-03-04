import { useState } from 'react'
import * as api from '@/sdk/sdk'
import { useResult, input } from './useResult'

function GetNow() {
	const { Result, set, clear } = useResult()

	return (
		<div>
			<button className={input} onClick={() => api.GetNow().then(set, set)}>
				Get
			</button>
			<button className={input} onClick={clear}>
				Clear
			</button>
			<Result />
		</div>
	)
}

function SetNow() {
	const [now, setNow] = useState('')
	const { Result, set, clear } = useResult()

	return (
		<div>
			<input className={input} type="datetime-local" value={now} onChange={e => setNow(e.target.value)} />
			<button className={input} onClick={() => api.SetNow(new Date(now)).then(set, set)}>
				Set
			</button>
			<button className={input} onClick={clear}>
				Clear
			</button>
			<Result />
		</div>
	)
}

function ResetNow() {
	const { Result, set, clear } = useResult()

	return (
		<div>
			<button className={input} onClick={() => api.ResetNow().then(set, set)}>
				Reset
			</button>
			<button className={input} onClick={clear}>
				Clear
			</button>
			<Result />
		</div>
	)
}

export function Now() {
	return (
		<div>
			<h1>GetNow</h1>
			<GetNow />
			<hr />
			<h1>SetNow</h1>
			<SetNow />
			<hr />
			<h1>ResetNow</h1>
			<ResetNow />
		</div>
	)
}
