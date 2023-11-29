package defs

import (
	"sync"
)

type SyncMap[K comparable, V any] struct {
	data sync.Map
}

func (m *SyncMap[K, V]) Set(key K, value V) {
	m.data.Store(key, value)
}

func (m *SyncMap[K, V]) Get(key K) (V, bool) {
	if store, exist := m.data.Load(key); exist {
		if value, ok := store.(V); ok {
			return value, true
		}
	}
	var zero V
	return zero, false
}

func (m *SyncMap[K, V]) Delete(key K) {
	m.data.Delete(key)
}

func (m *SyncMap[K, V]) FromMap(input map[K]V) {
	for key, value := range input {
		m.Set(key, value)
	}
}

func (m *SyncMap[K, V]) ToMap() map[K]V {
	out := make(map[K]V)
	m.Range(func(key K, value V) bool {
		out[key] = value
		return true
	})
	return out
}

func (m *SyncMap[K, V]) Range(f func(K, V) bool) {
	m.data.Range(func(key, value any) bool {
		k, ok1 := key.(K)
		v, ok2 := value.(V)
		if ok1 && ok2 {
			return f(k, v)
		}
		return false
	})
}

func NewSyncMap[K comparable, V any](input map[K]V) *SyncMap[K, V] {
	m := new(SyncMap[K, V])
	m.FromMap(input)
	return m
}
