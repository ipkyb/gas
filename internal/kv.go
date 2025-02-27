package internal

type kv[T any] struct {
	Key   string
	Value T
}
