import { ReactNode } from "react";

export function Card({ children }: { children: ReactNode }) {
	return <div className="rounded-lg bg-gray-100 p-4 shadow-md">{children}</div>
}
