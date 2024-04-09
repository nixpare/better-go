package main

func filter[T any](v []T, f func(i int, x T) bool) []T {
	u := make([]T, 0)
	for i, x := range v {
		if f(i, x) {
			u = append(u, x)
		}
	}
	return u
}

func reduce[T, S any](v []T, f func(i int, x T, prev S) S, init S) S {
	for i, x := range v {
		init = f(i, x, init)
	}
	return init
}

func transform[T, S any](v []T, f func(i int, x T) S) []S {
    u := make([]S, 0, len(v))
    for i, x := range v {
        u = append(u, f(i, x))
    }
    return u
}
