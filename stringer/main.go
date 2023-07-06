package main

import "fmt"

type Hex int

func (h Hex) String() string {
	return fmt.Sprintf("%x", int(h))
}

func dump[T fmt.Stringer](v T) {
	fmt.Printf("0x%s\n", v.String())
}

func main() {
	dump(Hex(1))
	dump(Hex(15))
	dump(Hex(16))
}
