export function Loading() {
	return (
		<div className="flex h-full w-full items-center justify-center space-x-2">
			<span className="sr-only">Loading...</span>
			<div className="h-2 w-2 animate-bounce rounded-full bg-black [animation-delay:-0.3s]"></div>
			<div className="h-2 w-2 animate-bounce rounded-full bg-black [animation-delay:-0.15s]"></div>
			<div className="h-2 w-2 animate-bounce rounded-full bg-black"></div>
		</div>
	)
}
