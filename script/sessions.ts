const SESSION_URL = 'https://sitcon.org/2025/sessions.json'
const DB_FILE = './sessions-new.db'
const SLIDO_CSV = './slido.csv'
import * as input from './input.ts'
import * as output from './output.ts'

import sqlite3 from 'better-sqlite3'

function mergeSameTitle(sessions: output.Session[], curr: output.Session) {
	const len = sessions.length
	if (len == 0) return [curr]
	if (sessions[len - 1].title == curr.title) {
		sessions[len - 1].end = curr.end
		return sessions
	} else {
		return [...sessions, curr]
	}
}

async function FetchAndParse() {
	const body: input.Root = await fetch(SESSION_URL).then(res => res.json())

	const rooms = body.rooms.map(item => item.zh.name)
	const speakers = Object.fromEntries(body.speakers.map(item => [item.id, item.zh.name]))
	const sessionType = Object.fromEntries(body.session_types.map(item => [item.id, item.zh.name]))

	const Sessions = Object.fromEntries(
		rooms.map(room => {
			return [
				room,
				body.sessions
					.map(s => {
						if (s.room != room || !s.broadcast?.includes(room)) return null
						return {
							id: s.id,
							idx: 0,
							title: s.zh.title,
							type: sessionType[s.type],
							room: room,
							broadcastTo: s.broadcast.filter(r => r != room) || [],
							broadcastFrom: s.broadcast.includes(room) && room != s.room ? room : '',
							start: new Date(s.start).getTime() / 1000,
							end: new Date(s.end).getTime() / 1000,

							speaker: s.speakers.map(id => speakers[id]).join('、'),

							qa: s.qa || '',
							slido_id: '', // TODO:
							slido_link: s.slide || '',
							slido_admin_link: '', // TODO:
							co_write: s.co_write || '',
						} as output.Session
					})
					.filter(i => !!i)
					.map(s => {
						if (['休息', '午餐', '點心'].includes(s.title)) {
							s.broadcastTo = []
							s.broadcastFrom = ''
							s.speaker = ''
							// TODO: use a unique id
						}
						return s
					})
					.reduce(mergeSameTitle, []),
			]
		}),
	)

	return Sessions
}

function SaveToDB(rooms: Record<string, output.Session[]>) {
	const db = new sqlite3(DB_FILE)

	const insert = db.prepare(`
        INSERT OR REPLACE INTO sessions (
			title, 
			id, 
			room, 
			broadcastTo, 
			broadcastFrom, 
			start, 
			end, 
			speaker,
			qa,
			slido_id, 
			slido_link,
			slido_admin_link,
			co_write
        ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`)

	const insertMany = db.transaction((rooms: Record<string, output.Session[]>) => {
		for (let room in rooms) {
			for (let s of rooms[room]) {
				insert.run(
					s.title,
					s.id,
					s.room,
					JSON.stringify(s.broadcastTo),
					s.broadcastFrom,
					s.start,
					s.end,
					s.speaker,
					s.qa,
					s.slido_id,
					s.slido_link,
					s.slido_admin_link,
					s.co_write,
				)
			}
		}
	})

	insertMany(rooms)

	db.close()
}

FetchAndParse().then(s => console.log(s['R0']))
