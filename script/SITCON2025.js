function parseInput(speakers) {
	return function (s) {
		const extraData = {
			speaker: speakers[s.speakers] || [],
			qa: s.qa || '',
			slide: s.slide || '',
			co_write: s.co_write || '',
			record: s.record || '',
			live: s.live || '',
			tags: s.tags || [],
			url: s.uri || '',
			description: s.zh.description,
		}

		if (s.broadcast) {
			return (
				s.broadcast?.map(r => ({
					idx: 0,
					start: new Date(s.start).getTime() / 1000,
					end: new Date(s.end).getTime() / 1000,
					session_id: s.id,
					room: r,
					next: '',
					title: s.zh.title,
					data: extraData,
				})) || []
			)
		} else {
			return [
				{
					idx: 0,
					start: new Date(s.start).getTime() / 1000,
					end: new Date(s.end).getTime() / 1000,
					session_id: s.id,
					room: s.room,
					next: '',
					title: s.zh.title,
					data: extraData,
				},
			]
		}
	}
}

function mergeSameTitle(sessions, curr) {
	const len = sessions.length
	if (len == 0) return [curr]
	if (sessions[len - 1].room == curr.room && sessions[len - 1].title == curr.title) {
		sessions[len - 1].end = curr.end
		return sessions
	} else {
		return [...sessions, curr]
	}
}

function removeSpeakerFromRest(s) {
	if (['休息', '午餐', '點心'].includes(s.title)) {
		s.data.speaker = []
	}
	return s
}

function setNext(s, idx, arr) {
	if (idx < arr.length - 1) {
		s.next = arr[idx + 1].id
	}
	return s
}

function setIdx(s, idx) {
	s.idx = idx
	return s
}

function main(data) {
	const rooms = data.rooms.map(item => item.zh.name)
	const speakers = Object.fromEntries(data.speakers.map(item => [item.id, item.zh.name]))

	const sessions = data.sessions
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
