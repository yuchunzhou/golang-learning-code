package package1

import (
	"fmt"
	"package-usage/package1/package2"
)

func Foo() {
	fmt.Printf("output from package1: %d\n", package2.GLOBAL_V)
}
