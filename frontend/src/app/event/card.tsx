import { ReactNode } from 'react'
import { twMerge } from 'tailwind-merge'

export function Card({ children, className }: { children: ReactNode; className?: string }) {
	return <div className={twMerge('rounded-lg bg-gray-100 p-4 shadow-md', className)}>{children}</div>
}
