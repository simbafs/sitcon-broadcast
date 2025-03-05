'use client'

import { Countdown } from './countdown'
import { Now } from './now'
import { Session } from './session'
import { Special } from './special'
import { SSE } from './sse'

export default function Page() {
	return (
		<div className="p-20">
			<details>
				<summary>SSE</summary>
				<SSE />
			</details>
			<details>
				<summary>Now</summary>
				<Now />
			</details>
			<details>
				<summary>Session</summary>
				<Session />
			</details>
			<details>
				<summary>Countdown</summary>
				<Countdown />
			</details>
			<details>
				<summary>Special</summary>
				<Special />
			</details>
		</div>
	)
}
