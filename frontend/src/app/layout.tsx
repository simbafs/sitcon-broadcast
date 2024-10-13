import { SSEProvider } from '@/hooks/useSSE'
import './globals.css'
import { Noto_Sans_TC } from 'next/font/google'

const font = Noto_Sans_TC({
	weight: ['400', '700'],
	subsets: ['latin'],
})

type Props = {
	children: React.ReactNode
}

export default function RootLayout({ children }: Props) {
	return (
		<html className={font.className}>
			<body>
				<SSEProvider url="/api/sse" maxLength={3}>
					{children}
				</SSEProvider>
			</body>
		</html>
	)
}
