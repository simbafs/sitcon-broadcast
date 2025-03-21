type Page =
	| {
			name: string
			link: string
	  }
	| {
			name: string
			children: Page[]
	  }

const pages: Page[] = []

function LinkTree({ pages }: { pages: Page[] }) {
	return (
		<ul className="ml-4 list-disc">
			{pages.map(page => {
				if ('link' in page) {
					return (
						<li key={page.link}>
							<a href={page.link}>{page.name}</a>
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
