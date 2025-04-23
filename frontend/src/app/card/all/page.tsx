'use client'

// TODO: 顯示成時間表
import { Card } from '@/components/card'
import { Loading } from '@/components/loading'
import { useSessions } from '@/hooks/useSessions'
import { useQuery } from '@/hooks/util/useQuery'

export default function Page() {
	const [room] = useQuery('room', 'R0')
	const sessions = useSessions(room)

	return (
		<div className="m-6 flex flex-col items-center">
			<div className=" max-w-2xl">
				<h1 className="m-10 text-center text-5xl font-bold">{room}</h1>
				{sessions ? sessions.map(session => <Card key={session.session_id} session={session} />) : <Loading />}
			</div>
		</div>
	)
}
