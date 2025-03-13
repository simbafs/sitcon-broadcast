import { tv } from 'tailwind-variants'

export const btn = tv({
	base: 'p-2 rounded m-2',
	variants: {
		color: {
			blue: 'bg-blue-300 hover:bg-blue-400 active:bg-blue-500',
			green: 'bg-green-300 hover:bg-green-400 active:bg-green-500',
			red: 'bg-red-300 hover:bg-red-400 active:bg-red-500',
			yellow: 'bg-yellow-300 hover:bg-yellow-400 active:bg-yellow-500',
			gray: 'bg-gray-300 hover:bg-gray-400 active:bg-gray-500',
		},
		disabled: {
			true: 'opacity-50 cursor-not-allowed',
			false: '',
		},
	},
	defaultVariants: {
		color: 'blue',
		disabled: false,
	},
})
