import { parseAsString, useQueryState } from 'nuqs'

export function useQuery(name: string, defaultValue: string) {
	return useQueryState(name, parseAsString.withDefault(defaultValue))
}
