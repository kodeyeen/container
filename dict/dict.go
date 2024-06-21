package dict

import "iter"

type Dict[K comparable, V any] map[K]V

func New[K comparable, V any](capacity int) Dict[K, V] {
	return make(Dict[K, V], capacity)
}

func (d Dict[K, V]) Get(key K) (V, bool) {
	val, ok := d[key]
	return val, ok
}

func (d Dict[K, V]) Set(key K, val V) {
	d[key] = val
}

func (d *Dict[K, V]) Delete(key K) {
	delete(*d, key)
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

func (d *Dict[K, V]) Clear() {
	clear(*d)
}
