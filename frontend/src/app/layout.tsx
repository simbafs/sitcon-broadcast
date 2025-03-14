import { NuqsAdapter } from 'nuqs/adapters/next/app'
import './globals.css'
import { Noto_Sans_TC } from 'next/font/google'
import { ToastContainer } from 'react-toastify'

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
				<NuqsAdapter>{children}</NuqsAdapter>
				<ToastContainer position="bottom-center" />
			</body>
		</html>
	)
}
