package main

import (
	"fmt"
	"reflect"
	"runtime"
)

func foo(a int, b string) {
	_ = a
	_ = b
}

func getFunctionName(f any) {
	pc := reflect.ValueOf(f).Pointer()
	fmt.Printf("pc: 0x%x\n", pc)

	fn := runtime.FuncForPC(pc)

	full_name := fn.Name()
	fmt.Println(full_name)

	file, line := fn.FileLine(pc)
	fmt.Printf("file: %s, line: %d\n", file, line)
}

func getCallerName(skip int) {
	pc, _, _, ok := runtime.Caller(skip)
	if !ok {
		panic("oops")
	}

	fn := runtime.FuncForPC(pc)
	fmt.Println(fn.Name())
}

func main() {
	fmt.Printf("foo ptr: %p\n", foo)
	getFunctionName(foo)

	getCallerName(0) // getCallerName
	getCallerName(1) // main
}
