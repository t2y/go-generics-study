package main

import "fmt"

type Hex int

func (h Hex) String() string {
	return fmt.Sprintf("%x", int(h))
}

func dumpWithTypeParameter[T fmt.Stringer](v T) {
	fmt.Printf("0x%s\n", v.String())
}

func dumpWithInterface(v fmt.Stringer) {
	fmt.Printf("0x%s\n", v.String())
}

func main() {
	// type parameter (static)
	dumpWithTypeParameter(Hex(1))
	dumpWithTypeParameter(Hex(15))
	dumpWithTypeParameter(Hex(16))

	fmt.Println("================================")

	// interface (dynamic)
	dumpWithInterface(Hex(1))
	dumpWithInterface(Hex(15))
	dumpWithInterface(Hex(16))
}
