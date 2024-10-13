export type Session = {
	id: string
	title: string
	type: string
	speakers: string[]
	room: string
	broadcast: string[]
	start: number
	end: number
}

export const ZeroSession: Session = {
	id: '',
	title: '',
	type: '',
	speakers: [],
	room: '',
	broadcast: [],
	start: 0,
	end: 0,
}

export type Sessions = Record<string, Session[]>
