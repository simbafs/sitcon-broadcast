import { createContext, useContext, useState } from 'react'
import { twMerge } from 'tailwind-merge'

function NumberInput({ initValue, setInitValue }: { initValue: string; setInitValue: (n: string) => void }) {
	return (
		<input
			className="w-full border-gray-500 border-2 rounded-lg p-1 outline-none focus:border-blue-500"
			type="txt"
			inputMode="decimal"
			value={initValue}
			onChange={e => setInitValue(e.target.value)}
			onFocus={e => e.target.select()}
		/>
	)
}

export function Editor({
	isOpen,
	init,
	res,
	rej,
}: {
	isOpen: boolean
	init: number
	res: (value: number) => void
	rej: () => void
}) {
	const [hour, setHour] = useState(Math.floor(init / 60).toString())
	const [minute, setMinute] = useState((init % 60).toString())

	return (
		<div
			className={twMerge(
				isOpen ? 'grid' : 'hidden',
				'fixed top-0 left-0 h-screen w-screen bg-black/50 place-items-center',
			)}
			onClick={rej}
		>
			<form
				className="rounded-lg bg-white flex flex-col gap-4 w-80 items-center py-16 px-8"
				onClick={e => e.stopPropagation()}
				onSubmit={e => {
					e.preventDefault()
					res(Number(hour) * 60 + Number(minute))
				}}
			>
				<h1>Set Time</h1>
				<div className="flex gap-4 w-fulla">
					<NumberInput initValue={hour} setInitValue={setHour} />
					:
					<NumberInput initValue={minute} setInitValue={setMinute} />
				</div>
				<div className="flex gap-4 w-full">
					<button
						className="rounded-md bg-gray-500 text-white p-2 font-bold w-full"
						onClick={rej}
						type="button"
					>
						Cancel
					</button>
					<button className="rounded-md bg-blue-500 text-white p-2 font-bold w-full" type="submit">
						Save
					</button>
				</div>
			</form>
		</div>
	)
}

const EditorContext = createContext<(init: number, res: (value: number) => void, rej: () => void) => void>(() => {})

export function EditorProvider({
	children,
	Editor,
}: {
	children: React.ReactNode
	Editor: React.FC<{ isOpen: boolean; init: number; res: (value: number) => void; rej: () => void }>
}) {
	const [initValue, setInitValue] = useState(0)
	const [res, setRes] = useState<(value: number) => void>(() => {})
	const [rej, setRej] = useState<() => void>(() => {})
	const [isOpen, setIsOpen] = useState(false)

	const setContext = (init: number, res: (value: number) => void, rej: () => void) => {
		// console.log('setContext')
		setIsOpen(true)
		setInitValue(init)
		setRes(() => (value: number) => {
			res(value)
			setIsOpen(false)
		})
		setRej(() => () => {
			rej()
			setIsOpen(false)
		})
	}

	return (
		<EditorContext.Provider value={setContext}>
			{children}
			<Editor isOpen={isOpen} init={initValue} res={res} rej={rej} key={initValue} />
		</EditorContext.Provider>
	)
}

export function useEditor() {
	const setInitValue = useContext(EditorContext)

	return (initValue: number) =>
		new Promise<number>((res, rej) => {
			// console.log('in promise')
			setInitValue(initValue, res, rej)
		})
}
