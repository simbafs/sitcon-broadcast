type Method = 'POST' | 'GET' | 'PUT' | 'DELETE'

export async function api<T extends any>(path: string, method: Method, body?: any) {
	let input = `/api`

	if (!path.startsWith('/')) input += '/' + path
	else input += path

	if (!input.endsWith('/')) input += '/'

	return fetch(`/api${path}`, {
		method,
		headers: {
			'Content-Type': 'application/json',
		},
		body: JSON.stringify(body),
	})
		.then(res => res.json())
		.then(body => {
			if (body.errors){
				console.error(body)

				throw new Error(`${body.title}: ${body.detail}`)
			}
			return body as T
		})
}
