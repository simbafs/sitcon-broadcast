import { useState } from 'react'

export const input = 'm-2 p-2 bg-blue-200'
export const output = 'm-2 p-2 bg-green-200'

export function useResult() {
    const [result, setResult] = useState('')

    function Result() {
        return <pre className={output}>{result}</pre>
    }

    function set(r: any) {
        if (r instanceof Error) setResult('error: ' + r.message)
        else setResult(JSON.stringify(r, null, 2))
    }

    function clear() {
        setResult('')
    }

    return { Result, set, clear }
}
