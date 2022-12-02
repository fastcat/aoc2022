package i

func Muster[T, U any](f func(T) (U, error)) func(T) U {
	return func(t T) U {
		u, err := f(t)
		if err != nil {
			panic(err)
		}
		return u
	}
}
