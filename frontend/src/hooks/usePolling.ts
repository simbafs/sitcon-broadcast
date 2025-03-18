import { useEffect, useState } from 'react'

// WARN: maybe replace this function with a more robust one, this is untested
function isEqual<T>(a: T, b: T) {
	if (typeof a !== typeof b) return false

	if (typeof a === 'object' || Array.isArray(a)) {
		if (Array.isArray(a) && Array.isArray(b) && a.length !== b.length) return false
		for (const key in a) {
			if (!isEqual(a[key], b[key])) return false
		}
		return true
	}

	return a === b
}

interface PollingOpt {
	interval?: number // 初始輪詢間隔
	maxRetries?: number // 允許的最大錯誤次數
	maxBackoff?: number // 最大的輪詢間隔
}

export function usePolling<T>(getter: () => Promise<T>, defaultValue: T, opt: PollingOpt = {}): T {
	const {
		interval = 5000, // 預設 5 秒
		maxRetries = 3, // 預設最多 3 次錯誤
		maxBackoff = 60000, // 預設最大 60 秒間隔
	} = opt

	const [value, setValue] = useState<T>(defaultValue)
	const [errorCount, setErrorCount] = useState(0)
	const currentInterval = Math.min(interval * Math.pow(2, Math.floor(errorCount / maxRetries)), maxBackoff)

	useEffect(() => {
		let id: NodeJS.Timeout

		const fetchValue = async () => {
			try {
				const newValue = await getter()
				setValue(value => {
					if (!isEqual(newValue, value)) {
						setErrorCount(0) // 成功時重置錯誤計數
						return newValue
					}
					return value
				})
			} catch (error) {
				console.error('Error fetching value:', error)
				setErrorCount(prev => prev + 1)
			}
		}

		fetchValue() // 立即執行一次
		id = setInterval(fetchValue, currentInterval)

		return () => clearInterval(id)
	}, [currentInterval, getter])

	return value
}
