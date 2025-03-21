import { JsonEditor } from 'json-edit-react'
import { Card } from './card'

export function ShowJSON({ data, title, setData }: { data: any; setData?: (data: any) => void; title: string }) {
	return (
		<Card className="overflow-scroll">
			<h2 className="text-lg font-semibold">{title}</h2>
			<JsonEditor data={data} setData={setData} collapse={1} viewOnly={!setData} />
		</Card>
	)
}
