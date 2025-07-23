import { Card } from './components/Card'
import { QRCode } from './components/QRcode'
import { useDisplay } from './hooks/useDisplay'

export function App() {
	const { display, id, showQRcode } = useDisplay()

	if (!display) {
		return <div>Connecting......</div>
	} else if (!id) {
		return <div>Waiting ID...</div>
	} else if (showQRcode) {
		return (
			<>
				<QRCode text={id} />
				<h1 className="text-3xl font-bold">{id}</h1>
			</>
		)
	} else {
		return <Card display={display} />
	}
}
