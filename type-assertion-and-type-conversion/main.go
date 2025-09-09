package main

import "fmt"

type Foo interface {
	foo()
}

type Bar interface {
	bar()
}

type Struct struct {
}

func (s *Struct) foo() {
	fmt.Println("foo")
}

func (s *Struct) bar() {
	fmt.Println("bar")
}

func foo(arg Foo) {
	arg.(Bar).bar()
}

func main() {
	// type assertion
	var s1 interface{} = "hello"
	s2, ok := s1.(string)
	fmt.Printf("s2: %s ok: %t\n", s2, ok)

	var s3 any
	s4, ok := s3.(string)
	fmt.Printf("s4: %s ok: %t\n", s4, ok)

	s := &Struct{}
	foo(s)

	// type conversion
	var a int16 = 0x1234
	b := int32(a)
	fmt.Printf("b: 0x%x\n", b)

	var c int32 = 0x12345678
	d := int16(c)
	fmt.Printf("d: 0x%x\n", d)
}
