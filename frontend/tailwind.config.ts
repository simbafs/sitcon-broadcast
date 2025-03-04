import type { Config } from 'tailwindcss'

const config: Config = {
	content: ['./src/**/*.{js,ts,jsx,tsx,mdx}'],
	theme: {
		extend: {
			fontFamily: {
				card: ['GenRyuMin2', 'Noto Serif'],
			},
		},
	},
	plugins: [],
}
export default config
