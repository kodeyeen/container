package dict

type Dict[K comparable, V any] map[K]V

func New[K comparable, V any](capacity int) Dict[K, V] {
	return make(Dict[K, V], capacity)
}
