'use client'
import { NuqsAdapter } from 'nuqs/adapters/next/app'
import { Noto_Sans_TC } from 'next/font/google'
import { ToastContainer } from 'react-toastify'

import './globals.css'
import { SSEProvider } from '@/hooks/util/useSSE'

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
				<SSEProvider url="/sse">
					<NuqsAdapter>{children}</NuqsAdapter>
				</SSEProvider>
				<ToastContainer position="bottom-right" />
			</body>
		</html>
	)
}
