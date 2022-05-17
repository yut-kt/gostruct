package gostruct_test

import (
	"strconv"
	"testing"

	"github.com/yut-kt/gostruct"
)

func generateKeys() (key1s, key2s []string) {
	N := 1000
	key1s = make([]string, N)
	key2s = make([]string, N)
	for i := 0; i < 1000; i++ {
		s := strconv.Itoa(i)
		key1s[i] = s
		key2s[i] = s
	}
	return
}

type Keys struct {
	key1, key2 string
}

// Generate

func BenchmarkMapMapGenerate(b *testing.B) {
	key1s, key2s := generateKeys()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m := make(map[string]map[string]int)
		for _, key1 := range key1s {
			m[key1] = make(map[string]int)
			for _, key2 := range key2s {
				m[key1][key2] = 1
			}
		}
	}
}

func BenchmarkMapStructGenerate(b *testing.B) {
	key1s, key2s := generateKeys()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m := make(map[Keys]int)
		for _, key1 := range key1s {
			for _, key2 := range key2s {
				m[Keys{key1: key1, key2: key2}] = 1
			}
		}
	}
}

func BenchmarkMap2DimGenerate(b *testing.B) {
	key1s, key2s := generateKeys()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m := gostruct.Map2Dim[string, string, int]{}
		for _, key1 := range key1s {
			for _, key2 := range key2s {
				m.Set(key1, key2, 1)
			}
		}
	}
}

// Access

func BenchmarkMapMapAccess(b *testing.B) {
	key1s, key2s := generateKeys()
	m := make(map[string]map[string]int)
	for _, key1 := range key1s {
		m[key1] = make(map[string]int)
		for _, key2 := range key2s {
			m[key1][key2] = 1
		}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sum := 0
		for _, key1 := range key1s {
			for _, key2 := range key2s {
				sum += m[key1][key2]
			}
		}
	}
}

func BenchmarkMapStructAccess(b *testing.B) {
	key1s, key2s := generateKeys()
	m := make(map[Keys]int)
	for _, key1 := range key1s {
		for _, key2 := range key2s {
			m[Keys{key1: key1, key2: key2}] = 1
		}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sum := 0
		for _, key1 := range key1s {
			for _, key2 := range key2s {
				sum += m[Keys{key1: key1, key2: key2}]
			}
		}
	}
}

func BenchmarkMap2DimAccess(b *testing.B) {
	key1s, key2s := generateKeys()
	m := gostruct.Map2Dim[string, string, int]{}
	for _, key1 := range key1s {
		for _, key2 := range key2s {
			m.Set(key1, key2, 1)
		}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sum := 0
		for _, key1 := range key1s {
			for _, key2 := range key2s {
				sum += m.Get(key1, key2)
			}
		}
	}
}

// Access with check

func BenchmarkMapMapAccessWithCheck(b *testing.B) {
	key1s, key2s := generateKeys()
	m := make(map[string]map[string]int)
	for _, key1 := range key1s {
		m[key1] = make(map[string]int)
		for _, key2 := range key2s {
			m[key1][key2] = 1
		}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sum := 0
		for _, key1 := range key1s {
			for _, key2 := range key2s {
				if v, ok := m[key1][key2]; ok {
					sum += v
				}
			}
		}
	}
}

func BenchmarkMapStructAccessWithCheck(b *testing.B) {
	key1s, key2s := generateKeys()
	m := make(map[Keys]int)
	for _, key1 := range key1s {
		for _, key2 := range key2s {
			m[Keys{key1: key1, key2: key2}] = 1
		}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sum := 0
		for _, key1 := range key1s {
			for _, key2 := range key2s {
				if v, ok := m[Keys{key1: key1, key2: key2}]; ok {
					sum += v
				}
			}
		}
	}
}

func BenchmarkMap2DimAccessWithCheck(b *testing.B) {
	key1s, key2s := generateKeys()
	m := gostruct.Map2Dim[string, string, int]{}
	for _, key1 := range key1s {
		for _, key2 := range key2s {
			m.Set(key1, key2, 1)
		}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sum := 0
		for _, key1 := range key1s {
			for _, key2 := range key2s {
				if v, ok := m.Gets(key1, key2); ok {
					sum += v
				}
			}
		}
	}
}
