import ReactJson from 'react-json-view'
import { Card } from './card'

export function ShowJSON({ data, title }: { data: any; title: string }) {
	return (
		<Card>
			<h2 className="text-lg font-semibold">{title}</h2>
			<ReactJson src={data} collapsed={1} indentWidth={2} />
		</Card>
	)
}
