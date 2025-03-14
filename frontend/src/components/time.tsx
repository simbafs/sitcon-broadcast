import { parseTime } from '@/sdk/time'

function pad2(n: number) {
	return n.toString().padStart(2, '0')
}

export function Time({ time }: { time: number }) {
	const t = parseTime(time)
	return (
		<p>
			{t.year}-{pad2(t.month)}-{pad2(t.day)} {pad2(t.hours)}:{pad2(t.minutes)}:{pad2(t.seconds)}
		</p>
	)
}
