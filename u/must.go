package u

func Must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

func Muster[T, U any](f func(T) (U, error)) func(T) U {
	return func(v T) U {
		return Must(f(v))
	}
}
