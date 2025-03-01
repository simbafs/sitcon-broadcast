import { SSEProvider } from '@/hooks/useSSE'
import './globals.css'
import { Noto_Sans_TC } from 'next/font/google'
import { NuqsAdapter } from 'nuqs/adapters/next/app'

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
				<NuqsAdapter>
					<SSEProvider url="/api/sse" maxLength={3}>
						{children}
					</SSEProvider>
				</NuqsAdapter>
			</body>
		</html>
	)
}
