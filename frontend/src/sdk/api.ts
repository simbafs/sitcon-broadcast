type Method = 'POST' | 'GET' | 'PUT' | 'DELETE'

export async function api<T extends any>(path: string, method: Method, body?: any) {
	let input = `/api`

	if (!path.startsWith('/')) input += '/' + path
	else input += path

	if (!input.endsWith('/')) input += '/'

	return fetch(`/api${path}`, {
		method,
		body: JSON.stringify(body),
	})
		.then(res => res.json())
		.then(body => {
			if (body.errors) throw new Error(body.error)
			return {
				...body,
				$schema: undefined,
			} as T
		})
}
