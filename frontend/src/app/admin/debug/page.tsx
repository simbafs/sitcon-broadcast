import { Now } from './now'

export default function Page() {
	return (
		<div className="m-8">
			<h1 className="text-2xl font-bold">Debug</h1>
			<h2 className="text-xl font-bold">Now</h2>
			<Now />
			<hr />
		</div>
	)
}
