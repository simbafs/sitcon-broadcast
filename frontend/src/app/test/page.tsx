'use client'
// import { createContext, useContext, useRef, useState } from 'react'

// useEditor.ts
// function Editor({
// 	isOpen,
// 	init,
// 	res,
// 	rej,
// }: {
// 	isOpen: boolean
// 	init: string
// 	res: (value: string) => void
// 	rej: () => void
// }) {
// 	const [value, setValue] = useState(init)
// 	return (
// 		<div style={{ display: isOpen ? 'block' : 'none' }} className="border border-black">
// 			<input value={value} onChange={e => setValue(e.target.value)} onFocus={e => e.target.select()} />
// 			<button type="button" onClick={() => res(value)}>
// 				Submit
// 			</button>
// 			<button type="button" onClick={() => rej()}>
// 				Cancel
// 			</button>
// 		</div>
// 	)
// }
//
// const EditorContext = createContext<(init: string, res: (value: string) => void, rej: () => void) => void>(() => {})
//
// function EditorProvider({ children }: { children: React.ReactNode }) {
// 	const [initValue, setInitValue] = useState('')
// 	const [res, setRes] = useState<(value: string) => void>(() => {})
// 	const [rej, setRej] = useState<() => void>(() => {})
// 	const [isOpen, setIsOpen] = useState(false)
//
// 	const setContext = (init: string, res: (value: string) => void, rej: () => void) => {
// 		console.log('setContext')
// 		setIsOpen(true)
// 		setInitValue(init)
// 		setRes(() => (value: string) => {
// 			res(value)
// 			setIsOpen(false)
// 		})
// 		setRej(() => () => {
// 			rej()
// 			setIsOpen(false)
// 		})
// 	}
//
// 	return (
// 		<EditorContext.Provider value={setContext}>
// 			{children}
// 			<div className="border border-black rounded m-2 p-2">
// 				init value: {initValue}
// 				<Editor isOpen={isOpen} init={initValue} res={res} rej={rej} key={initValue} />
// 			</div>
// 		</EditorContext.Provider>
// 	)
// }
//
// function useEditor() {
// 	const setInitValue = useContext(EditorContext)
//
// 	return (initValue: string) =>
// 		new Promise<string>((res, rej) => {
// 			console.log('in promise')
// 			setInitValue(initValue, res, rej)
// 		})
// }

'use client'
import { FC, useRef, useState } from 'react'

// useEditor.ts
type ModalComponent<T> = FC<{ isOpen: boolean; initValue: T; callback: (value: T, ok: boolean) => void }>

function useEditor<T>(Editor: ModalComponent<T>, zeroValue: T) {
	const [initValue, setInitValue] = useState<T | undefined>(undefined)
	const [callback, setCallback] = useState<(value: T, ok: boolean) => void>(() => {})
	const isOpen = initValue !== undefined
	const [key, setKey] = useState(0)

	const edit = (initValue: T) =>
		new Promise<T>((res, rej) => {
			setKey(key + 1)
			setInitValue(initValue) // open
			setCallback(() => (value: T, ok: boolean) => {
				if (ok) res(value)
				else rej()

				setInitValue(undefined) // close
			})
		})

	return [
		() => <Editor isOpen={isOpen} initValue={initValue || zeroValue} callback={callback} key={key} />,
		edit,
	] as const
}

// page.tsx
export default function Page() {
	return <Child />
}

function MyEditor({
	isOpen,
	initValue,
	callback,
}: {
	isOpen: boolean
	initValue: string
	callback: (value: string, ok: boolean) => void
}) {
	const ref = useRef<HTMLInputElement>(null)

	return (
		<div style={{ display: isOpen ? 'block' : 'none' }} className="border border-black">
			<input onFocus={e => e.target.select()} ref={ref} />
			<button type="button" onClick={() => callback(ref.current?.value || initValue, true)}>
				Submit
			</button>
			<button type="button" onClick={() => callback(initValue, false)}>
				Cancel
			</button>
		</div>
	)
}

function Child() {
	const [Editor, edit] = useEditor(MyEditor, '')
	const ref = useRef<HTMLInputElement>(null)

	return (
		<>
			<input ref={ref} onFocus={e => e.target.select()} />
			<button
				type="button"
				onClick={() => edit(ref.current?.value || 'default').then(console.log, console.error)}
			>
				Set Editor
			</button>
			<Editor />
		</>
	)
}
