package main

import "fmt"

func main() {
	// type assertion
	var s1 interface{} = "hello"
	s2, ok := s1.(string)
	fmt.Printf("s2: %s ok: %t\n", s2, ok)

	var s3 any
	s4, ok := s3.(string)
	fmt.Printf("s4: %s ok: %t\n", s4, ok)

	// type conversion
	var a int16 = 0x1234
	b := int32(a)
	fmt.Printf("b: 0x%x\n", b)

	var c int32 = 0x12345678
	d := int16(c)
	fmt.Printf("d: 0x%x\n", d)
}
