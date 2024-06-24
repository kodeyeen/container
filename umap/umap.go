package umap

import "iter"

type UMap[K comparable, V any] map[K]V

func New[K comparable, V any](capacity int) UMap[K, V] {
	return make(UMap[K, V], capacity)
}

func (m UMap[K, V]) Get(key K) (V, bool) {
	val, ok := m[key]
	return val, ok
}

func (m UMap[K, V]) Set(key K, val V) {
	m[key] = val
}

func (m UMap[K, V]) Delete(key K) {
	delete(m, key)
}

func (m UMap[K, V]) All() iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for key, val := range m {
			if !yield(key, val) {
				return
			}
		}
	}
}

func (m UMap[K, V]) Clear() {
	clear(m)
}

func (m UMap[K, V]) Len() int {
	return len(m)
}

func (m UMap[K, V]) String() string {
	return ""
}
