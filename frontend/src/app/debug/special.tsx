import { useState } from 'react'
import * as api from '@/sdk/sdk'
import { useResult, input } from './useResult'

function GetAllSpecial() {
	const { Result, set, clear } = useResult()

	return (
		<div>
			<button className={input} onClick={() => api.GetAllSpecial().then(set, set)}>
				Get
			</button>
			<button className={input} onClick={clear}>
				Clear
			</button>
			<Result />
		</div>
	)
}

function GetSpecialByID() {
	const [id, setID] = useState('')
	const { Result, set, clear } = useResult()

	return (
		<div>
			<input className={input} value={id} onChange={e => setID(e.target.value)} />
			<button className={input} onClick={() => api.GetSpecialByID(id).then(set, set)}>
				Get
			</button>
			<button className={input} onClick={clear}>
				Clear
			</button>
			<Result />
		</div>
	)
}

function UpdateSpecial() {
	const [id, setID] = useState('')
	const [special, setSpecial] = useState('')
	const { Result, set, clear } = useResult()

	return (
		<div>
			<input className={input} value={id} onChange={e => setID(e.target.value)} />
			<textarea className={input} value={special} onChange={e => setSpecial(e.target.value)} />
			<button className={input} onClick={() => api.UpdateSpecial(id, special).then(set, set)}>
				Update
			</button>
			<button className={input} onClick={clear}>
				Clear
			</button>
			<Result />
		</div>
	)
}

export function Special() {
	return <div>
	    <h1>GetAllSpecial</h1>
	    <GetAllSpecial />
	    <hr />
	    <h1>GetSpecialByID</h1>
	    <GetSpecialByID />
	    <hr />
	    <h1>UpdateSpecial</h1>
	    <UpdateSpecial />
	    <hr />
	</div>
}
