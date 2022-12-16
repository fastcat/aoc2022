package i

func KeysToMap[K comparable, V any](
	keys Iterable[K],
	value func(K) V,
) map[K]V {
	// could do this with Reduce too
	m := make(map[K]V)
	For(keys, func(k K, _ int) { m[k] = value(k) })
	return m
}

func ValuesToMap[K comparable, V any](
	values Iterable[V],
	key func(V) K,
) map[K]V {
	// could do this with Reduce too
	m := make(map[K]V)
	For(values, func(v V, _ int) { m[key(v)] = v })
	return m
}

func ToMap[T any, K comparable, V any](
	in Iterable[T],
	entry func(T) (K, V),
) map[K]V {
	// could do this with Reduce too
	m := make(map[K]V)
	For(in, func(t T, _ int) { k, v := entry(t); m[k] = v })
	return m
}
