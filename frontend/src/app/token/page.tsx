'use client'
import { useRouter } from 'next/navigation'
import { useState } from 'react'

export default function Page() {
    const router = useRouter()
    const [token, setToken] = useState('')
    const [valid, setValid] = useState<null | boolean>(null)

    return (
        <div className="w-full flex items-center justify-center h-screen bg-gray-100 gap-4">
            <form
                onSubmit={e => {
                    e.preventDefault()
                    fetch(`/verify?token=${token}`)
                        .then(res => res.json())
                        .then(data => {
                            if (data.error) {
                                setValid(false)
                            } else {
                                router.push('/')
                            }
                        })
                }}
                className="w-80 flex flex-col gap-4 items-center justify-center bg-white rounded-xl p-4 shadow-lg"
            >
                <h1 className="w-full text-center text-2xl">Token</h1>
                <input
                    type="text"
                    value={token}
                    onChange={e => setToken(e.target.value)}
                    className="border-black rounded p-4 border-2 w-full"
                />
                {valid === false && <p>Invalid token</p>}

                <button type="submit" className="border-black rounded p-4 border-2 w-full">
                    Check
                </button>
            </form>
        </div>
    )
}
