package gostruct

type Map2Dim[K1, K2 comparable, V any] map[K1]map[K2]V

func (m Map2Dim[K1, K2, V]) Set(key1 K1, key2 K2, val V) {
	if m[key1] == nil {
		m[key1] = make(map[K2]V)
		m[key1][key2] = val
	} else {
		m[key1][key2] = val
	}
}

func (m Map2Dim[K1, K2, V]) Get(key1 K1, key2 K2) V {
	if m[key1] == nil {
		m[key1] = make(map[K2]V)
	}
	return m[key1][key2]
}

func (m Map2Dim[K1, K2, V]) Gets(key1 K1, key2 K2) (val V, ok bool) {
	if m[key1] == nil {
		m[key1] = make(map[K2]V)
	}
	val, ok = m[key1][key2]
	return
}

func (m Map2Dim[K1, K2, V]) HasKey(key1 K1, key2 K2) bool {
	if m[key1] == nil {
		return false
	}
	_, ok := m[key1][key2]
	return ok
}
