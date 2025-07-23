import data from '../data/2025.json' with { type: 'json' }

const speakers = Object.fromEntries(data.speakers.map(s => [s.id, s.zh.name]))

type Session = (typeof data.sessions)[number]
type SessionInRoom = {
	title: string
	start: Date
	end: Date
	speakers: string[]
}

function getSessionInRoom(room: string): Session[] {
	return data.sessions.filter(s => s.room == room || s.broadcast?.includes(room))
}

function sortByTime(sessions: Session[]): Session[] {
	return sessions.sort((a, b) => {
		const aTime = new Date(a.start).getTime()
		const bTime = new Date(b.start).getTime()
		return aTime - bTime
	})
}

function toSessionInRoom(sessions: Session[]): SessionInRoom[] {
	return sessions.map(session => ({
		title: session.zh.title,
		start: new Date(session.start),
		end: new Date(session.end),
		speakers: (session.speakers as string[]).map(s => speakers[s] || s),
	}))
}

function fillGap(sessions: SessionInRoom[]): SessionInRoom[] {
	const result: SessionInRoom[] = []
	let lastEnd: Date | null = null

	for (const session of sessions) {
		if (lastEnd && session.start > lastEnd) {
			result.push({
				title: '空檔',
				start: lastEnd,
				end: session.start,
				speakers: [],
			})
		}
		result.push(session)
		lastEnd = session.end
	}

	return result
}

function mergeSameTitle(sessions: SessionInRoom[]): SessionInRoom[] {
	const result: SessionInRoom[] = []

	for (const session of sessions) {
		if (result.length === 0 || result[result.length - 1].title != session.title) {
			result.push(session)
			continue
		}

		const lastSession = result[result.length - 1]
		lastSession.end = session.end
		const speakers = new Set(lastSession.speakers.concat(session.speakers))
		lastSession.speakers = Array.from(speakers)
	}

	return result.map(s => ({
		...s,
		speakers: Array.from(new Set(s.speakers)),
	}))
}

function pipe<T>(value: T) {
	return {
		do<U>(fn: (input: T) => U) {
			return pipe(fn(value))
		},
		get() {
			return value
		},
	}
}

pipe(
	data.rooms.map(r => [
		r.id,
		pipe(getSessionInRoom(r.id)).do(sortByTime).do(toSessionInRoom).do(fillGap).do(mergeSameTitle).get(),
	]),
)
	.do(Object.fromEntries)
	.do(JSON.stringify)
	.do(console.log)
