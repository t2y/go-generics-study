package main

import "fmt"

func Keys[K comparable, V any](m map[K]V) []K {
	// https: //cs.opensource.google/go/x/exp/+/97b1e661:maps/maps.go;l=10
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// compile error: invalid map key type K (missing comparable constraint)
/*
func KeysOfAny[K any, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
*/

func compareValue[T comparable](v T) {
	fmt.Println(v, v == v)
}

// compile error: T does not satisfy comparable
/*
func compareAny[T any](v T) {
	compareValue(v)
}
*/

func compareInterface(v any) {
	compareValue(v)
}

/*
func g[T any](t T) {
	f(t)
}
*/

func main() {
	// get keys of a map
	m := map[string]int{"cat": 3, "dog": 2}
	fmt.Println(Keys(m)) // => [cat dog]

	compareValue("test") // => true
	// compareValue([]string{"test"}) // => compile error []string does not satisfy comparable

	compareInterface("string")   // => true
	compareInterface([]string{}) // => panic: runtime error: comparing uncomparable type []string
}
