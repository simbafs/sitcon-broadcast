import { useEffect, useState } from 'react'

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
	const [currentInterval, setCurrentInterval] = useState(interval)

	useEffect(() => {
		let id: NodeJS.Timeout

		const fetchValue = async () => {
			try {
				const newValue = await getter()
				setValue(newValue)
				setErrorCount(0) // 成功時重置錯誤計數
				setCurrentInterval(interval) // 回復原本的輪詢間隔
			} catch (error) {
				console.error('Error fetching value:', error)
				setErrorCount(prev => prev + 1)

				// 如果錯誤次數超過 maxRetries，則加倍間隔時間，但不超過 maxBackoff
				if (errorCount + 1 >= maxRetries) {
					setCurrentInterval(prev => Math.min(prev * 2, maxBackoff))
				}
			}
		}

		fetchValue() // 立即執行一次
		id = setInterval(fetchValue, currentInterval)

		return () => clearInterval(id)
	// eslint-disable-next-line react-hooks/exhaustive-deps
	}, [getter, currentInterval])

	useEffect(() => console.log({getter, currentInterval}), [getter, currentInterval])

	return value
}
