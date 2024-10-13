import { Session } from '@/types/card'

export function getNow(sessions: Session[]) {
	const now = new Date()
	const nowTime = now.getHours() * 60 + now.getMinutes()

	for (let i = 0; i < sessions.length; i++) {
		if (nowTime < sessions[i].start) {
			// console.log({ i, now, start: sessions[i-1].start })
			return i - 1
		}
	}

	// console.log('not found')
	return sessions.length - 1
}
