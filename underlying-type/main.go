package main

import "fmt"

type MyInt int

type MyNumber MyInt

type Numeric interface {
	~int
}

type NumericSlice interface {
	~[]int
}

func printNumeric[T Numeric](v T) {
	fmt.Printf("%T %v\n", v, v)
}

func printNumericSlice[T NumericSlice](v T) {
	fmt.Printf("%T %v\n", v, v)
}

type MyIntSlice []int

func main() {
	// all ~int
	printNumeric(10)
	printNumeric(MyInt(3))
	printNumeric(MyNumber(8))

	fmt.Println("================================")

	// all ~[]int
	i := []int{5, 11} // type literal
	// printNumeric(i) // => compile error: []int does not satisfy Numeric ([]int missing in ~int)
	printNumericSlice(i)
	printNumericSlice(MyIntSlice{6, 13, 17})
}
