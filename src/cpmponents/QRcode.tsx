import QRcode from 'qrcode'
import { useEffect, useState } from 'react'
import { twMerge } from 'tailwind-merge'

export function QRCode({ text, className }: { text: string; className?: string }) {
	const [img, setImg] = useState('')
	const [err, setErr] = useState<Error | null>(null)

	useEffect(() => {
		QRcode.toDataURL(text, {
			width: 300,
			margin: 0,
		})
			.then(setImg)
			.catch(err => {
				console.error('Error generating QR code:', err, 'text=', text)
				setErr(new Error(`Failed to generate QR code: ${err.message}`))
			})
	}, [text])

	if (err) {
		return <div>Error generating QR code: {err.message}</div>
	}

	if (!img) {
		return <div>Loading...</div>
	}

	return <img src={img} alt={`QR Code: ${text}`} className={twMerge('h-[300px] w-[300px]', className)} />
}
