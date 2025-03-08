export interface Root {
	sessions: Session[]
	speakers: Speaker[]
	session_types: SessionType[]
	rooms: Room[]
	tags: Tag[]
}

export interface Session {
	id: string
	type: string
	room: string
	broadcast?: string[]
	start: string
	end: string
	qa?: string
	slide?: string
	co_write?: string
	record: any
	live: any
	language: any
	uri?: string
	zh: Zh
	en: En
	speakers: string[]
	tags: string[]
}

export interface Zh {
	title: string
	description: string
}

export interface En {
	title: string
	description: string
}

export interface Speaker {
	id: string
	avatar: string
	zh: Zh2
	en: En2
}

export interface Zh2 {
	name: string
	bio: string
}

export interface En2 {
	name: string
	bio: string
}

export interface SessionType {
	id: string
	zh: Zh3
	en: En3
}

export interface Zh3 {
	name: string
	description: string
}

export interface En3 {
	name: string
	description: string
}

export interface Room {
	id: string
	zh: Zh4
	en: En4
}

export interface Zh4 {
	name: string
	description: string
}

export interface En4 {
	name: string
	description: string
}

export interface Tag {
	id: string
	zh: Zh5
	en: En5
}

export interface Zh5 {
	name: string
	description: string
}

export interface En5 {
	name: string
	description: string
}
