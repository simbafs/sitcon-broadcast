// TODO: merge adjacent sessions with the same title
// TODO: remove broadcast from 休息
const fs = require('fs')
const Database = require('better-sqlite3')
const csv = require('csv-parser')

const URL = 'https://sitcon.org/2025/sessions.json'
const DB_FILE = './sessions.db'
const SLIDO_CSV = './slido.csv'

let nextID = 0

function uniqueID() {
	return `休息-${nextID++}`
}

function convertTime(t) {
	return new Date(t).toISOString()
}

function mergeSessions(sessions) {
	if (sessions.length === 0) return []

	const merged = [sessions[0]]

	for (let i = 1; i < sessions.length; i++) {
		const prev = merged[merged.length - 1]
		const curr = sessions[i]

		if (prev.title === curr.title && prev.room === curr.room) {
			prev.end = curr.end
		} else {
			merged.push(curr)
		}
	}

	return merged
}

function fillGaps(sessions) {
	const sortedSessions = [...sessions].sort((a, b) => new Date(a.start) - new Date(b.start))
	const filled = []

	for (let i = 0; i < sortedSessions.length; i++) {
		if (i > 0 && new Date(sortedSessions[i].start) > new Date(sortedSessions[i - 1].end)) {
			filled.push({
				id: uniqueID(),
				title: '休息',
				type: 'Event',
				speakers: [],
				room: sortedSessions[i - 1].room,
				broadcast: [],
				start: sortedSessions[i - 1].end,
				end: sortedSessions[i].start,
				slido: '',
				slide: '',
				hackmd: '',
			})
		}
		filled.push(sortedSessions[i])
	}

	return filled
}

function saveSessionsToDB(sessions) {
	const db = new Database(DB_FILE)

	db.exec(`
		CREATE TABLE IF NOT EXISTS sessions (
			id TEXT PRIMARY KEY,
			title TEXT,
			type TEXT,
			speakers TEXT,
			room TEXT,
			broadcast TEXT,
			start TEXT,
			end TEXT,
			slido TEXT,
			slide TEXT,
			hackmd TEXT
		);
	`)

	db.exec('DELETE FROM sessions;')
	console.log('刪除後的紀錄數量:', db.prepare('SELECT COUNT(*) FROM sessions').get())

	const insertSession = db.prepare(`
		INSERT OR REPLACE INTO sessions (id, title, type, speakers, room, broadcast, start, end, slido, slide, hackmd)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);
	`)

	const insertMany = db.transaction(sessions => {
		const ids = new Set()
		const duplicatedIDs = []

		for (const session of sessions) {
			if (ids.has(session.id)) {
				duplicatedIDs.push(session.id)
				continue
			}
			ids.add(session.id)

			insertSession.run(
				session.id,
				session.title,
				session.type,
				JSON.stringify(session.speakers),
				session.room,
				JSON.stringify(session.broadcast),
				session.start,
				session.end,
				session.slido,
				session.slide,
				session.hackmd,
			)
		}

		if (duplicatedIDs.length > 0) {
			console.warn('跳過重複的 ID:', duplicatedIDs)
		}
	})

	insertMany(sessions)
	
	// for (let room of Object.keys(sessionsByRoom)) {
	// 	insertMany(sessionsByRoom[room])
	// }

	console.log('資料已成功寫入 SQLite')
	console.log('寫入的紀錄數量:', db.prepare('SELECT COUNT(*) FROM sessions').get())
	db.close()
}

function loadSlidoMappings(csvPath) {
	return new Promise((resolve, reject) => {
		const slidoMap = {}
		fs.createReadStream(csvPath)
			.pipe(csv({ headers: false }))
			.on('data', row => {
				const sessionID = row[0]?.trim()
				const slidoURL = row[1]?.trim()
				if (sessionID && slidoURL) {
					slidoMap[slidoURL] = sessionID
				}
			})
			.on('end', () => {
				console.log('成功載入 Slido 對應表:', Object.keys(slidoMap).length, '筆資料')
				resolve(slidoMap)
			})
			.on('error', reject)
	})
}

function loadAndSaveSpecial() {
	console.log('載入特殊議程')
	const cases = require('./special_cases.json')

	const db = new Database(DB_FILE)

	db.exec(`
		CREATE TABLE IF NOT EXISTS specials (
			id text NOT NULL, 
			data text NOT NULL DEFAULT (''), 
			PRIMARY KEY (id)
		) 
	`)

	db.exec('DELETE FROM specials;')
	console.log('刪除後的紀錄數量:', db.prepare('SELECT COUNT(*) FROM specials').get())

	const insertCase = db.prepare(`
		INSERT OR REPLACE INTO specials (id, data)
		VALUES (?, ?);
	`)

	for (const [id, data] of Object.entries(cases)) {
		insertCase.run(id, JSON.stringify(data))
	}

	console.log('資料已成功寫入 SQLite')
	db.close()
}

;(async () => {
	const slidoMap = await loadSlidoMappings(SLIDO_CSV)
	const data = await fetch(URL).then(res => res.json())

	console.log(data.sessions.length)

	const rooms = data.rooms.map(item => item.zh.name)
	const speakers = Object.fromEntries(data.speakers.map(item => [item.id, item.zh.name]))
	const sessionTypes = Object.fromEntries(data.session_types.map(item => [item.id, item.zh.name]))

	data.sessions.forEach(s => {
		if (['d01cb7', '710069', '12775d', '1e9ea4', '7540dc', 'e13fad', '06dc3c'].includes(s.id)) {
			console.log(s.id, s.broadcast)
		}
	})
	const sessions = data.sessions.map(s => ({
		id: s.id || '',
		title: s.zh.title || '',
		type: sessionTypes[s.type] || '',
		speakers: s.speakers.map(id => speakers[id]) || [],
		room: s.room || '',
		broadcast: s.broadcast || [],
		start: convertTime(s.start),
		end: convertTime(s.end),
		slido: slidoMap[s.qa?.trim()] || '', // 這裡使用 slido.csv 對應的 Slido 連結
		slide: s.slide || '',
		hackmd: s.co_write || '',
	}))

	console.log('====')
	sessions.forEach(s => {
		if (['d01cb7', '710069', '12775d', '1e9ea4', '7540dc', 'e13fad', '06dc3c'].includes(s.id)) {
			console.log(s.id, s.broadcast)
		}
	})

	// let filledSessions = {}
	// for (let room of rooms) {
	// 	let roomSessions = sessions.filter(s => s.room === room || s.broadcast.includes(room))
	// 	// TDOO: 2025 的議程表是緊湊的，不需要合併相鄰的議程和填補休息
	// 	// roomSessions = mergeSessions(roomSessions)
	// 	// filledSessions[room] = fillGaps(roomSessions)
	// 	filledSessions[room] = roomSessions
	// }

	// for (let room of rooms) {
	// 	filledSessions[room] = filledSessions[room].map(session => {
	// 		if (session.title === '休息') {
	// 			session.broadcast = []
	// 		}
	// 		return session
	// 	})
	// }

	saveSessionsToDB(sessions)
	loadAndSaveSpecial()
})()
