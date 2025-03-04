import { FC, useState } from 'react'

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
