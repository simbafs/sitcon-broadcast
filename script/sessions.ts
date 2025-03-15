const SESSION_URL = 'https://sitcon.org/2025/sessions.json'
const DB_FILE = './sitcon.db'
import * as input from './input.ts'
import * as output from './output.ts'

import sqlite3 from 'better-sqlite3'

function parseInput(speakers: Record<string, string>) {
	return function (s: input.Session) {
		const extraData = {
			speaker: s.speakers || [],
			qa: s.qa || '',
			slide: s.slide || '',
			co_write: s.co_write || '',
			record: s.record || '',
			live: s.live || '',
			tags: s.tags || [],
			url: s.uri || '',
			description: s.zh.description,
		}

		const b =
			s.broadcast?.map<output.Session>(r => ({
				idx: 0,
				start: new Date(s.start).getTime() / 1000,
				end: new Date(s.end).getTime() / 1000,
				id: s.id,
				room: r,
				next: '',
				title: s.zh.title,
				data: JSON.stringify(extraData),
			})) || []

		const r = [
			{
				idx: 0,
				start: new Date(s.start).getTime() / 1000,
				end: new Date(s.end).getTime() / 1000,
				id: s.id,
				room: s.room,
				next: '',
				title: s.zh.title,
				data: JSON.stringify(extraData),
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
		s.data = JSON.stringify({ ...JSON.parse(s.data), speaker: '' })
	}
	return s
}

function setNext(s: output.Session, idx: number, arr: output.Session[]) {
	if (idx < arr.length - 1) {
		s.next = arr[idx + 1].id
	}
	return s
}

function setIdx(s: output.Session, idx: number) {
	s.idx = idx
	return s
}

async function FetchAndParse() {
	const body: input.Root = await fetch(SESSION_URL).then(res => res.json())

	const rooms = body.rooms.map(item => item.zh.name)
	const speakers = Object.fromEntries(body.speakers.map(item => [item.id, item.zh.name]))

	const sessions = body.sessions
		.flatMap(parseInput(speakers))
		.filter(i => !!i)
		.map(removeSpeakerFromRest)
		.reduce(mergeSameTitle, [])

	const roomSessions = rooms.map(room =>
		sessions
			.filter(s => s.room == room)
			.toSorted((a, b) => a.start - b.start)
			.map(setNext)
			.map(setIdx),
	)

	return roomSessions.flat()
}

function SaveToDB(sessions: output.Session[]) {
	const db = new sqlite3(DB_FILE)

	db.exec('DELETE FROM sessions;')

	const insert = db.prepare(`
        INSERT OR REPLACE INTO sessions (
			idx,
            start,
            end,
            session_id,
            room,
            next,
            title,
            data
        ) VALUES (@idx, @start, @end, @id, @room, @next, @title, @data);
    `)

	const insertMany = db.transaction((ss: output.Session[]) => {
		for (let s of ss) {
			insert.run(s)
		}
	})

	insertMany(sessions)

	db.close()
}

FetchAndParse().then(SaveToDB)
