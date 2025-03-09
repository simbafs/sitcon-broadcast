import { useSessions } from '@/hooks/useSessions'
import { parseAsString, useQueryState } from 'nuqs'

export default function Page() {
	const [room] = useQueryState('room', parseAsString.withDefault('R0'))
	const sessions = useSessions(room)

	return <div>
	    <button>下一個</button>
	    <button>撤銷</button>
	</div>
}
