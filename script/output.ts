export type Session = {
	// 排序唯一依據
	idx: number 
	// 他們是個大問題
	start: number // unix timestamp in seconds
	end: number // unix timestamp in seconds

	//  純粹為了可以支援不同議程廳而設計的
	id: string // not unique in the while system, (room, id) or idx is unique
	room: string

	// 方便找下一個議程，也可以根據 start、end 來找
	next: string

	// 好像不太需要考慮同步控制，因為那些時候其他廳都是看 R0 直播
	// broadcastTo: string[] // for update clients that intreast in this session
	// broadcastFrom: string // is the session has broadcastFrom, it cannot be modify

	// 有點額外的額外資訊，不過蠻必要的
	title: string

	// extra data
	data: string

	// 額外資訊
	// qa: string
	// slido_id: string
	// slido_link: string
	// slido_admin_link: string
	// co_write: string //
}
