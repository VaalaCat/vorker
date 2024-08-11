package utils

import (
	"sort"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSyncMap_InitialState(t *testing.T) {
	var m SyncMap[string, int]

	require.Equal(t, 0, m.Len())
	require.EqualValues(t, []string{}, m.Keys())
	require.EqualValues(t, []int{}, m.Values())
}

func TestSyncMap_Clone(t *testing.T) {
	var m SyncMap[string, int]

	const key1, key2, val1, val2 = "a", "b", 111, 222

	m.Store(key1, val1)

	require.Equal(t, 1, m.Len())
	require.EqualValues(t, []string{key1}, m.Keys())
	require.EqualValues(t, []int{val1}, m.Values())

	v, loaded := m.Load(key1)

	require.True(t, loaded)
	require.EqualValues(t, val1, v)

	v, loaded = m.Load(key2)

	require.False(t, loaded)
	require.EqualValues(t, 0, v)

	var clone = m.Clone()

	require.EqualValues(t, m.Len(), clone.Len())

	v, loaded = clone.Load(key1)

	require.True(t, loaded)
	require.EqualValues(t, val1, v)

	m.Store(key1, val2) // overwrite in original map

	v, loaded = clone.Load(key1)

	require.True(t, loaded)
	require.EqualValues(t, val1, v)

	v, loaded = m.Load(key1)

	require.True(t, loaded)
	require.EqualValues(t, val2, v)
}

func TestSyncMap_Grow(t *testing.T) {
	var m SyncMap[string, int]

	require.EqualValues(t, 0, m.Len())

	m.Grow(3)

	require.EqualValues(t, 0, m.Len())
}

func TestSyncMap_Load(t *testing.T) {
	var m SyncMap[string, int]

	const key = "a"

	value, ok := m.Load(key) // not exists

	require.False(t, ok)
	require.EqualValues(t, 0, value)

	m.Store(key, 111)
	m.Store(key, 111)
	m.Store(key, 111) // repeated call

	value, ok = m.Load(key) // exists

	require.True(t, ok)
	require.EqualValues(t, 111, value)
}

func TestSyncMap_Store(t *testing.T) {
	var m SyncMap[string, int]

	m.Grow(2)

	const (
		key1, key2 = "a", "b"
		val1, val2 = 123, 321
	)

	m.Store(key1, val1)

	require.Equal(t, 1, m.Len())
	require.EqualValues(t, []string{key1}, m.Keys())
	require.EqualValues(t, []int{val1}, m.Values())

	m.Store(key2, val2)
	m.Store(key2, val2) // repeated call

	require.Equal(t, 2, m.Len())

	var wantKeys, gotKeys = []string{key2, key1}, m.Keys()

	sort.Strings(wantKeys)
	sort.Strings(gotKeys)
	require.EqualValues(t, wantKeys, gotKeys)

	var wantValues, gotValues = []int{val1, val2}, m.Values()

	sort.Ints(wantValues)
	sort.Ints(gotValues)
	require.EqualValues(t, wantValues, gotValues)
}

func TestSyncMap_LoadOrStore(t *testing.T) {
	var m SyncMap[string, float64]

	const (
		key        = "a"
		val1, val2 = 123.123, 321.321
	)

	v, loaded := m.LoadOrStore(key, val1)

	require.False(t, loaded)
	require.EqualValues(t, val1, v)
	require.Equal(t, 1, m.Len())

	v, loaded = m.LoadOrStore(key, val2) // another value is passed

	require.True(t, loaded)
	require.EqualValues(t, val1, v)
	require.Equal(t, 1, m.Len())
}

func TestSyncMap_LoadAndDelete(t *testing.T) {
	var m SyncMap[string, int]

	const key, val = "a", 123

	v, loaded := m.LoadAndDelete(key)
	require.False(t, loaded)
	require.EqualValues(t, 0, v)

	m.Store(key, val)

	require.Equal(t, 1, m.Len())
	require.EqualValues(t, []string{key}, m.Keys())
	require.EqualValues(t, []int{val}, m.Values())

	v, loaded = m.LoadAndDelete(key)
	require.True(t, loaded)
	require.EqualValues(t, val, v)

	require.Equal(t, 0, m.Len())
	require.EqualValues(t, []string{}, m.Keys())
	require.EqualValues(t, []int{}, m.Values())

	v, loaded = m.LoadAndDelete(key) //nolint:ineffassign
	v, loaded = m.LoadAndDelete(key) //nolint:ineffassign
	v, loaded = m.LoadAndDelete(key) // repeated call

	require.False(t, loaded)
	require.EqualValues(t, 0, v)

	require.Equal(t, 0, m.Len())
	require.EqualValues(t, []string{}, m.Keys())
	require.EqualValues(t, []int{}, m.Values())
}

func TestSyncMap_Delete(t *testing.T) {
	var m SyncMap[string, int]

	const key, val = "a", 123

	m.Delete(key)
	m.Delete(key) // repeated call

	require.Equal(t, 0, m.Len())
	require.EqualValues(t, []string{}, m.Keys())
	require.EqualValues(t, []int{}, m.Values())

	m.Store(key, val)

	require.Equal(t, 1, m.Len())
	require.EqualValues(t, []string{key}, m.Keys())
	require.EqualValues(t, []int{val}, m.Values())

	m.Delete(key)
	m.Delete(key)
	m.Delete(key) // repeated call

	require.Equal(t, 0, m.Len())
	require.EqualValues(t, []string{}, m.Keys())
	require.EqualValues(t, []int{}, m.Values())
}

func TestSyncMap_Range(t *testing.T) {
	var m SyncMap[string, int]

	const (
		key1, key2 = "a", "b"
		val1, val2 = 123, 321
	)

	var iter uint

	m.Range(func(key string, val int) bool {
		iter++

		return false
	})

	require.EqualValues(t, 0, iter)

	iter = 0 // reset

	m.Store(key1, val1)
	m.Store(key1, val1) // repeated call
	m.Store(key2, val2)
	m.Store(key2, val2) // repeated call

	require.Equal(t, 2, m.Len())

	m.Range(func(key string, val int) bool {
		if key == key1 {
			require.EqualValues(t, val1, val)
		} else if key == key2 {
			require.EqualValues(t, val2, val)
		}

		iter++

		return true
	})

	require.EqualValues(t, 2, iter)

	iter = 0 // reset

	m.Range(func(key string, val int) bool {
		iter++

		return false
	})

	require.EqualValues(t, 1, iter)
}

func TestSyncMap_Struct(t *testing.T) {
	type some struct{ foo string }

	m := SyncMap[[2]int, some]{}
	require.Equal(t, 0, m.Len())

	var key = [2]int{1, 2}

	val, ok := m.Load(key)
	require.False(t, ok)
	require.EqualValues(t, some{}, val) // NOT nil

	m.Store(key, some{"bar"})
	require.Equal(t, 1, m.Len())

	val, ok = m.Load(key)
	require.True(t, ok)
	require.EqualValues(t, some{"bar"}, val)
}

func TestSyncMap_Map(t *testing.T) {
	type some map[uint]sync.Mutex

	m := SyncMap[uint, *some]{}
	require.Equal(t, 0, m.Len())

	var key uint = 1

	val, ok := m.Load(key)
	require.False(t, ok)
	require.Nil(t, val) // nil here is correct

	var mu sync.Mutex

	m.Store(key, &some{1: mu}) //nolint:govet
	require.Equal(t, 1, m.Len())

	val, ok = m.Load(key)
	require.True(t, ok)
	require.EqualValues(t, &some{1: mu}, val) //nolint:govet
}

//go:noinline
func TestNoCopy_ConcurrentUsage(t *testing.T) { // race detector provocation
	var (
		m  SyncMap[string, int]
		wg sync.WaitGroup
	)

	for i := 0; i < 100; i++ {
		wg.Add(12)

		go func() { defer wg.Done(); m.Grow(3) }()
		go func() { defer wg.Done(); _, _ = m.LoadOrStore("foo", 1) }() // +
		go func() { defer wg.Done(); m.Store("foo", 1) }()              // +
		go func() { defer wg.Done(); _, _ = m.Load("foo") }()
		go func() { defer wg.Done(); _, _ = m.LoadAndDelete("foo") }() // -
		go func() { defer wg.Done(); m.Delete("foo") }()               // -
		go func() { defer wg.Done(); m.Range(func(_ string, _ int) bool { return true }) }()
		go func() { defer wg.Done(); m.Range(func(_ string, _ int) bool { return false }) }()
		go func() { defer wg.Done(); _ = m.Len() }()
		go func() { defer wg.Done(); _ = m.Keys() }()
		go func() { defer wg.Done(); _ = m.Values() }()
		go func() { defer wg.Done(); _ = m.Clone() }()
	}

	wg.Wait()
}

// BenchmarkSyncMap_NativeMapMutex-8   	20247865	        60.29 ns/op	       0 B/op	       0 allocs/op
// BenchmarkSyncMap_SyncMapUnderTheHood-8   	 9286198	       131.2 ns/op	      32 B/op	       2 allocs/op
func BenchmarkSyncMap_NativeMap(b *testing.B) {
	b.ReportAllocs()

	var (
		m     = SyncMap[string, int]{}
		v, ok = 0, false
	)

	const key = "a"

	for i := 0; i < b.N; i++ {
		m.Store(key, 1)
		v, ok = m.Load(key)
		m.Delete(key)
	}

	require.True(b, ok)
	require.EqualValues(b, 1, v)
}

// BenchmarkSyncMap_Stdlib-8      	13189734	        93.28 ns/op	      16 B/op	       1 allocs/op
func BenchmarkSyncMap_Stdlib(b *testing.B) {
	b.ReportAllocs()

	var (
		m         = sync.Map{}
		v, ok any = 0, false
	)

	const key = "a"

	for i := 0; i < b.N; i++ {
		m.Store(key, 1)
		v, ok = m.Load(key)
		m.Delete(key)
	}

	require.True(b, ok.(bool))
	require.EqualValues(b, 1, v)
}
