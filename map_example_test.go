package gostruct_test

import (
	"fmt"
	"github.com/yut-kt/gostruct"
)

func ExampleMap2Dim() {
	m2d := gostruct.Map2Dim[string, string, int]{}
	k1, k2, val := "key1", "key2", 1

	// NonSet Get
	fmt.Printf("Get:%v\n", m2d.Get(k1, k2))
	// NonSet Gets
	v, ok := m2d.Gets(k1, k2)
	fmt.Printf("Gets(v:%v, ok:%v)\n", v, ok)
	// NonSet HasKey
	fmt.Printf("HasKey:%v\n", m2d.HasKey(k1, k2))

	// Set
	m2d.Set(k1, k2, val)

	// Get
	fmt.Printf("Get:%v\n", m2d.Get(k1, k2))
	// Gets
	v, ok = m2d.Gets(k1, k2)
	fmt.Printf("Gets(v:%v, ok:%v)\n", v, ok)
	// HasKey
	fmt.Printf("HasKey:%v\n", m2d.HasKey(k1, k2))
	// Output:
	// Get:0
	// Gets(v:0, ok:false)
	// HasKey:false
	// Get:1
	// Gets(v:1, ok:true)
	// HasKey:true
}
