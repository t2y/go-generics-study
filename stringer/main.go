package main

import "fmt"

type Hex int

func (h Hex) String() string {
	return fmt.Sprintf("%x", int(h))
}

func dumpWithTypeParameter[T fmt.Stringer](v T) {
	// more efficient?
	fmt.Printf("0x%s\n", v.String())
}

func dumpWithInterface(v fmt.Stringer) {
	// good readability?
	fmt.Printf("0x%s\n", v.String())
}

func main() {
	// type parameter (static)
	dumpWithTypeParameter(Hex(1))  // => 0x1
	dumpWithTypeParameter(Hex(15)) // => 0xf
	dumpWithTypeParameter(Hex(16)) // => 0x10

	fmt.Println("================================")

	// interface (dynamic)
	dumpWithInterface(Hex(1))  // => 0x1
	dumpWithInterface(Hex(15)) // => 0xf
	dumpWithInterface(Hex(16)) // => 0x10
}
