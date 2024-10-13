export type Card = {
	id: string
	title: string
	type: string
	speakers: string[]
	room: string
	broadcast: string[]
	start: number
	end: number
}

export const ZeroCard: Card = {
	id: '',
	title: '',
	type: '',
	speakers: [],
	room: '',
	broadcast: [],
	start: 0,
	end: 0,
}
