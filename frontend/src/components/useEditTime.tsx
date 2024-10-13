import { FC, useRef, useState } from 'react'
import { twMerge } from 'tailwind-merge'

export function MyEditor({
	isOpen,
	initValue,
	callback,
}: {
	isOpen: boolean
	initValue: number
	callback: (value: number, ok: boolean) => void
}) {
	const hour = useRef<HTMLInputElement>(null)
	const minute = useRef<HTMLInputElement>(null)

	return (
		<div
			className={twMerge(
				isOpen ? 'grid' : 'hidden',
				'fixed top-0 left-0 h-screen w-screen bg-black/50 place-items-center',
			)}
			onClick={() => callback(0, false)}
		>
			<form
				className="rounded-lg bg-white flex flex-col gap-4 w-80 items-center py-16 px-8"
				onClick={e => e.stopPropagation()}
				onSubmit={e => {
					e.preventDefault()
					callback(Number(hour.current?.value || 0) * 60 + Number(minute.current?.value || 0), true)
				}}
			>
				<h1>Set Time</h1>
				<div className="flex gap-4 w-fulla">
					<input
						ref={hour}
						className="w-full border-gray-500 border-2 rounded-lg p-1 outline-none focus:border-blue-500"
						type="txt"
						defaultValue={Math.floor(initValue / 60)}
						inputMode="decimal"
						onFocus={e => e.target.select()}
					/>
					:
					<input
						ref={minute}
						className="w-full border-gray-500 border-2 rounded-lg p-1 outline-none focus:border-blue-500"
						type="txt"
						defaultValue={initValue % 60}
						inputMode="decimal"
						onFocus={e => e.target.select()}
					/>
				</div>
				<div className="flex gap-4 w-full">
					<button
						className="rounded-md bg-gray-500 text-white p-2 font-bold w-full"
						onClick={() => callback(0, false)}
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

export type ModalComponent<T> = FC<{ isOpen: boolean; initValue: T; callback: (value: T, ok: boolean) => void }>
export type edit<T> = (value: T) => Promise<number>

export function useEditor<T>(Editor: ModalComponent<T>, zeroValue: T) {
	const [initValue, setInitValue] = useState<T | undefined>(undefined)
	const [callback, setCallback] = useState<(value: T, ok: boolean) => void>(() => {})
	const isOpen = initValue !== undefined
	const [key, setKey] = useState(0)

	const edit = (initValue: T) => {
		// console.log(`edit with ${initValue}`)
		return new Promise<T>((res, rej) => {
			setKey(key + 1)
			setInitValue(initValue) // open
			setCallback(() => (value: T, ok: boolean) => {
				if (ok) res(value)
				else rej()

				setInitValue(undefined) // close
			})
		})
	}

	return [
		() => <Editor isOpen={isOpen} initValue={initValue || zeroValue} callback={callback} key={key} />,
		edit,
	] as const
}
