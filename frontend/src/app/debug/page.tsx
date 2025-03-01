'use client'
import { useAllSSE } from '@/hooks/useSSE'
import { btn } from '@/varients/btn'
import { useRef, useState } from 'react'

function SSE() {
	const msg = useAllSSE()

	if (!msg) return <div>loading...</div>

	return (
		<>
			<h1 className="text-3xl">Server Sent Events</h1>
			<div className="m-4 flex gap-4">
				{Object.entries(msg).map(([name, msgs]) => (
					<div key={name} className="rounded border border-black p-4">
						<h2>{name}</h2>
						<ul>
							{msgs.map((msg, i) => (
								<li key={i}>{JSON.stringify(msg, null, 1)}</li>
							))}
						</ul>
					</div>
				))}
			</div>
		</>
	)
}

function Now() {
	const [res, setRes] = useState('')
	const hour = useRef<HTMLInputElement>(null)
	const minute = useRef<HTMLInputElement>(null)

	const send = () => {
		setRes('loading...')
		const time = (hour.current?.valueAsNumber || 0) * 60 + (minute.current?.valueAsNumber || 0)
		console.log(time)
		fetch('/api/now', {
			method: 'POST',
			body: time.toString(),
		})
			.then(res => res.json())
			.then(data => {
				if (data.error) {
					setRes(data.error)
				} else {
					setRes(data.data)
				}
			})
			.catch(e => {
				setRes(e.toString())
			})
	}

	const reset = () => {
		setRes('loading...')
		fetch('/api/now', {
			method: 'DELETE',
		})
			.then(res => res.json())
			.then(data => {
				if (data.error) {
					setRes(data.error)
				} else {
					setRes(data.data)
				}
			})
			.catch(e => {
				setRes(e.toString())
			})
	}

	const getNow = () => {
		setRes('loading...')
		fetch('/api/now')
			.then(res => res.json())
			.then(data => {
				if (data.error) {
					setRes(data.error)
				} else {
					if (!hour.current || !minute.current) return
					hour.current.valueAsNumber = Math.floor(data.now / 60)
					minute.current.valueAsNumber = data.now % 60
				}
				setRes('')
			})
			.catch(e => {
				setRes(e.toString())
			})
	}

	return (
		<>
			<h1>Now</h1>
			<div className="flex flex-col gap-4">
				<div className="flex gap-4">
					<input
						className="rounded border-2 border-black"
						type="number"
						min="0"
						max="24"
						step="1"
						defaultValue="9"
						ref={hour}
					/>
					<input
						className="rounded border-2 border-black"
						type="number"
						min="0"
						max="59"
						step="1"
						defaultValue="0"
						ref={minute}
					/>
				</div>
				<div className="flex gap-4">
					<button className={btn()} onClick={send}>
						Set
					</button>
					<button className={btn()} onClick={reset}>
						Reset
					</button>
					<button className={btn()} onClick={getNow}>
						Get
					</button>
				</div>
				<p>{res}</p>
			</div>
		</>
	)
}

export default function Page() {
	return (
		<div className="px-20">
			<SSE />
			<Now />
		</div>
	)
}
