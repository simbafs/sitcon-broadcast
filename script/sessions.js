// TODO: merge adjacent sessions with the same title
// TODO: remove broadcast from 休息
const URL = 'https://sitcon.org/2025/sessions.json'

let nextID = 0

function uniqueID() {
	return `休息-${nextID++}`
}

function convertTime(t) {
	const time = new Date(t)
	return time.getHours() * 60 + time.getMinutes()
}

/**
 * @typedef {Object} Session
 * @property {string} id
 * @property {string} title
 * @property {string} type
 * @property {string[]} speakers
 * @property {string} room
 * @property {string[]} broadcast
 * @property {number} start
 * @property {number} end
 */

/**
 * 合併相鄰且標題相同的會議
 * @param {Session[]} sessions
 * @returns {Session[]}
 */
function mergeSessions(sessions) {
	if (sessions.length === 0) return []

	const merged = [sessions[0]]

	for (let i = 1; i < sessions.length; i++) {
		const prev = merged[merged.length - 1]
		const curr = sessions[i]

		// 如果標題相同，則合併時間區間
		if (prev.title === curr.title && prev.room === curr.room) {
			prev.end = Math.max(prev.end, curr.end)
		} else {
			merged.push(curr)
		}
	}

	return merged
}

/**
 * 填補空隙
 * @param {Session[]} sessions
 * @returns {Session[]}
 */
function fillGaps(sessions) {
	const sortedSessions = [...sessions].sort((a, b) => a.start - b.start)
	/** @type{Session[]} */
	const filled = []

	for (let i = 0; i < sortedSessions.length; i++) {
		if (i > 0 && sortedSessions[i].start > sortedSessions[i - 1].end) {
			// 插入「休息」時間
			filled.push({
				id: uniqueID(),
				title: '休息',
				type: 'Event',
				speakers: [],
				room: sortedSessions[i - 1].room,
				broadcast: [], // 確保不含 broadcast
				start: sortedSessions[i - 1].end,
				end: sortedSessions[i].start,
			})
		}
		filled.push(sortedSessions[i])
	}

	return filled
}

;(async () => {
	const data = await fetch(URL).then(res => res.json())

	const rooms = data.rooms.map(item => item.zh.name)
	const speakers = Object.fromEntries(data.speakers.map(item => [item.id, item.zh.name]))
	const sessionTypes = Object.fromEntries(data.session_types.map(item => [item.id, item.zh.name]))

	/** @type {Session[]} */
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

	// 根據不同場地填補空隙
	const filledSessions = {}

	for (let room of rooms) {
		let roomSessions = sessions.filter(s => s.room === room || s.broadcast.includes(room))

		// 合併相鄰相同標題的會議
		roomSessions = mergeSessions(roomSessions)

		// 填補空隙
		filledSessions[room] = fillGaps(roomSessions)
	}

	// 確保所有「休息」的 broadcast 欄位都清空
	for (let room of rooms) {
		filledSessions[room] = filledSessions[room].map(session => {
			if (session.title === '休息') {
				session.broadcast = []
			}
			return session
		})
	}

	// 建立 ID 對應表
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
