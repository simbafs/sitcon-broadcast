import { Session } from '@/sdk/sdk'

export function getCurrent(sessions: Session[]) {
	if (sessions.length == 0) return 0

	const now = new Date()

	for (let i = 0; i < sessions.length; i++) {
		if (sessions[i].end > now) {
			return i
		}
	}
	return sessions.length - 1
}
