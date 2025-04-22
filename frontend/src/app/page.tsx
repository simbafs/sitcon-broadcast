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

function pageWithRoom(link: string): Page[] {
	return ['R0', 'R1', 'R2', 'R3', 'S'].map(room => ({
		name: room,
		link: `${link}?room=${room}`,
	}))
}

const pages: Page[] = [
	page('Home', '/'),
	directory('card', [page('admin', '/card/admin'), page('all', '/card/all'), ...pageWithRoom('/card')]),
	directory('counter', [page('admin', '/counter/admin'), ...pageWithRoom('/counter')]),
	directory('debug', [page('debug', '/debug'), page('sse', '/debug/sse')]),
	page('event', '/event'),
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
