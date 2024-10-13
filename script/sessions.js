// TODO: merge adjacent sessions with the same title
// remove broadcast from 休息
const URL = 'https://sitcon.org/2024/sessions.json'

let nextID = 0

function uniqueID() {
	return `id-${nextID++}`
}

function convertTime(t) {
	const time = new Date(t)
	return time.getHours() * 60 + time.getMinutes()
}

/**
 * @param {Array<{id: string, title: string, type: string, speakers: string[], room: string, broadcast: string[], start: number, end: number}>} sessions
 */
function fillGaps(sessions) {
	sessions.sort((a, b) => a.start - b.start)
	const n = sessions.length
	for (let i = 1; i < n; i++) {
		if (sessions[i].start > sessions[i - 1].end) {
			sessions.push({
				id: uniqueID(),
				title: '休息',
				type: 'Event',
				speakers: [],
				room: sessions[i - 1].room,
				broadcast: [],
				start: sessions[i - 1].end,
				end: sessions[i].start,
			})
		}
	}
	sessions.sort((a, b) => a.start - b.start)

	return sessions
}

;(async () => {
	const data = await fetch(URL).then(res => res.json())

	const rooms = data.rooms.map(item => item.zh.name)
	const speakers = Object.fromEntries(data.speakers.map(item => [item.id, item.zh.name]))
	const sessionTypes = Object.fromEntries(data.session_types.map(item => [item.id, item.zh.name]))

	const sessions = data.sessions.map(s => ({
		id: s.id,
		title: s.zh.title,
		type: sessionTypes[s.type],
		speakers: s.speakers.map(id => speakers[id]),
		room: s.room,
		broadcast: s.broadcast || [],
		start: convertTime(s.start),
		end: convertTime(s.end),
	}))

	// fill all gaps with 「休息時間」

	const filledSessions = {}

	for (let room of rooms) {
		filledSessions[room] = fillGaps(sessions.filter(s => s.room === room || s.broadcast.includes(room)))
	}

	const idMap = {}

	for (let room of rooms) {
		for (let [index, s] of filledSessions[room].entries()) {
			idMap[s.id] = {
				room,
				index,
			}
		}
	}

	console.log(
		JSON.stringify({
			sessions: filledSessions,
			idMap,
			nextID,
		}),
	)
})()
