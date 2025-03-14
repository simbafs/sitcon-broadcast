'use client'
import { constructNow, GetNow, type Now, parseNow, ResetNow, SetNow } from '@/sdk/now'
import { btn } from '@/style/btn'
import { useReducer } from 'react'
import { toast } from 'react-toastify'

export function Now() {
	const [now, setNow] = useReducer(
		(state: Now, action: number | Partial<Now>) => {
			if (typeof action === 'number') return parseNow(action)
			return { ...state, ...action }
		},
		parseNow(Date.now() / 1000),
	)

	return (
		<div>
			<div className="grid grid-cols-6">
				<input
					type="number"
					value={now.year}
					placeholder="Year"
					onChange={e => setNow({ year: parseInt(e.target.value) })}
					className={btn()}
				/>
				<input
					type="number"
					value={now.month}
					placeholder="Month"
					onChange={e => setNow({ month: parseInt(e.target.value) })}
					className={btn()}
				/>
				<input
					type="number"
					value={now.day}
					placeholder="Day"
					onChange={e => setNow({ day: parseInt(e.target.value) })}
					className={btn()}
				/>
				<input
					type="number"
					value={now.hours}
					placeholder="Hours"
					onChange={e => setNow({ hours: parseInt(e.target.value) })}
					className={btn()}
				/>
				<input
					type="number"
					value={now.minutes}
					placeholder="Minutes"
					onChange={e => setNow({ minutes: parseInt(e.target.value) })}
					className={btn()}
				/>
				<input
					type="number"
					value={now.seconds}
					placeholder="Seconds"
					onChange={e => setNow({ seconds: parseInt(e.target.value) })}
					className={btn()}
				/>
			</div>
			<button
				onClick={() =>
					GetNow()
						.then(setNow)
						.then(() => toast('取得時間成功'))
						.catch(e => toast(`取得時間失敗: ${e.message}`))
				}
				className={btn()}
			>
				取得時間
			</button>
			<button
				onClick={() =>
					SetNow(constructNow(now))
						.then(() => toast('設定時間成功'))
						.catch(e => toast(`設定時間失敗: ${e}`))
				}
				className={btn()}
			>
				設定時間
			</button>
			<button
				onClick={() =>
					ResetNow()
						.then(() => toast('重設時間成功'))
						.catch(e => toast(`重設時間失敗: ${e.message}`))
				}
				className={btn()}
			>
				重設時間
			</button>
		</div>
	)
}
