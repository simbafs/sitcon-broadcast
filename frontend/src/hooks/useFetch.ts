import { useEffect, useState } from 'react'
import { toast } from 'react-toastify'

export function useFetch<T = any>(url: string) {
	const [result, setResult] = useState<T>()

	useEffect(() => {
		fetch(url)
			.then(res => res.json())
			.then(setResult)
			.catch((e: Error) => toast(e.message))
	}, [url])

	return result
}
