export type Session = {
	title: string
	id: string // not unique in the while system, (room, id) or idx is unique
	room: string
	broadcastTo: string[] // for update clients that intreast in this session
	broadcastFrom: string // is the session has broadcastFrom, it cannot be modify
	start: number // unix timestamp in seconds
	end: number // unix timestamp in seconds

	speaker: string

	// links
	qa: string
	slido_id: string
	slido_link: string
	slido_admin_link: string
	co_write: string //
}
