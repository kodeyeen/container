package dict

import "iter"

type Dict[K comparable, V any] map[K]V

func New[K comparable, V any](capacity int) Dict[K, V] {
	return make(Dict[K, V], capacity)
}

func (d Dict[K, V]) All() iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range d {
			if !yield(k, v) {
				return
			}
		}
	}
}
