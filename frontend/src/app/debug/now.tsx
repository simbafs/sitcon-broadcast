import { useEffect, useState } from 'react'
import * as api from '@/sdk/sdk'
import { useResult, input } from './useResult'
import { formatTime } from '@/utils/formatTime'

function GetNow() {
	const { Result, set, clear, result } = useResult()

	return (
		<div>
			<button className={input} onClick={() => api.GetNow().then(set, set)}>
				Get
			</button>
			<button className={input} onClick={clear}>
				Clear
			</button>
			<p>{result && formatTime(new Date(JSON.parse(result)))}</p>
			<Result />
		</div>
	)
}

function SetNow() {
	const [now, setNow] = useState('')
	const { Result, set, clear } = useResult()

	useEffect(() => console.log(now), [now])

	const loadNow = () => {
		api.GetNow().then(d =>
			setNow(
				`${d.getFullYear()}-${(d.getMonth() + 1).toString().padStart(2, '0')}-${d.getDate().toString().padStart(2, '0')}T${d.getHours().toString().padStart(2, '0')}:${d.getMinutes().toString().padStart(2, '0')}`,
			),
		)
	}

	return (
		<div>
			<button className={input} onClick={loadNow}>
				Load Now
			</button>
			<br />
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
