const SESSION_URL = 'https://sitcon.org/2025/sessions.json'
const DB_FILE = './sitcon.db'
import * as input from './input.ts'
import * as output from './output.ts'

import sqlite3 from 'better-sqlite3'

function parseInput(speakers: Record<string, string>) {
	return function (s: input.Session) {
		const b =
			s.broadcast?.map<output.Session>(r => ({
				start: new Date(s.start).getTime() / 1000,
				end: new Date(s.end).getTime() / 1000,

				id: s.id,
				room: r,

				next: '',
				title: s.zh.title,
				speaker: s.speakers.map(id => speakers[id]).join(', '),
			})) || []
		const r = [
			{
				start: new Date(s.start).getTime() / 1000,
				end: new Date(s.end).getTime() / 1000,

				id: s.id,
				room: s.room,

				next: '',
				title: s.zh.title,
				speaker: s.speakers.map(id => speakers[id]).join(', '),
			} as output.Session,
		]

		return b.length == 0 ? r : b
	}
}

function mergeSameTitle(sessions: output.Session[], curr: output.Session) {
	const len = sessions.length
	if (len == 0) return [curr]
	if (sessions[len - 1].room == curr.room && sessions[len - 1].title == curr.title) {
		sessions[len - 1].end = curr.end
		return sessions
	} else {
		return [...sessions, curr]
	}
}

function removeSpeakerFromRest(s: output.Session) {
	if (['休息', '午餐', '點心'].includes(s.title)) {
		s.speaker = ''
	}
	return s
}

async function FetchAndParse() {
	const body: input.Root = await fetch(SESSION_URL).then(res => res.json())

	const rooms = body.rooms.map(item => item.zh.name)
	const speakers = Object.fromEntries(body.speakers.map(item => [item.id, item.zh.name]))
	// const sessionType = Object.fromEntries(body.session_types.map(item => [item.id, item.zh.name]))

	const sessions = body.sessions
		.flatMap(parseInput(speakers))
		.filter(i => !!i)
		.map(removeSpeakerFromRest)
		.reduce(mergeSameTitle, [])

	// set next by room
	const roomSessions = rooms.map(room => sessions.filter(s => s.room == room))
	for (let room of roomSessions) {
		room.sort((a, b) => a.start - b.start)
		for (let i = 0; i < room.length - 1; i++) {
			room[i].next = room[i + 1].id
		}
	}

	return roomSessions.flat()
}

function SaveToDB(sessions: output.Session[]) {
	const db = new sqlite3(DB_FILE)

	db.exec('DELETE FROM sessions;')

	const insert = db.prepare(`
        INSERT OR REPLACE INTO sessions (
			start,
			end,
			session_id,
			room,
			next,
			title,
			speaker
        ) VALUES (?, ?, ?, ?, ?, ?, ?)
	`)

	const insertMany = db.transaction((ss: output.Session[]) => {
		for (let s of ss) {
			insert.run(s.start, s.end, s.id, s.room, s.next, s.title, s.speaker)
		}
	})

	insertMany(sessions)

	db.close()
}

FetchAndParse().then(SaveToDB)
