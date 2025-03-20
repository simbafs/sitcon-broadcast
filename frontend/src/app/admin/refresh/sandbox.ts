export function sandbox(script: string, data: any): Promise<any> {
    console.log({script, data})
	return new Promise((resolve, reject) => {
		try {
			// 創建一個獨立的 iframe 作為沙箱環境
			const iframe = document.createElement('iframe')
			iframe.style.display = 'none'
			document.body.appendChild(iframe)

			// 獲取 iframe 的 window 對象，作為沙箱環境
			const sandboxWindow: any = iframe.contentWindow

			if (!sandboxWindow) return reject(new Error('Failed to create sandbox'))

			// 在沙箱內執行腳本
			sandboxWindow.eval(script)
			sandboxWindow.console = console

			// 確保 main 存在
			if (typeof sandboxWindow.main === 'function') {
				const result = sandboxWindow.main(data)
				document.body.removeChild(iframe) // 清理 iframe
				resolve(result)
			} else {
				document.body.removeChild(iframe)
				reject(new Error("Function 'main' is not defined in the script."))
			}
		} catch (error) {
			reject(error)
		}
	})
}
