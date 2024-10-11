import { SSEProvider } from '@/hooks/useSSE'
import '@/styles/globals.css'

type Props = {
	children: React.ReactNode
}

export default function RootLayout({ children }: Props) {
	return (
		<html>
			<body>
				<SSEProvider url="/api/sse" maxLength={3}>{children}</SSEProvider>
			</body>
		</html>
	)
}
