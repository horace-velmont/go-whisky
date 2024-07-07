package ptr

// Of returns parameter's pointer
func Of[T any](t T) *T {
	return &t
}
