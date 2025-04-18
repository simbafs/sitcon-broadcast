type Page =
	| {
		name: string
		link: string
	}
	| {
		name: string
		children: Page[]
	}

function page(name: string, link: string): Page {
	return {
		name,
		link,
	}
}

function directory(name: string, children: Page[]): Page {
	return {
		name,
		children,
	}
}

const pages: Page[] = [
	directory('Card', [
		page('R0', '/card?room=R0'),
		page('R1', '/card?room=R1'),
		page('R2', '/card?room=R2'),
		page('R3', '/card?room=R3'),
		page('S', '/card?room=S'),
	]),
	directory('Admin', [
		directory('card', [
			page('R0', '/admin/card?room=R0'),
			page('R1', '/admin/card?room=R1'),
			page('R2', '/admin/card?room=R2'),
			page('R3', '/admin/card?room=R3'),
			page('S', '/admin/card?room=S'),
		]),
		page('debug', '/admin/debug'),
		page('event', '/admin/event'),
	]),
]

function LinkTree({ pages }: { pages: Page[] }) {
	return (
		<ul className="ml-4 list-disc">
			{pages.map(page => {
				if ('link' in page) {
					return (
						<li key={page.link}>
							<a className="underline underline-offset-2 hover:underline-offset-1" href={page.link}>
								{page.name}
							</a>
						</li>
					)
				} else {
					return (
						<li key={page.name}>
							{page.name}
							<LinkTree pages={page.children} />
						</li>
					)
				}
			})}
		</ul>
	)
}

export default function Page() {
	return (
		<div className="mx-20 mt-10 flex flex-col">
			<LinkTree pages={pages} />
		</div>
	)
}
