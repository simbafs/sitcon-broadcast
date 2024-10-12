'use client'
import { createContext, useContext, useState } from 'react'

// useEditor.ts
function Editor({
	isOpen,
	init,
	res,
	rej,
}: {
	isOpen: boolean
	init: string
	res: (value: string) => void
	rej: () => void
}) {
	const [value, setValue] = useState(init)
	return (
		<div style={{ display: isOpen ? 'block' : 'none' }} className="border border-black">
			<input value={value} onChange={e => setValue(e.target.value)} onFocus={e => e.target.select()} />
			<button type="button" onClick={() => res(value)}>
				Submit
			</button>
			<button type="button" onClick={() => rej()}>
				Cancel
			</button>
		</div>
	)
}

const EditorContext = createContext<(init: string, res: (value: string) => void, rej: () => void) => void>(() => {})

function EditorProvider({ children }: { children: React.ReactNode }) {
	const [initValue, setInitValue] = useState('')
	const [res, setRes] = useState<(value: string) => void>(() => {})
	const [rej, setRej] = useState<() => void>(() => {})
	const [isOpen, setIsOpen] = useState(false)

	const setContext = (init: string, res: (value: string) => void, rej: () => void) => {
		console.log('setContext')
		setIsOpen(true)
		setInitValue(init)
		setRes(() => (value: string) => {
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
			<div className="border border-black rounded m-2 p-2">
				init value: {initValue}
				<Editor isOpen={isOpen} init={initValue} res={res} rej={rej} key={initValue} />
			</div>
		</EditorContext.Provider>
	)
}

function useEditor() {
	const setInitValue = useContext(EditorContext)

	return (initValue: string) =>
		new Promise<string>((res, rej) => {
			console.log('in promise')
			setInitValue(initValue, res, rej)
		})
}

// page.tsx
export default function Page() {
	return (
		<EditorProvider>
			<Child />
		</EditorProvider>
	)
}

function Child() {
	const editor = useEditor()
	const [value, setValue] = useState('')

	return (
		<>
			<input value={value} onChange={e => setValue(e.target.value)} onFocus={e => e.target.select()} />
			<button type="button" onClick={() => editor(value).then(console.log, console.error)}>
				Set Editor
			</button>
		</>
	)
}
